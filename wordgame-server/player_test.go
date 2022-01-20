package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWinGame(t *testing.T) {
	// Single word list so we know what to expect
	words := []string{"CAT"}
	store := NewMemoryStore()
	player := player{words: words, store: store}

	game, err := player.NewGame()
	require.NoError(t, err)
	assert.Equal(t, "CAT", game.Word)
	assert.Equal(t, "___", game.Current)

	game, err = player.Guess(game.ID, 'B')
	require.NoError(t, err)
	assert.Equal(t, "___", game.Current)

	game, err = player.Guess(game.ID, 'C')
	require.NoError(t, err)
	assert.Equal(t, "C__", game.Current)

	game, err = player.Guess(game.ID, 'T')
	require.NoError(t, err)
	assert.Equal(t, "C_T", game.Current)

	game, err = player.Guess(game.ID, 'A')
	require.NoError(t, err)
	assert.Equal(t, "CAT", game.Current)
	// Win!

	// Should be deleted now
	game, err = player.Guess(game.ID, 'A')
	require.Error(t, err)
}

func TestLoseGame(t *testing.T) {
	// Single word list so we know what to expect
	words := []string{"CAT"}
	store := NewMemoryStore()
	player := player{words: words, store: store}

	game, err := player.NewGame()
	require.NoError(t, err)
	assert.Equal(t, "CAT", game.Word)
	assert.Equal(t, "___", game.Current)

	game, err = player.Guess(game.ID, 'B')
	require.NoError(t, err)
	assert.Equal(t, "___", game.Current)

	game, err = player.Guess(game.ID, 'F')
	require.NoError(t, err)
	assert.Equal(t, "___", game.Current)

	game, err = player.Guess(game.ID, 'M')
	require.NoError(t, err)
	assert.Equal(t, "___", game.Current)

	game, err = player.Guess(game.ID, 'D')
	require.NoError(t, err)
	assert.Equal(t, "___", game.Current)

	game, err = player.Guess(game.ID, 'X')
	require.NoError(t, err)
	assert.Equal(t, "___", game.Current)

	game, err = player.Guess(game.ID, 'Y')
	require.NoError(t, err)
	assert.Equal(t, "___", game.Current)
	// Loss!

	// Should be deleted now
	game, err = player.Guess(game.ID, 'L')
	require.Error(t, err)
}
