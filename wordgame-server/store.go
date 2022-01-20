package main

import "github.com/pkg/errors"

type MemoryStore struct {
	games map[string]GameState
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		games: make(map[string]GameState),
	}
}

func (s *MemoryStore) SaveGame(game GameState) error {
	s.games[game.ID] = game
	return nil
}

func (s *MemoryStore) LoadGame(gameID string) (*GameState, error) {
	game, ok := s.games[gameID]
	if !ok {
		return nil, errors.Errorf("not found: %s", gameID)
	}
	return &game, nil
}

func (s *MemoryStore) DeleteGame(gameID string) error {
	_, ok := s.games[gameID]
	if !ok {
		return errors.Errorf("not found: %s", gameID)
	}
	delete(s.games, gameID)
	return nil
}
