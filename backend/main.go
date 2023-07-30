package main

import (
	"log"
	"net/http"
	"os"

	"github.com/BrayOOi/tic-tac-toe/pkg/websocket"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func initWS(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.Upgrade(w, r)

	if err != nil {
		log.Print("upgrade:", err)
	}

	client := &websocket.Client{
		Conn: ws,
		Pool: pool,
	}

	client.Read()
}

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not found in the env file")
	}

	pool := websocket.NewPool()
	go pool.Start()

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/health", handlerReadiness)

	// game
	v1Router.Get("/game", func(w http.ResponseWriter, r *http.Request) {
		initWS(pool, w, r)
	})

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Printf("Server starting on PORT %s\n", port)

	log.Fatal(srv.ListenAndServe())
}
