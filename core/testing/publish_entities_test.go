package entitymanager

import (
	// "context"
	"log"
	"net"
	"testing"

	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
	entitymanager "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/core/entitymanager"
	"google.golang.org/grpc"
)

func TestPublishEntities(t *testing.T) {
	lis, err := net.Listen("tcp", ":50050")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	server := entitymanager.NewServer()
	components.RegisterEntityManagerServer(grpcServer, server)

	go grpcServer.Serve(lis)
	defer grpcServer.Stop()

	// sampleEntities := []*components.Entity{}
	// for i := 0; i < 10; i++ {
	// 	sampleEntities = append(sampleEntities, sampledata.SampleEntity())
	// }

}
