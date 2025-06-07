package sampledata

import (
	"time"

	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func SampleEntity() *components.Entity {
	return &components.Entity{
		EntityId:            nil,
		Description:         Pointer("Sample entity for testing"),
		IsLive:              Pointer(true),
		CreatedTime:         timestamppb.New(time.Now()),
		ExpiryTime:          timestamppb.New(time.Now().Add(48 * time.Hour)),
		Status:              SampleStatus(),
		Location:            SampleLocation(),
		Provenance:          SampleProvenance(),
		Aliases:             SampleAliases(),
		Ontology:            SampleOntology(),
		GeoDetails:          SampleGeoDetails(),
		GeoShape:            SampleGeoShape(),
		LocationUncertainty: SampleLocationUncertainty(),
		Tracked:             SampleTracked(),
		Correlation:         SampleCorrelation(),
		MilView:             SampleMilView(),
		Sensors:             SampleSensors(),
		Payloads:            SamplePayloads(),
		PowerState:          SamplePowerState(),
	}
}

func Pointer[T any](pointer T) *T {
	return &pointer
}
