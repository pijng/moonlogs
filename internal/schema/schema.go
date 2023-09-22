package schema

import (
	"context"
	"log"
	"moonlogs/ent"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func Generate(client *ent.Client) {
	err := entc.Generate("./ent/schema", &gen.Config{})
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
