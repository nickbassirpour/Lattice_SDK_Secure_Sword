package entitymanager

import (
	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
)

func UpdateStatus(entity *components.Entity, new_status *components.Status) error {
	if new_status.Code != nil {
		entity.Status.Code = new_status.Code
	}
	if new_status.Message != nil {
		entity.Status.Message = new_status.Message
	}
	if len(new_status.Details) > 0 {
		entity.Status.Details = new_status.Details
	}
	return nil
}
