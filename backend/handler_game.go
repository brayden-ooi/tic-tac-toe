package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GamePayload struct {
	Id     string       `json:"id"`
	State  [3][3]string `json:"state"`
	Status GameStatus   `json:"status"`
	Result string       `json:"result"`
}

func (db *DB) handlerCreateGame(w http.ResponseWriter, r *http.Request) {
	game, err := CreateGame(db)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldnt create game: %v", err))
	}

	respondWithJSON(w, 201, GamePayload{
		Id:     game.id,
		State:  game.state,
		Status: game.status,
		Result: game.result,
	})
}

func (db *DB) handlerGetGame(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Id string `json:"id"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	game, err := db.GetGame(params.Id)

	if err != nil {
		respondWithError(w, 400, "Invalid game id")
		return
	}

	respondWithJSON(w, 200, game)
}

// func handlerUpdateGame(w http.ResponseWriter, r *http.Request) {

// }
