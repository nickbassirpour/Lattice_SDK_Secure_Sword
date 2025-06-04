package sampledata

import (
	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
)

func SampleGeoShape() *components.GeoShape {
	return &components.GeoShape{
		Point:     &components.Point{Position: SamplePosition()},
		Line:      &components.Line{Position: SamplePosition()},
		Polygon:   SamplePolygon(),
		Ellipse:   SampleEllise(),
		Ellipsoid: SampleEllipsoid(),
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
		&components.LinearRing{Position: SamplePosition(), HeightM: Pointer(float32(9032.235))},
	}
	ring := &components.Ring{
		Positions: linear_rings,
	}
	rings = append(rings, ring)
	return rings
}

func SampleEllise() *components.Ellipse {
	return &components.Ellipse{
		SemiMajorAxisM: Pointer(float64(236.31)),
		SemiMinorAxisM: Pointer(float64(765.31)),
		OrientationD:   Pointer(float64(43.34)),
		HeightM:        Pointer(float64(432.31)),
	}
}

func SampleEllipsoid() *components.Ellipsoid {
	return &components.Ellipsoid{
		ForwardAxisM: Pointer(float64(2398.35)),
		SideAxisM:    Pointer(float64(1398.35)),
		UpAxisM:      Pointer(float64(238.35)),
	}
}
