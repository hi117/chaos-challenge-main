package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

// server handles the transport layer, dispatching requests to the business
// logic.
type server struct {
	player player
}

func (s *server) NewGame(w http.ResponseWriter, req *http.Request) {
	if req.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		_, _ = w.Write(nil)
		return
	}

	game, err := s.player.NewGame()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	resp, err := json.Marshal(game)
	if err != nil {
		http.Error(w, "marshal response: "+err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	// Not much we can do about an error at this point, however in a real
	// production server we'd probably want to log this.
	_, _ = w.Write(resp)
}

func (s *server) Guess(w http.ResponseWriter, req *http.Request) {
	if req.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		_, _ = w.Write(nil)
		return
	}

	var guess struct {
		// ID is the game ID.
		ID string `json:"id"`
		// Guess is the character guessed.
		Guess string `json:"guess"`
	}
	if err := json.NewDecoder(req.Body).Decode(&guess); err != nil {
		http.Error(w, "decode request: "+err.Error(), http.StatusBadRequest)
	}

	// Validate request
	if guess.ID == "" {
		http.Error(w, "guess ID is empty", http.StatusBadRequest)
	}
	if len(guess.Guess) != 1 {
		http.Error(w, "guess should be single character", http.StatusBadRequest)
	}
	guess.Guess = strings.ToUpper(guess.Guess)

	game, err := s.player.Guess(guess.ID, byte(guess.Guess[0]))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	resp, err := json.Marshal(game)
	if err != nil {
		http.Error(w, "marshal response: "+err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// Not much we can do about an error at this point, however in a real
	// production server we'd probably want to log this.
	_, _ = w.Write(resp)
}
