package main

import (
	"log"

	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
	clientmethods "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/cmd/client/clientmethods"
	sampledata "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/cmd/client/sampledata"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50050", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := components.NewEntityManagerClient(conn)
	entity := sampledata.SampleEntity()

	clientmethods.SendEntityRequest(client, entity)

}
