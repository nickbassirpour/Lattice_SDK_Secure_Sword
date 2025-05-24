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

func (s *Server) PublishEntity(ctx context.Context, req *components.PublishEntityRequest) (*components.PublishEntityResponse, error) {
	if req == nil || req.GetEntity() == nil {
		return nil, errors.New("PublishEntityRequest or Entity is nil")
	}

	entity := req.GetEntity()
	log.Printf("Received request: %+v", entity)

	if entity.EntityId == "" {
		return CreateEntity(entity, s)
	} else {
		UpdateEntity(entity, s)
	}
}

func CreateEntity(entity *components.Entity, s *Server) (*components.PublishEntityResponse, error) {
	if err := validation.ValidateEntity(entity); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validation failed: %v", err)
	}

	s.mu.Lock()
	s.entities[entity.EntityId] = entity
	s.mu.Unlock()

	return &components.PublishEntityResponse{
		Success: true,
		Message: "Published Entity",
	}, nil
}
