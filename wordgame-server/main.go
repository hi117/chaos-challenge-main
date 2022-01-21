package main

import (
	"go.elastic.co/apm/module/apmhttp"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	mux := http.NewServeMux()

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	words, err := loadWords("words.txt")
	if err != nil {
		log.Fatal(err)
	}

	store := NewMemoryStore()
	player := player{store: store, words: words}
	server := server{player: player}

	mux.HandleFunc("/new", server.NewGame)
	mux.HandleFunc("/guess", server.Guess)

	log.Printf("Serving with %d words", len(words))
	if err := http.ListenAndServe(":"+port, apmhttp.Wrap(mux)); err != nil {
		log.Fatal(err)
	}

}
