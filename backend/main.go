package main

//go:generate go run ./scripts/gqlgen.go
import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/go-sql-driver/mysql"
	"github.com/harshpreet93/next-gqlgen-ent/backend/ent"
	"github.com/harshpreet93/next-gqlgen-ent/backend/ent/migrate"
	"github.com/harshpreet93/next-gqlgen-ent/backend/graph"
	"github.com/harshpreet93/next-gqlgen-ent/backend/graph/generated"
	"github.com/joho/godotenv"
)

func migrateDB(conn *ent.Client) error {
	ctx := context.Background()
	err := conn.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	return err
}

func initDBConn(dbHost string, dbPort int, dbPass string) *ent.Client {
	client, err := ent.Open("mysql",
		fmt.Sprintf("app:%s@tcp(%s:%d)/app?parseTime=true", dbPass, dbHost, dbPort))
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	return client
}

// Defining the Graphql handler
func graphqlHandler(dbClient *ent.Client) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file

	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &graph.Resolver{
			DBClient: dbClient,
		},
	}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
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
	if os.Getenv("ENV") != "prod" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalf("Error parsing port from DB_PORT %+v", dbPort)
	}
	dbPass := os.Getenv("DB_PW")
	conn := initDBConn(dbHost, dbPort, dbPass)
	defer conn.Close()
	err = migrateDB(conn)

	r := gin.Default()
	r.POST("/query", graphqlHandler(conn))
	r.GET("/", playgroundHandler())
	r.Run()
}
