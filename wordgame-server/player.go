package main

import (
	"math/rand"
	"strings"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

const (
	startingGuesses = 6
)

type GameState struct {
	// ID uniquely identifies the game.
	ID string `json:"id"`
	// Word is the actual word that the user is guessing.
	Word string `json:"-"` // Do not include in JSON
	// Current is the current state of the board.
	Current string `json:"current"`
	// GuessesRemaining is how many guesses are left.
	GuessesRemaining int `json:"guesses_remaining"`
}

// CheckWin checks whether the game is in a winning state.
func (g *GameState) CheckWin() bool {
	for _, char := range g.Current {
		if char == '_' {
			return false
		}
	}

	return true
}

type Store interface {
	SaveGame(game GameState) error
	LoadGame(gameID string) (*GameState, error)
	DeleteGame(gameID string) error
}

// player handles the business logic of the game
type player struct {
	store Store
	words []string
}

func (g *player) NewGame() (*GameState, error) {
	wordIndex := rand.Intn(len(g.words))
	word := g.words[wordIndex]

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.Wrap(err, "generate game ID")
	}

	game := GameState{
		ID:               id.String(),
		Word:             word,
		Current:          strings.Repeat("_", len(word)),
		GuessesRemaining: startingGuesses,
	}
	if err := g.store.SaveGame(game); err != nil {
		return nil, errors.Wrap(err, "save initial game state")
	}

	return &game, nil
}

func (p *player) Guess(gameID string, guess byte) (*GameState, error) {
	game, err := p.store.LoadGame(gameID)
	if err != nil {
		return nil, errors.Wrap(err, "load game")
	}

	matched := false
	newCurrent := []byte{}
	for i, char := range game.Current {
		if guess == game.Word[i] {
			matched = true
			newCurrent = append(newCurrent, guess)
		} else {
			newCurrent = append(newCurrent, byte(char))
		}
	}
	game.Current = string(newCurrent)

	if matched && game.CheckWin() {
		if err := p.store.DeleteGame(gameID); err != nil {
			return nil, errors.Wrap(err, "clear won game")
		}

		return game, nil
	}

	if !matched {
		game.GuessesRemaining--
	}
	if game.GuessesRemaining == 0 {
		if err := p.store.DeleteGame(gameID); err != nil {
			return nil, errors.Wrap(err, "clear lost game")
		}
		return game, nil
	}

	if err := p.store.SaveGame(*game); err != nil {
		return nil, errors.Wrap(err, "update save state")
	}

	return game, nil
}
