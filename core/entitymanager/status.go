package entitymanager

import (
	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
)

func UpdateStatus(original_status *components.Status, new_status *components.Status) error {
	if new_status.Code != nil {
		original_status.Code = new_status.Code
	}
	if new_status.Message != nil {
		original_status.Message = new_status.Message
	}
	if len(new_status.Details) > 0 {
		original_status.Details = new_status.Details
	}
	return nil
}
