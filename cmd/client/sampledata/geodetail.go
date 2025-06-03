package sampledata

import (
	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
)

func SampleGeoDetails() *components.GeoDetails {
	return &components.GeoDetails{
		Type:        Pointer(components.GeoType(0)),
		ControlArea: &components.ControlArea{Type: components.ControlAreaType_CONTROL_AREA_TYPE_KEEP_IN_ZONE.Enum()},
		Acm:         &components.Acm{AcmType: components.AcmType_ACM_DETAIL_TYPE_INVALID.Enum(), AcmDescription: Pointer("Sample assignment")},
	}
}
