package entitymanager

import (
	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
)

func UpdateComponents(entity *components.Entity, new_data *components.Entity) (*components.Entity, error) {
	if new_data.Description != nil {
		entity.Description = new_data.Description
	}
	if new_data.IsLive != nil {
		entity.IsLive = new_data.IsLive
	}
	if new_data.CreatedTime != nil {
		entity.CreatedTime = new_data.CreatedTime
	}
	if new_data.ExpiryTime != nil {
		entity.ExpiryTime = new_data.ExpiryTime
	}
	if new_data.Status != nil {
		err := UpdateStatus(entity, new_data.Status)
		if err != nil {
			return nil, err
		}
	}
	if new_data.Location != nil {
		err := UpdateLocation(entity, new_data.Location)
		if err != nil {
			return nil, err
		}
	}
	if new_data.LocationUncertainty != nil {
		err := UpdateLocationUncertainty(entity, new_data.LocationUncertainty)
		if err != nil {
			return nil, err
		}
	}
	if new_data.GeoShape != nil {
		err := UpdateGeoShape(entity, new_data.GeoShape)
		if err != nil {
			return nil, err
		}
	}
	if new_data.GeoDetails != nil {
		err := UpdateGeoDetails(entity, new_data.GeoDetails)
		if err != nil {
			return nil, err
		}
	}
	if new_data.Aliases != nil {
		err := UpdateAliases(entity, new_data.Aliases)
		if err != nil {
			return nil, err
		}
	}
	if new_data.Tracked != nil {
		err := UpdateTracked(entity, new_data.Tracked)
		if err != nil {
			return nil, err
		}
	}
	if new_data.Correlation != nil {
		err := UpdateCorrelation(entity, new_data.Correlation)
		if err != nil {
			return nil, err
		}
	}
	if new_data.MilView != nil {
		err := UpdateMilView(entity, new_data.MilView)
		if err != nil {
			return nil, err
		}
	}
	if new_data.Ontology != nil {
		err := UpdateOntology(entity, new_data.Ontology)
		if err != nil {
			return nil, err
		}
	}
	if new_data.Sensors != nil {
		err := UpdateSensors(entity, new_data.Sensors.Sensors)
		if err != nil {
			return nil, err
		}
	}
	if new_data.Payloads != nil {
		err := UpdatePayloads(entity, new_data.Payloads)
		if err != nil {
			return nil, err
		}
	}
	if new_data.PowerState != nil {
		err := UpdatePowerState(entity, new_data.PowerState)
		if err != nil {
			return nil, err
		}
	}
	return entity, nil
}

/*
	PowerState PowerState `json:"powerState"`

	Overrides Overrides `json:"overrides"`

	Indicators Indicators `json:"indicators"`

	TargetPriority TargetPriority `json:"targetPriority"`

	Signal Signal `json:"signal"`

	TransponderCodes TransponderCodes `json:"transponderCodes"`

	DataClassification DataClassification `json:"dataClassification"`

	TaskCatalog TaskCatalog `json:"taskCatalog"`

	Relationships Relationships `json:"relationships"`

	VisualDetails VisualDetails `json:"visualDetails"`

	Dimensions Dimensions `json:"dimensions"`

	RouteDetails RouteDetails `json:"routeDetails"`

	Schedules Schedules `json:"schedules"`

	Health Health `json:"health"`

	GroupDetails GroupDetails `json:"groupDetails"`

	Supplies Supplies `json:"supplies"`

	Orbit Orbit `json:"orbit"`

*/

func Pointer[T any](pointer T) *T {
	return &pointer
}
