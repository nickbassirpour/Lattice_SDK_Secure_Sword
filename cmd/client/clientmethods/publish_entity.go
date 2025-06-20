package clientmethods

import (
	"context"
	"fmt"
	"log"
	"time"

	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
)

func SendEntityRequest(client components.EntityManagerClient, entity *components.Entity) {
	req := &components.PublishEntityRequest{Entity: entity}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.PublishEntity(ctx, req)
	if err != nil {
		log.Fatalf("could not publish entity: %v", err)
	}

	fmt.Printf("Response: %v\n", resp)

}
