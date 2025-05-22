package sampledata

import (
	"time"

	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func SampleEntity() *components.Entity {
	return &components.Entity{
		EntityId:    "sample-entity",
		Description: "Sample entity for testing",
		IsLive:      true,
		CreatedTime: timestamppb.New(time.Now()),
		ExpiryTime:  timestamppb.New(time.Now().Add(48 * time.Hour)),
		Status:      SampleStatus(),
		Location:    SampleLocation(),
	}
}
