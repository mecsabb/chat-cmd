package main

type RoomManager struct {
	roomID      string
	activeUsers map[*Client]bool
	server      *Server
}

func (rm *RoomManager) broadcast() {
	defer func() {
		rm.server.destroy <- rm
	}()

}
