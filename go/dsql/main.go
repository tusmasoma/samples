package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	_ "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type Owner struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	City      string `json:"city"`
	Telephone string `json:"telephone"`
}

const (
	REGION = "us-east-1"
)

type Env string

const (
	envPrd   Env = "prd"
	envLocal Env = "local"
)

// GetEnv returns the current environment (production or local).
func GetEnv() Env {
	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}
	return Env(env)
}

func (e Env) IsLocal() bool {
	return e == envLocal
}

func (e Env) IsProduction() bool {
	return strings.HasPrefix(string(e), string(envPrd))
}

func GenerateDbConnectAdminAuthToken(creds *credentials.Credentials, clusterEndpoint string) (string, error) {
	// the scheme is arbitrary and is only needed because validation of the URL requires one.
	endpoint := "https://" + clusterEndpoint
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return "", err
	}
	values := req.URL.Query()
	values.Set("Action", "DbConnectAdmin")
	req.URL.RawQuery = values.Encode()

	signer := v4.Signer{
		Credentials: creds,
	}
	_, err = signer.Presign(req, nil, "dsql", REGION, 15*time.Minute, time.Now())
	if err != nil {
		return "", err
	}

	url := req.URL.String()[len("https://"):]

	return url, nil
}

func getConnection(ctx context.Context, clusterEndpoint string) (*pgx.Conn, error) {
	var connConfig *pgx.ConnConfig
	var err error

	env := GetEnv()
	switch {
	case env.IsProduction():
		// Aurora DSQL 用の接続設定
		url := fmt.Sprintf("postgres://%s:5432/postgres?user=admin&sslmode=verify-full", clusterEndpoint)

		sess, err := session.NewSession()
		if err != nil {
			return nil, err
		}

		creds, err := sess.Config.Credentials.Get()
		if err != nil {
			return nil, err
		}
		staticCredentials := credentials.NewStaticCredentials(
			creds.AccessKeyID,
			creds.SecretAccessKey,
			creds.SessionToken,
		)

		token, err := GenerateDbConnectAdminAuthToken(staticCredentials, clusterEndpoint)
		if err != nil {
			return nil, err
		}

		connConfig, err = pgx.ParseConfig(url)
		if err != nil {
			return nil, err
		}
		connConfig.Password = token
	case env.IsLocal():
		// 通常の PostgreSQL 用の接続設定
		clusterEndpoint = "localhost"
		url := "postgres://postgres:postgres@localhost:5432/test?sslmode=disable"
		connConfig, err = pgx.ParseConfig(url)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unsupported environment")
	}

	conn, err := pgx.ConnectConfig(ctx, connConfig)
	return conn, err
}

func example(clusterEndpoint string) error {
	ctx := context.Background()

	// Establish connection
	conn, err := getConnection(ctx, clusterEndpoint)
	if err != nil {
		return err
	}

	// Create owner table
	_, err = conn.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS owner (
			id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			name VARCHAR(255),
			city VARCHAR(255),
			telephone VARCHAR(255)
		)
	`)
	if err != nil {
		return err
	}

	// insert data
	query := `INSERT INTO owner (id, name, city, telephone) VALUES ($1, $2, $3, $4)`
	_, err = conn.Exec(ctx, query, uuid.New(), "John Doe", "Anytown", "555-555-0150")

	if err != nil {
		return err
	}

	owners := []Owner{}
	// Define the SQL query to insert a new owner record.
	query = `SELECT id, name, city, telephone FROM owner where name='John Doe'`

	rows, err := conn.Query(ctx, query)
	defer rows.Close()

	owners, err = pgx.CollectRows(rows, pgx.RowToStructByName[Owner])
	fmt.Println(owners)
	if err != nil || owners[0].Name != "John Doe" || owners[0].City != "Anytown" {
		panic("Error retrieving data")
	}

	// Delete some data
	_, err = conn.Exec(ctx, `DELETE FROM owner where name='John Doe'`)
	if err != nil {
		return err
	}

	defer conn.Close(ctx)

	return nil
}

func main() {
	cluster_endpoint := "foo0bar1baz2quux3quuux4.dsql.us-east-1.on.aws"
	err := example(cluster_endpoint)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to run example: %v\n", err)
		os.Exit(1)
	}
}
