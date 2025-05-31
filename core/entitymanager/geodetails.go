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
