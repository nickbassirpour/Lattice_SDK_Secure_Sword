package entitymanager

import (
	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
)

func UpdateRouteDetails(entity *components.Entity, newRouteDetails *components.RouteDetails) error {
	if newRouteDetails.DestinationName != nil {
		entity.RouteDetails.DestinationName = newRouteDetails.DestinationName
	}
	if newRouteDetails.EstimatedArrivalTime != nil {
		entity.RouteDetails.EstimatedArrivalTime = newRouteDetails.EstimatedArrivalTime
	}
	return nil
}
