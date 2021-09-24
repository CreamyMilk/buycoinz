package schema

import (
	"log"

	"github.com/graphql-go/graphql"
)

var Schema graphql.Schema

func Init() {
	var err error
	setupQueries()
	setupMuations()
	rootQuery := graphql.ObjectConfig{Name: "Query", Fields: queryFields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery), Mutation: mutationsTypes}
	Schema, err = graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}
}
