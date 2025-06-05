package sampledata

import (
	"time"

	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func SampleProvenance() *components.Provenance {
	return &components.Provenance{
		IntegrationName:   Pointer("sample-integration"),
		DataType:          Pointer("ADSB"),
		SourceUpdateTime:  timestamppb.New(time.Now()),
		SourceId:          Pointer("8235-dfg-25-dgf-2-2"),
		SourceDescription: Pointer("sample-provenance"),
	}
}
