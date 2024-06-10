package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"todoapp/controllers"
	"todoapp/internal/database"
	"todoapp/router"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// load the env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// get the port number
	portNumber := os.Getenv("PORT")
	fmt.Println("The port number is ", portNumber)
	if portNumber == "" {
		log.Fatal("No port number provided")
	}
	// get the database credentiantials
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB url is not found in env variables")
	}
	conn, err := sql.Open("postgres", dbURL)
	apiCfg := controllers.ApiConf{DB: database.New(conn)}
	if err != nil {
		log.Fatal("Cannot connect to the database", err)
	}
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	userRouter := router.RouteUser(&apiCfg)
	todoRouter := router.RouteTodo(&apiCfg)
	r.Mount("/user", userRouter)
	r.Mount("/todo", todoRouter)
	srv := &http.Server{
		Addr:    ":" + portNumber,
		Handler: r,
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("There was an error with the server", err)
	}
}
