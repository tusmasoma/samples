package dsql

import (
	"os"
	"strings"
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
