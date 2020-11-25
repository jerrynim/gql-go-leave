package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	database "github.com/jerrynim/gql-leave/db"
	"github.com/jerrynim/gql-leave/graph/generated"
	"github.com/jerrynim/gql-leave/middleware"
	"github.com/jerrynim/gql-leave/resolver"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func init() {
	// Log error if .env file does not exist
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found")
	}

  }


// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	server := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}}))
	server.SetRecoverFunc(func(ctx context.Context, err interface{}) error {
		log.Print((err))
		
		return errors.New(fmt.Sprintf("%v", err))
	})
	return func(c *gin.Context) {
		server.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	db,err := database.GetDatabase()
	if err!=nil {
		log.Println(err.Error(),"데이터베이스 연결 에러")
	}
	database.RunMigrations(db)
	port := os.Getenv("PORT")
	fmt.Print(port)
    if port == "" {
		port = ":8000"
    }
	
	
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Middleware())

	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())
	r.Run(port)
}
