package game

import (
	"errors"
	"log"

	"github.com/google/uuid"
)

type Game struct {
	Id     string       `json:"id"`
	State  [3][3]string `json:"state"`
	Status GameStatus   `json:"status"`
	Result string       `json:"result"` // winnter | "stale"
}

type GameStatus string

const (
	GameStatusPlaying   GameStatus = "playing"
	GameStatusCompleted GameStatus = "completed"
)

func NewGame() (string, Game, error) {
	var state [3][3]string

	id := (uuid.New()).String()

	if id == "" {
		return "", Game{}, errors.New("id generation failed")
	}

	return id, Game{
		Id:     id,
		State:  state,
		Status: GameStatusPlaying,
	}, nil
}

func (game *Game) Move(s string, x, y int) *Game {
	if game.Status != GameStatusPlaying {
		log.Fatal(errors.New("game has ended"))
		// return nil, errors.New("game has ended")
	}

	if x < 0 || x > 2 || y < 0 || y > 2 {
		log.Fatal(errors.New("move not allowed"))
		// return nil, errors.New("move not allowed")
	}

	if game.State[y][x] != "" {
		log.Fatal(errors.New("cell already taken"))
		// return nil, errors.New("cell already taken")
	}

	game.State[y][x] = s

	if winner := game.CheckWin(); winner != "" {
		game.Status = GameStatusCompleted
		game.Result = winner
	} else if isStale := game.CheckStale(); isStale {
		game.Status = GameStatusCompleted
		game.Result = "stale"
	}

	return game
}

func (game *Game) CheckWin() string {
	// check if row line
	for _, r := range game.State {
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
		cellTaken := game.State[0][x]

		if cellTaken == "" {
			continue
		}

		for y := 1; y < 3; y++ {
			if game.State[y][x] == "" {
				cellTaken = ""
				break
			}

			if cellTaken != game.State[y][x] {
				cellTaken = ""
				break
			}

			if y == 2 {
				return cellTaken
			}
		}
	}

	// check diagonal
	if game.State[0][0] != "" && game.State[0][0] == game.State[1][1] && game.State[0][0] == game.State[2][2] {
		return game.State[0][0]
	}

	if game.State[0][2] != "" && game.State[0][2] == game.State[1][1] && game.State[0][2] == game.State[2][0] {
		return game.State[0][2]
	}

	return ""
}

func (game *Game) CheckStale() bool {
	for y, r := range game.State {
		for x := range r {
			if game.State[y][x] == "" {
				return false
			}
		}
	}

	return true
}

func (game *Game) Print() string {
	var output = "\n"

	for y, r := range game.State {
		for x, cell := range r {
			if cell != "" {
				output += game.State[y][x]
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
