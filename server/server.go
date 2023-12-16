package main

type Server struct {
	connections map[string]*RoomManager
	create      chan *RoomManager
	destroy     chan *RoomManager
}

func (s *Server) spin() {
	for {
		select {
		case room := <-s.create:
			s.connections[room.roomID] = room
		case room := <-s.destroy:
			if _, ok := s.connections[room.roomID]; ok {
				delete(s.connections, room.roomID)
			}
		}
	}
}

func (s *Server) createRoom() {
	// this needs to handle the response to /createRoom endpoitn
}
