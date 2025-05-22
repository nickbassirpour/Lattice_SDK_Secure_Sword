package sampledata

import components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"

func SampleStatus() *components.Status {
	return &components.Status{
		Code:    0,
		Message: "Entity created",
	}
}
