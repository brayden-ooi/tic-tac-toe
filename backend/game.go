package main

import (
	"errors"
	"log"
)

type Game struct {
	state [3][3]string
	// p1     uuid.UUID
	// p2     uuid.UUID
	status GameStatus
	result string // winner | "stale"
}

type GameStatus string

const (
	GameStatusPlaying   GameStatus = "playing"
	GameStatusCompleted GameStatus = "completed"
)

func CreateGame() Game {
	var state [3][3]string

	return Game{
		state:  state,
		status: GameStatusPlaying,
	}
}

func (game *Game) Move(s string, x, y int) *Game {
	if game.status != GameStatusPlaying {
		log.Fatal(errors.New("game has ended"))
		// return nil, errors.New("game has ended")
	}

	if x < 0 || x > 2 || y < 0 || y > 2 {
		log.Fatal(errors.New("move not allowed"))
		// return nil, errors.New("move not allowed")
	}

	if game.state[y][x] != "" {
		log.Fatal(errors.New("cell already taken"))
		// return nil, errors.New("cell already taken")
	}

	game.state[y][x] = s

	if winner := game.CheckWin(); winner != "" {
		game.status = GameStatusCompleted
		game.result = winner
	} else if isStale := game.CheckStale(); isStale {
		game.status = GameStatusCompleted
		game.result = "stale"
	}

	return game
}

func (game *Game) CheckWin() string {
	// check if row line
	for _, r := range game.state {
		cellTaken := r[0]

		if cellTaken == "" {
			continue
		}

		for i, cell := range r {
			if cell == "" {
				cellTaken = ""
				break
			}

			if cellTaken != cell {
				cellTaken = ""
				break
			}

			if i == 2 {
				return cellTaken
			}
		}
	}

	// check if col line
	for x := 0; x < 3; x++ {
		cellTaken := game.state[0][x]

		if cellTaken == "" {
			continue
		}

		for y := 1; y < 3; y++ {
			if game.state[y][x] == "" {
				cellTaken = ""
				break
			}

			if cellTaken != game.state[y][x] {
				cellTaken = ""
				break
			}

			if y == 2 {
				return cellTaken
			}
		}
	}

	// check diagonal
	if game.state[0][0] != "" && game.state[0][0] == game.state[1][1] && game.state[0][0] == game.state[2][2] {
		return game.state[0][0]
	}

	if game.state[0][2] != "" && game.state[0][2] == game.state[1][1] && game.state[0][2] == game.state[2][0] {
		return game.state[0][2]
	}

	return ""
}

func (game *Game) CheckStale() bool {
	for y, r := range game.state {
		for x := range r {
			if game.state[y][x] == "" {
				return false
			}
		}
	}

	return true
}

func (game *Game) Print() string {
	var output = "\n"

	for y, r := range game.state {
		for x, cell := range r {
			if cell != "" {
				output += game.state[y][x]
			} else {
				output += " "
			}

			if x != 2 {
				output += "|"
			}
		}

		output += "\n"

		if y != 2 {
			output += "-+-+-\n"
		}
	}

	return output
}
