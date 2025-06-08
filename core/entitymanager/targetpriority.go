package entitymanager

import (
	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
)

func UpdateTargetPriority(entity *components.Entity, new_target_priority *components.TargetPriority) error {
	if new_target_priority.HighValueTarget != nil {
		new_hvt := new_target_priority.HighValueTarget
		existing_hvt := entity.TargetPriority.HighValueTarget
		if new_hvt.IsHighValueTarget != nil {
			existing_hvt.IsHighValueTarget = new_hvt.IsHighValueTarget
		}
		if new_hvt.TargetPriority != nil {
			existing_hvt.TargetPriority = new_hvt.TargetPriority
		}
		if new_hvt.TargetMatches != nil && len(new_hvt.TargetMatches) > 0 {
			existing_hvt.TargetMatches = append(existing_hvt.TargetMatches, new_hvt.TargetMatches...)
		}
	}
	if new_target_priority.Threat != nil {
		entity.TargetPriority.Threat = new_target_priority.Threat
	}
	return nil
}
