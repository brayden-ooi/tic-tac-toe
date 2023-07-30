package websocket

import (
	"fmt"

	"github.com/BrayOOi/tic-tac-toe/pkg/game"
)

type ActionWClient struct {
	act Action
	c   *Client
}

type Pool struct {
	Message chan ActionWClient
	Clients map[*Client]*game.Game
	Games   map[*game.Game][]*Client
}

func NewPool() *Pool {
	return &Pool{
		Message: make(chan ActionWClient),
		Clients: make(map[*Client]*game.Game),
		Games:   make(map[*game.Game][]*Client),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case message := <-pool.Message:
			if message.act.Type == "create" {
				// validate if client already has a game associated
				if _, ok := pool.Clients[message.c]; ok {
					message.c.RespondError("not allowed to start another game")
					break
				}

				_, gameIns, err := game.NewGame()
				if err != nil {
					message.c.RespondError(fmt.Sprintf("Couldnt create game: %v", err))
					break
				}

				// update the pool
				pool.Clients[message.c] = &gameIns
				pool.Games[&gameIns] = []*Client{message.c}
				message.c.RespondJSON(gameIns)

			}
		}
	}
}
