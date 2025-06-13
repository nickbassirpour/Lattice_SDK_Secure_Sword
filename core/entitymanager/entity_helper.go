package entitymanager

import (
	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
)

func UpdateComponents(entity *components.Entity, newData *components.Entity) (*components.Entity, error) {
	if newData.Description != nil {
		entity.Description = newData.Description
	}
	if newData.IsLive != nil {
		entity.IsLive = newData.IsLive
	}
	if newData.CreatedTime != nil {
		entity.CreatedTime = newData.CreatedTime
	}
	if newData.ExpiryTime != nil {
		entity.ExpiryTime = newData.ExpiryTime
	}
	if newData.Status != nil {
		err := UpdateStatus(entity, newData.Status)
		if err != nil {
			return nil, err
		}
	}
	if newData.Location != nil {
		err := UpdateLocation(entity, newData.Location)
		if err != nil {
			return nil, err
		}
	}
	if newData.LocationUncertainty != nil {
		err := UpdateLocationUncertainty(entity, newData.LocationUncertainty)
		if err != nil {
			return nil, err
		}
	}
	if newData.GeoShape != nil {
		err := UpdateGeoShape(entity, newData.GeoShape)
		if err != nil {
			return nil, err
		}
	}
	if newData.GeoDetails != nil {
		err := UpdateGeoDetails(entity, newData.GeoDetails)
		if err != nil {
			return nil, err
		}
	}
	if newData.Aliases != nil {
		err := UpdateAliases(entity, newData.Aliases)
		if err != nil {
			return nil, err
		}
	}
	if newData.Tracked != nil {
		err := UpdateTracked(entity, newData.Tracked)
		if err != nil {
			return nil, err
		}
	}
	if newData.Correlation != nil {
		err := UpdateCorrelation(entity, newData.Correlation)
		if err != nil {
			return nil, err
		}
	}
	if newData.MilView != nil {
		err := UpdateMilView(entity, newData.MilView)
		if err != nil {
			return nil, err
		}
	}
	if newData.Ontology != nil {
		err := UpdateOntology(entity, newData.Ontology)
		if err != nil {
			return nil, err
		}
	}
	if newData.Sensors != nil {
		newSensors, err := UpdateSensors(entity.Sensors.Sensors, newData.Sensors.Sensors)
		if err != nil {
			return nil, err
		}
		entity.Sensors.Sensors = newSensors
	}
	if newData.Payloads != nil {
		err := UpdatePayloads(entity, newData.Payloads)
		if err != nil {
			return nil, err
		}
	}
	if newData.PowerState != nil {
		err := UpdatePowerState(entity, newData.PowerState)
		if err != nil {
			return nil, err
		}
	}
	if newData.Overrides != nil {
		err := UpdateOverrides(entity, newData.Overrides)
		if err != nil {
			return nil, err
		}
	}
	if newData.TargetPriority != nil {
		err := UpdateTargetPriority(entity, newData.TargetPriority)
		if err != nil {
			return nil, err
		}
	}
	if newData.Signal != nil {
		err := UpdateSignal(entity, newData.Signal)
		if err != nil {
			return nil, err
		}
	}
	if newData.TransponderCodes != nil {
		err := UpdateTransponderCodes(entity, newData.TransponderCodes)
		if err != nil {
			return nil, err
		}
	}
	if newData.DataClassification != nil {
		err := UpdateDataClassification(entity, newData.DataClassification)
		if err != nil {
			return nil, err
		}
	}
	if newData.TaskCatalog != nil {
		err := UpdateTaskCatalog(entity, newData.TaskCatalog)
		if err != nil {
			return nil, err
		}
	}
	if newData.Relationships != nil {
		err := UpdateRelationships(entity, newData.Relationships)
		if err != nil {
			return nil, err
		}
	}
	if newData.VisualDetails != nil {
		err := UpdateVisualDetails(entity, newData.VisualDetails)
		if err != nil {
			return nil, err
		}
	}
	if newData.Dimensions != nil {
		err := UpdateDimensions(entity, newData.Dimensions)
		if err != nil {
			return nil, err
		}
	}
	if newData.RouteDetails != nil {
		err := UpdateRouteDetails(entity, newData.RouteDetails)
		if err != nil {
			return nil, err
		}
	}
	if newData.Schedules != nil {
		err := UpdateSchedules(entity, newData.Schedules)
		if err != nil {
			return nil, err
		}
	}
	if newData.Health != nil {
		err := UpdateHealth(entity, newData.Health)
		if err != nil {
			return nil, err
		}
	}
	if newData.GroupDetails != nil {
		err := UpdateGroupDetails(entity, newData.GroupDetails)
		if err != nil {
			return nil, err
		}
	}
	return entity, nil
}

/*

Supplies Supplies `json:"supplies"`

Orbit Orbit `json:"orbit"`
*/

func Pointer[T any](pointer T) *T {
	return &pointer
}
