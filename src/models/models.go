package models

import (
	"os"

	_ "github.com/lib/pq"
	"github.com/sergot/tibiago/ent"
)

func ConnectDatabase() (*ent.Client, error) {
	connection_env := "DATABASE_URL"

	client, err := ent.Open("postgres", os.Getenv(connection_env))
	if err != nil {
		return nil, err
	}

	return client, nil
}
