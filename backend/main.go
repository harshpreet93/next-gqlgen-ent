package backend

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/harshpreet93/next-gqlgen-ent/backend/ent"
	"github.com/harshpreet93/next-gqlgen-ent/backend/ent/migrate"
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

}
