package main

// func (db *DB) handlerGetGame(w http.ResponseWriter, r *http.Request) {
// 	type parameters struct {
// 		Id string `json:"id"`
// 	}
// 	decoder := json.NewDecoder(r.Body)

// 	params := parameters{}

// 	err := decoder.Decode(&params)

// 	if err != nil {
// 		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
// 		return
// 	}

// 	game, err := db.GetGame(params.Id)

// 	if err != nil {
// 		respondWithError(w, 400, "Invalid game id")
// 		return
// 	}

// 	respondWithJSON(w, 200, game)
// }
