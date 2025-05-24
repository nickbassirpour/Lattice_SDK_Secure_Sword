package main

import (
	"log"
	"net"

	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
	entitymanager "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/core/entitymanager"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50050")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	server := entitymanager.NewServer()
	components.RegisterEntityManagerServer(grpcServer, server)

	log.Println("gRPC server listening on :50050")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
