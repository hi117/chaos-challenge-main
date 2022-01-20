package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

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

	http.HandleFunc("/new", server.NewGame)
	http.HandleFunc("/guess", server.Guess)

	log.Printf("Serving with %d words", len(words))
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}

}
