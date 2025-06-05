package entitymanager

import (
	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
)

func UpdateMilView(entity *components.Entity, new_milview *components.MilView) error {
	if new_milview.Disposition != nil {
		entity.MilView.Disposition = new_milview.Disposition
	}
	if new_milview.Environment != nil {
		entity.MilView.Environment = new_milview.Environment
	}
	if new_milview.Nationality != nil {
		entity.MilView.Nationality = new_milview.Nationality
	}
	return nil
}
