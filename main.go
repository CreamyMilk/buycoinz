package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/CreamyMilk/buycoinz/schema"
	"github.com/CreamyMilk/buycoinz/storage"
	"github.com/graphql-go/handler"
	"github.com/joho/godotenv"
)

func configureEnvironementVariables() {
	ENV := os.Getenv("ENV")
	if ENV == "" {
		err := godotenv.Load()
		log.Println("Loading ENV from file")
		if err != nil {
			log.Fatal("Error loading .env file", err)
		}
	}
	if os.Getenv("PORT") == "" {
		os.Setenv("PORT", "8080")
	}
}
func main() {
	configureEnvironementVariables()
	storage.InitalizeDB()
	schema.Init()

	h := handler.New(&handler.Config{
		Schema:   &schema.Schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", h)

	port := ":" + os.Getenv("PORT")
	fmt.Printf("🚀 Server ready at http://localhost%s/graphql\n", port)

	http.ListenAndServe(port, nil)
}
