package sampledata

import components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"

func SamplePosition() *components.Position {
	return &components.Position{
		LatitudeDegrees:   Pointer(37.7749),
		LongitudeDegrees:  Pointer(-122.4194),
		AltitudeAglMeters: Pointer(1256.34),
	}
}
