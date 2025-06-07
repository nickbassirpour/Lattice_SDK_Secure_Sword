package entitymanager

import (
	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
)

func UpdateLocation(entity *components.Entity, new_location *components.Location) error {
	if new_location.Position != nil {
		err := UpdatePosition(entity.Location.Position, new_location.Position)
		if err != nil {
			return err
		}
		// Add check for location and match with country borders
		// once found, log and send report if in allied territory
		// if in specific fleet (3rd fleet, 5th fleet, etc.) tell Navy
		// space := checkAirspaceOrWaterspace(new_location.Position)
	}

	if new_location.VelocityEnu != nil {
		if new_location.VelocityEnu.E != nil {
			entity.Location.VelocityEnu.E = new_location.VelocityEnu.E
		}
		if new_location.VelocityEnu.N != nil {
			entity.Location.VelocityEnu.N = new_location.VelocityEnu.N
		}
		if new_location.VelocityEnu.U != nil {
			entity.Location.VelocityEnu.U = new_location.VelocityEnu.U
		}
	}

	if new_location.SpeedMps != nil {
		entity.Location.SpeedMps = new_location.SpeedMps
	}

	if new_location.Acceleration != nil {
		if new_location.Acceleration.E != nil {
			entity.Location.Acceleration.E = new_location.Acceleration.E
		}

		if new_location.Acceleration.N != nil {
			entity.Location.Acceleration.N = new_location.Acceleration.N
		}

		if new_location.Acceleration.U != nil {
			entity.Location.Acceleration.U = new_location.Acceleration.U
		}
	}

	if new_location.AttitudeEnu != nil {
		if new_location.AttitudeEnu.X != nil {
			entity.Location.AttitudeEnu.X = new_location.AttitudeEnu.X
		}
		if new_location.AttitudeEnu.Y != nil {
			entity.Location.AttitudeEnu.Y = new_location.AttitudeEnu.Y
		}
		if new_location.AttitudeEnu.Z != nil {
			entity.Location.AttitudeEnu.Z = new_location.AttitudeEnu.Z
		}
		if new_location.AttitudeEnu.W != nil {
			entity.Location.AttitudeEnu.W = new_location.AttitudeEnu.W
		}
	}

	return nil
}
