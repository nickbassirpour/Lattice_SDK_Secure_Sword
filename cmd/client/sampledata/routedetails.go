package sampledata

import (
	"time"

	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func SampleRouteDetails() *components.RouteDetails {
	return &components.RouteDetails{
		DestinationName:      Pointer("Los Angeles"),
		EstimatedArrivalTime: timestamppb.New(time.Now()),
	}
}
