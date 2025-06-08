package sampledata

import (
	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
)

func SampleIndicators() *components.Indicators {
	return &components.Indicators{
		Simulated:  Pointer(true),
		Exercise:   Pointer(false),
		Emergency:  Pointer(true),
		C2:         Pointer(false),
		Egressable: Pointer(false),
		Starred:    Pointer(true),
	}
}
