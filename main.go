package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rajeev0928/GoTest/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {

	
	godotenv.Load("D:/web dev/GoProject/.env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

	dbURL := os.Getenv("DB_URL")
	log.Print(dbURL)
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}
	db, errdb := sql.Open("postgres", dbURL)
	if errdb != nil {
		log.Fatal(errdb)
	}
	dbQueries := database.New(db)

	apiCfg := apiConfig{
		DB: dbQueries,
	}

	go startScraping(dbQueries,10,5 * time.Minute)

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1router := chi.NewRouter()

	v1router.Get("/healthz", handlerReadiness)
	v1router.Get("/err", handlerErr)

	v1router.Post("/users", apiCfg.handlerUsersCreate)
	v1router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerUsersGet))

	v1router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerFeedsCreate))
	v1router.Get("/feeds", apiCfg.handlerFeedsGet)

	v1router.Post("/feeds/follow", apiCfg.middlewareAuth(apiCfg.handlerFeedFollowCreate))
	v1router.Get("/feeds/follow", apiCfg.middlewareAuth(apiCfg.handlerFeedFollowsGet))
	v1router.Delete("/feeds/follow/{feed_follow_id}", apiCfg.middlewareAuth(apiCfg.handlerFeedFollowDelete))

	v1router.Get("/posts", apiCfg.middlewareAuth(apiCfg.handlerPostsGet))

	router.Mount("/v1", v1router)

	srv := &http.Server{
		Addr:    ":" + portString,
		Handler: router,
	}

	log.Printf("Server starting on port %v", portString)
	
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
