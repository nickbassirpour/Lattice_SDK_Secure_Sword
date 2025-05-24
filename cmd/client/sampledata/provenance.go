package sampledata

import (
	"time"

	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func SampleProvenance() *components.Provenance {
	return &components.Provenance{
		IntegrationName:  "sample-integration",
		DataType:         "ADSB",
		SourceUpdateTime: timestamppb.New(time.Now()),
	}
}
