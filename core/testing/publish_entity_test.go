package entitymanager

import (
	"context"
	"testing"

	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
	sampledata "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/cmd/client/sampledata"
	entitymanager "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/core/entitymanager"
)

func TestPublishEntity(t *testing.T) {
	server := entitymanager.NewServer()
	sample_entity := sampledata.SampleEntity()

	req := &components.PublishEntityRequest{
		Entity: sample_entity,
	}

	resp, err := server.PublishEntity(context.Background(), req)

	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	if !resp.Success {
		t.Errorf("Expected Success, got error: %v", resp.Success)
	}

	if resp.Message != "Published Entity" {
		t.Errorf("Unexpected message: %v", resp.Message)
	}
}

func TestUpdatePublishedEntity(t *testing.T) {
	server := entitymanager.NewServer()
	sample_entity := sampledata.SampleEntity()

	req := &components.PublishEntityRequest{
		Entity: sample_entity,
	}

	server.PublishEntity(context.Background(), req)

}

func TestPublishEntity_InvalidRequest(t *testing.T) {
	server := &entitymanager.Server{}

	_, err := server.PublishEntity(context.Background(), nil)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestPublishEntity_NilEntity(t *testing.T) {
	server := &entitymanager.Server{}

	req := &components.PublishEntityRequest{Entity: nil}
	_, err := server.PublishEntity(context.Background(), req)
	if err == nil {
		t.Fatal("expected error for nil entity, got nil")
	}
}
