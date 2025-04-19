package dsql

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/jackc/pgx/v5"
	"github.com/tusmasoma/go-microservice-k8s/go/pkg/log"
)

func NewPostgresConn(ctx context.Context) (*pgx.Conn, error) {
	conn, err := getConnection(ctx)
	if err != nil {
		log.Error("Failed to connect to database", log.Ferror(err))
		return nil, err
	}
	if err := conn.Ping(ctx); err != nil {
		log.Error("Failed to ping database", log.Ferror(err))
		return nil, err
	}
	return conn, nil
}

const (
	REGION = "us-east-1"
)

const (
	dbPrefix = "POSTGRES_"
)

func generateDbConnectAdminAuthToken(creds *credentials.Credentials, clusterEndpoint string) (string, error) {
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

func getConnection(ctx context.Context) (*pgx.Conn, error) {
	var connConfig *pgx.ConnConfig
	conf, err := NewDBConfig(ctx, dbPrefix)
	if err != nil {
		log.Error("Failed to load database config", log.Ferror(err))
		return nil, err
	}
	env := GetEnv()
	switch {
	case env.IsProduction():
		// Aurora DSQL 用の接続設定
		// cluster_endpoint := "foo0bar1baz2quux3quuux4.dsql.us-east-1.on.aws"
		url := fmt.Sprintf("postgres://%s:%s/%s?user=%s&sslmode=verify-full",
			conf.ClusterEndpoint,
			conf.Port,
			conf.DBName,
			conf.User,
		)

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

		token, err := generateDbConnectAdminAuthToken(staticCredentials, conf.ClusterEndpoint)
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
		// clusterEndpoint = "localhost"
		url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			conf.User, conf.Password, conf.ClusterEndpoint, conf.Port, conf.DBName)
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
