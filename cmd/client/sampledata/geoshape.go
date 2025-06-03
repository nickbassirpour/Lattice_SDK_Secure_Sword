package sampledata

import (
	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
)

func SampleGeoShape() *components.GeoShape {
	return &components.GeoShape{
		Point: SamplePosition(),
		Line: SamplePosition(),
		Polygon: SamplePolygon(),
		// Ellipse: ,
		// Ellipsoid: ,
	}
}

func SamplePolygon() *components.Polygon {
	return &components.Polygon{
		Rings: SampleRings(),
	}
}

func SampleRings() []*components.Ring {
	rings := []*components.Ring{}
	linear_rings := []*components.LinearRing{
		&components.LinearRing{Position: SamplePosition(), HeightM: Pointer(float32(9032.235))}
	}
	rings.append(linear_rings)
	return rings
}
