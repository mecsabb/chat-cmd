package main

import (
	"net/http"
)

// handleCreate - i.e. backend response to client call - chat create
//   - this needs to print a website link (e.g. localhost:XXXX) that the user can put into
//     their browser to view the chatroom (how does this work?) -- this is client code --
//   - this also will generate the key
func handleCreate(w http.ResponseWriter, r *http.Request) {

}

// handleJoin - i.e. backend response to client call - chat join -k <key>
//   - create a Client struct for the client information
//   - pass the client into the join channel of the roommanager corresponding to the client's key
func handleJoin(w http.ResponseWriter, r *http.Request) {

}

func main() {
	// TODO
	// need to handle the response to a new user joining
	server := &Server{
		connections: make(map[string]*RoomManager),
		create:      make(chan *RoomManager),
		destroy:     make(chan *RoomManager),
	}

	go server.spin()

	mux := http.NewServeMux()

	mux.Handle("/create", http.HandlerFunc(handleCreate))
	mux.Handle("/join", http.HandlerFunc(handleJoin))
}
