package entitymanager

import (
	"context"
	"errors"
	"log"

	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
	"github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/core/validation"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	components.UnimplementedEntityManagerServer
}

func (s *Server) PublishEntity(ctx context.Context, req *components.PublishEntityRequest) (*components.PublishEntityResponse, error) {
	if req == nil || req.GetEntity() == nil {
		return nil, errors.New("PublishEntityRequest or Entity is nil")
	}

	log.Printf("Received request: %+v", req.GetEntity())
	if err := validation.ValidateEntity(req.GetEntity()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validation failed: %v", err)
	}

	return &components.PublishEntityResponse{
		Success: true,
		Message: "Published Entity",
	}, nil
}
