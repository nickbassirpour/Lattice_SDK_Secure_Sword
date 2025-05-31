package entitymanager

import (
	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
)

func UpdateGeoDetails(entity *components.Entity, new_geo_details *components.GeoDetails) error {
	if new_geo_details.Type != nil {
		entity.GeoDetails.Type = new_geo_details.Type
	}
	if new_geo_details.ControlArea != nil && new_geo_details.ControlArea.Type != nil {
		entity.GeoDetails.ControlArea.Type = new_geo_details.ControlArea.Type
	}
	if new_geo_details.Acm != nil {
		if new_geo_details.Acm.AcmType != nil {
			entity.GeoDetails.Acm.AcmType = new_geo_details.Acm.AcmType
		}
		if new_geo_details.Acm.AcmDescription != nil {
			entity.GeoDetails.Acm.AcmDescription = new_geo_details.Acm.AcmDescription
		}
	}
	return nil
}

/*
package entities

// A component that describes a geo-entity.
type GeoDetails struct {
	Type        GeoType     `json:"type"`
	ControlArea ControlArea `json:"controlArea"`
	Acm         Acm         `json:"acm"`
}


// Determines the type of control area being represented by the geo-entity,
// in which an asset can, or cannot, operate.
type ControlArea struct {
	Type ControlAreaType `json:"type"`
}

type ControlAreaType int

const (
	CONTROL_AREA_TYPE_INVALID ControlAreaType = iota
	CONTROL_AREA_TYPE_KEEP_IN_ZONE
	CONTROL_AREA_TYPE_KEEP_OUT_ZONE
	CONTROL_AREA_TYPE_DITCH_ZONE
	CONTROL_AREA_TYPE_LOITER_ZONE
)

type Acm struct {
	AcmType AcmType `json:"acmType"`

	// Used for loosely typed associations, such as assignment to a specific fires unit. Limit to 150 characters.
	AcmDescription string `json:"acmDescription"`
}

type AcmType int

const (
	ACM_DETAIL_TYPE_INVALID AcmType = iota
	ACM_DETAIL_TYPE_LANDING_ZONE
)

*/
