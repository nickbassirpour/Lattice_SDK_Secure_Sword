package sampledata

import (
	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
)

func SampleGroupDetails() *components.GroupDetails {
	return &components.GroupDetails{
		Team:    nil,
		Echelon: sampleEchelon(),
	}
}

func sampleEchelon() *components.Echelon {
	return &components.Echelon{
		ArmyEchelon: Pointer(components.ArmyEchelon_ARMY_ECHELON_ARMY),
	}
}
