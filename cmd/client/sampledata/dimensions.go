package sampledata

import (
	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
)

func SampleDimensions() *components.Dimensions {
	return &components.Dimensions{
		LengthM: Pointer(float32(90.7)),
	}
}
