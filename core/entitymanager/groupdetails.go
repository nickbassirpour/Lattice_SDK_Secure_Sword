package entitymanager

import (
	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
)

func UpdateGroupDetails(entity *components.Entity, newGroupDetails *components.GroupDetails) error {
	if newGroupDetails.Team != nil {
		entity.GroupDetails.Team = newGroupDetails.Team
	}
	if newGroupDetails.Echelon != nil {
		entity.GroupDetails.Echelon = newGroupDetails.Echelon
	}
	return nil
}
