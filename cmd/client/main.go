package main

import (
	"context"
	"log"
	"time"

	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.PublishEntity(ctx, &components.PublishEntityRequest{Entity: sampledata.SampleEntity()})
	if err != nil {
		log.Fatalf("could not publish entity: %v", err)
	}

	log.Printf("Response: %v", resp)
}
