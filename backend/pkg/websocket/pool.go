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

type Session struct {
	PlayerToMove *Client
	Players      []*Client
}

type Pool struct {
	Message chan ActionWClient
	Clients map[*Client]*game.Game
	Games   map[string]Session
}

var gameSymbol = map[int]string{
	0: "O",
	1: "X",
}

func NewPool() *Pool {
	return &Pool{
		Message: make(chan ActionWClient),
		Clients: make(map[*Client]*game.Game),
		Games:   make(map[string]Session),
	}
}

func (p *Pool) getGame(id string) (*game.Game, error) {
	session, ok := p.Games[id]

	if !ok {
		return nil, errors.New("invalid id")
	}

	if len(session.Players) != 1 {
		return nil, errors.New("invalid id")
	}

	gamePtr, ok := p.Clients[session.Players[0]]

	if !ok {
		return nil, errors.New("invalid id")
	}

	return gamePtr, nil
}

func (p *Pool) getPlayerSymbol(c *Client, id string) (string, error) {
	for p, v := range p.Games[id].Players {
		if v == c {
			return gameSymbol[p], nil
		}
	}
	return "", errors.New("no player found")
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
				pool.Games[id] = Session{PlayerToMove: message.c, Players: []*Client{message.c}}
				message.c.RespondJSON(gameIns)

			} else if message.act.Type == "join" {
				id, ok := message.act.Payload.(string)
				if !ok {
					message.c.RespondError("unrecognized payload format")
					break
				}

				session, ok := pool.Games[id]

				// validate if game id is valid
				if !ok {
					message.c.RespondError("invalid id submitted")
					break
				}

				// validate if game already has two players
				if len(session.Players) != 1 {
					message.c.RespondError("cannot join game")
					break
				}

				// try grabbing game ref
				gamePtr, err := pool.getGame(id)
				if err != nil {
					message.c.RespondError(err.Error())
					break
				}

				// add Client to game and send game state
				pool.Clients[message.c] = gamePtr
				playerArr := pool.Games[id].Players
				playerArr = append(playerArr, message.c)
				pool.Games[id] = Session{PlayerToMove: pool.Games[id].PlayerToMove, Players: playerArr}

				message.c.RespondJSON(&gamePtr)
			} else if message.act.Type == "update" {
				x, y, err := GetJoinPayload(message.act.Payload)
				if err != nil {
					message.c.RespondError(err.Error())
					break
				}

				// validate if the game exist
				game, ok := pool.Clients[message.c]
				if !ok {
					fmt.Println(message.c, pool)
					message.c.RespondError("no game registered")
					break
				}

				// validate if the player is registered
				symbol, err := pool.getPlayerSymbol(message.c, game.Id)
				if err != nil {
					message.c.RespondError("no game registered 2")
					break
				}

				// validate if there is enough players
				if len(pool.Games[game.Id].Players) != 2 {
					message.c.RespondError("cannot start the game")
					break
				}

				// validate if the player should make the move
				playerToMove := pool.Games[game.Id].PlayerToMove
				if playerToMove != message.c {
					message.c.RespondError("another player has yet to make a move")
					break
				}

				nextGame, err := game.Move(symbol, x, y)
				if err != nil {
					message.c.RespondError(err.Error())
					break
				}

				for _, client := range pool.Games[game.Id].Players {
					// update player to move
					if playerToMove != client {
						nextSession := Session{PlayerToMove: client, Players: pool.Games[game.Id].Players}
						pool.Games[game.Id] = nextSession
					}

					client.RespondJSON(&nextGame)
				}
			} else {
				message.c.RespondError("unrecognized payload")
			}
		}
	}
}
