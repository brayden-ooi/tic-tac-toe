package main

import (
	"errors"
)

// database
type DB struct {
	games map[string]Game
}

func initDB() DB {
	games := make(map[string]Game)

	kvStore := DB{games}

	return kvStore
}

func (db *DB) GetGame(id string) (Game, error) {
	game, ok := db.games[id]

	if !ok {
		return game, errors.New("invalid game id")
	}

	return game, nil
}
