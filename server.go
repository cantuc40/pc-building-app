package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/cantuc40/gqlgen-todos/graph"
	"github.com/cantuc40/gqlgen-todos/graph/generated"
	"github.com/cantuc40/gqlgen-todos/graph/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const defaultPort = "8083"

var db *gorm.DB

func initDB() {
	var err error
	dataSourceName := "root:@tcp(localhost:3306)/?parseTime=True"
	db, err = gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	//db.LogMode(true)

	// Create the database. This is a one-time step.
	// Comment out if running multiple times - You may see an error otherwise
	db.Exec("CREATE DATABASE test_db")
	db.Exec("USE test_db")

	// Migration to create tables for Order and Item schema
	db.AutoMigrate(
		&model.User{},
		&model.PartsDb{},
		&model.Motherboard{},
		&model.CPU{},
		&model.Storage{},
		&model.Memory{},
		&model.PowerSupply{},
		&model.GraphicsCard{},
		&model.Case{},
		&model.Monitor{},
		&model.OperatingSystem{},
	)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	//initDB()
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		//DB: db,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
