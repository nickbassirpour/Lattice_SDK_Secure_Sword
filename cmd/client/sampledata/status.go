package sampledata

import components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"

func SampleStatus() *components.Status {
	return &components.Status{
		Code:    Pointer(int32(0)),
		Message: Pointer("Entity created"),
	}
}
