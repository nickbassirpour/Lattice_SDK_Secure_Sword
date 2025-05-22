package entitymanager

import (
	"context"
	"log"

	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
)

type Server struct {
	components.UnimplementedEntityManagerServer
}

func (s *Server) PublishEntity(ctx context.Context, req *components.PublishEntityRequest) (*components.PublishEntityResponse, error) {
	log.Printf("Received request: %+v", req.GetEntity())

	return &components.PublishEntityResponse{
		Success: true,
		Message: "Published Entity",
	}, nil
}
