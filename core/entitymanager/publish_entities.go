package entitymanager

import (
	"io"
	"log"

	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
)

func (s *Server) PublishEntities(stream components.EntityManager_PublishEntitiesServer) error {
	for {
		entity, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&components.PublishEntitiesResponse{
				Success: true,
				Message: "All entities published",
			})
		}
		if err != nil {
			return err
		}

		log.Printf("Received request: %+v", entity)

		if entity.EntityId == nil {
			CreateEntity(entity, s)
		} else {
			UpdateEntity(entity, s)
		}
	}
}
