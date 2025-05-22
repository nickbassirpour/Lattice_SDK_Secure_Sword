package sampledata

import components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"

func SampleLocation() *components.Location {
	return &components.Location{
		Position: SamplePosition(),
	}
}
