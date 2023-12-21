package main

import "log"

type RoomManager struct {
	roomID      string
	activeUsers map[*Client]bool
	server      *Server
	join        chan *Client
	leave       chan *Client
	msgqueue    chan string
}

func (rm *RoomManager) broadcast() {
	defer func() {
		rm.server.destroy <- rm
	}()
	for {
		select {
		case client, ok := <-rm.join:
			if !ok {
				log.Println("RoomManager (ID: %s) Join Queue Closed")
				return
			}
			rm.activeUsers[client] = true
		case client, ok := <-rm.leave:
			if !ok {
				log.Println("RoomManager (ID: %s) Leave Queue Closed")
				return
			}
			delete(rm.activeUsers, client)
			close(client.msgqueue)
			if len(rm.activeUsers) == 0 {
				log.Println("RoomManager (ID: %s) Closed Due To 0 Users")
				return
			}
		case message, ok := <-rm.msgqueue:
			if !ok {
				log.Println("RoomManager (ID: %s) Message Queue Closed")
				return
			}
			for user := range rm.activeUsers {
				user.msgqueue <- message
			}
		}
	}

}
