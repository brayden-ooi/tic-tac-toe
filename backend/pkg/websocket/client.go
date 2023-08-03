package websocket

import (
	"errors"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
}

// inbound message format
type Action struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

// outbound message format
type Message struct {
	Type int         `json:"type"`
	Body interface{} `json:"body"`
}

func (c *Client) Read() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		var msg Action
		err := c.Conn.ReadJSON(&msg)
		if err != nil {
			log.Println(err)
			break
		}

		c.Pool.Message <- ActionWClient{act: msg, c: c}
		fmt.Printf("Message Received: %+v\n", msg)
	}
}

func (c *Client) RespondError(err string) {
	c.Conn.WriteJSON(Message{Type: 0, Body: err})
}

func (c *Client) RespondJSON(payload interface{}) {
	c.Conn.WriteJSON(Message{Type: 1, Body: payload})
}

func GetJoinPayload(payload interface{}) (int, int, error) {
	arr, ok := payload.([]interface{})
	if !ok {
		return 0, 0, errors.New("unrecognized payload format")
	}

	x, ok := arr[0].(float64)
	if !ok {
		return 0, 0, errors.New("unrecognized payload format")
	}

	y, ok := arr[1].(float64)
	if !ok {
		return 0, 0, errors.New("unrecognized payload format")
	}

	return int(x), int(y), nil
}
