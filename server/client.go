package main

import (
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	manager  *RoomManager
	msgqueue chan string
	ws       *websocket.Conn
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// Allow specific origins
		allowedOrigin := "http://127.0.0.1:5173"
		if r.Header.Get("Origin") == allowedOrigin {
			return true
		}
		return false
	},
}

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

func (c *Client) listen() {

}

func (c *Client) write() {

}
