package entitymanager

import (
	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
)

func UpdateAliases(entity *components.Entity, new_aliases *components.Aliases) error {
	if new_aliases.Name != nil {
		entity.Aliases.Name = new_aliases.Name
	}
	if new_aliases.AlternateIds != nil {
		if len(new_aliases.AlternateIds) > 0 {
			for i := 0; i < len(new_aliases.AlternateIds); i++ {
				entity.Aliases.AlternateIds = append(entity.Aliases.AlternateIds, new_aliases.AlternateIds[i])
			}
		}
	}
	return nil
}
