package database

import (
	"backend/pkg/ent"
	"context"
	"log"

	_ "github.com/lib/pq"
)

func NewPostgreSQL(uri string) (*ent.Client, error) {
	client, err := ent.Open("postgres", uri)

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client, err
}
