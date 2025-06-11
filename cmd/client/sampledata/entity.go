package sampledata

import (
	"time"

	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
	"google.golang.org/protobuf/types/known/anypb"
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
		Overrides:           SampleOverrides(),
		Indicators:          SampleIndicators(),
		TargetPriority:      SampleTargetPriority(),
		Signal:              SampleSignal(),
		TransponderCodes:    SampleTransponderCodes(),
		DataClassification:  SampleDataClassification(),
		TaskCatalog:         SampleTaskCatalog(),
	}
}

func Pointer[T any](pointer T) *T {
	return &pointer
}

func WrapEntityAsAny(entity *components.Entity) (*anypb.Any, error) {
	entity_set_to_any, err := anypb.New(entity)
	if err != nil {
		return nil, err
	}
	return entity_set_to_any, nil
}
