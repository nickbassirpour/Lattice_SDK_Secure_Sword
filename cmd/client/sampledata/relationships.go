package sampledata

import (
	"time"

	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func SampleRelationships() *components.Relationships {
	relationships := []*components.Relationship{
		sampleRelationship(),
	}

	return &components.Relationships{
		Relationships: relationships,
	}
}

func sampleRelationship() *components.Relationship {
	return &components.Relationship{
		RelatedEntityId:  Pointer("gg-12098-skjfk"),
		RelationshipId:   Pointer("11-89083-sjsdf"),
		RelationshipType: sampleRelationshipType(),
	}
}

func sampleRelationshipType() *components.RelationshipType {
	return &components.RelationshipType{
		TrackedBy: sampleTrackedBy(),
	}
}

func sampleTrackedBy() *components.TrackedBy {
	return &components.TrackedBy{
		ActivelyTrackingSensors:  SampleSensors(),
		LastMeasurementTimeStamp: timestamppb.New(time.Now()),
	}
}
