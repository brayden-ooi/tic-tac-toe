package websocket

import (
	"errors"
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
	Games   map[string][]*Client
}

func NewPool() *Pool {
	return &Pool{
		Message: make(chan ActionWClient),
		Clients: make(map[*Client]*game.Game),
		Games:   make(map[string][]*Client),
	}
}

func (p *Pool) getGame(id string) (*game.Game, error) {
	p1, ok := p.Games[id]

	if !ok {
		return nil, errors.New("invalid id")
	}

	if len(p1) != 1 {
		return nil, errors.New("invalid id")
	}

	gamePtr, ok := p.Clients[p1[0]]

	if !ok {
		return nil, errors.New("invalid id")
	}

	return gamePtr, nil
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

				id, gameIns, err := game.NewGame()
				if err != nil {
					message.c.RespondError(fmt.Sprintf("Couldnt create game: %v", err))
					break
				}

				// update the pool
				pool.Clients[message.c] = &gameIns
				pool.Games[id] = []*Client{message.c}
				message.c.RespondJSON(gameIns)

			} else if message.act.Type == "join" {
				clientArr, ok := pool.Games[message.act.Payload]

				// validate if game id is valid
				if !ok {
					message.c.RespondError("invalid id submitted")
					break
				}

				// validate if game already has two players
				if len(clientArr) != 1 {
					message.c.RespondError("cannot join game")
					break
				}

				// try grabbing game ref
				gamePtr, err := pool.getGame(message.act.Payload)
				if err != nil {
					message.c.RespondError(err.Error())
					break
				}

				// add Client to game and send game state
				pool.Clients[message.c] = gamePtr
				pool.Games[message.act.Payload] = append(pool.Games[message.act.Payload], message.c)

				message.c.RespondJSON(&gamePtr)
			} else {
				fmt.Println("game updated!")
			}
		}
	}
}
