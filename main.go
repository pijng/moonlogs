package main

import (
	"context"
	"log"
	"moonlogs/api/server"
	"moonlogs/ent"
	"moonlogs/internal/config"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	client, err := ent.Open("sqlite3", "file:./database.sqlite?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	err = entc.Generate("./ent/schema", &gen.Config{})
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	config.SetClient(client)

	server.Serve()
}
