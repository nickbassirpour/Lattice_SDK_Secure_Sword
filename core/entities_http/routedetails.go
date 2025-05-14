package entities

import "time"

// Additional information about an entity's route.
type RouteDetails struct {
	// Free form text giving the name of the entity's destination
	DestinationName string `json:"destinationName"`

	// Estimated time of arrival at destination
	EstimatedArrivalTime time.Time `json:"estimatedArrivalTime"`
}
