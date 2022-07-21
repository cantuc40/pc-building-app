package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/cantuc40/gqlgen-todos/graph"
	"github.com/cantuc40/gqlgen-todos/graph/generated"
	"github.com/cantuc40/gqlgen-todos/graph/model"
	"github.com/rs/cors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDB() {
	var err error

	//For dataSourceName, change admin to your mysql username and Admin for the password
	dataSourceName := "admin:Admin@tcp(localhost:3306)/?parseTime=True"
	db, err = gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	//db.LogMode(true)

	// Create the database. This is a one-time step.
	// Comment out if running multiple times - You may see an error otherwise
	db.Exec("CREATE DATABASE pc_builder_db")
	db.Exec("USE pc_builder_db")

	start := time.Now()
	// Migration to create tables for all computer parts and users
	db.AutoMigrate(
		&model.User{},
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

	fmt.Println(time.Since(start))
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8083"
	}

	initDB()
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		DB: db,
	}}))

	handler := cors.Default().Handler(srv)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
	http.Handle("/cors", handler)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
