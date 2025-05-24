package entitymanager

import (
	"sync"

	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
)

type Server struct {
	components.UnimplementedEntityManagerServer
	mu       sync.Mutex
	entities map[string]*components.Entity
}

func NewServer() *Server {
	return &Server{
		entities: make(map[string]*components.Entity),
	}
}
