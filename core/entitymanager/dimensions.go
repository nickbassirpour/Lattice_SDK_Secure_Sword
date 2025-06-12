package entitymanager

import (
	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
)

func UpdateDimensions(entity *components.Entity, newDimensions *components.Dimensions) error {
	if newDimensions.LengthM != nil {
		entity.Dimensions.LengthM = newDimensions.LengthM
	}
	return nil
}
