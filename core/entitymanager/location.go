package entitymanager

import (
	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
)

func UpdateLocation(original_location *components.Location, new_location *components.Location) error {
	if new_location.Position != nil {
		err := UpdatePosition(original_location.Position, new_location.Position)
		if err != nil {
			return err
		}
		space := checkAirspaceOrWaterspace(new_location.Position)
	}

	if new_location.VelocityEnu != nil {
		original_location.VelocityEnu.E = new_location.VelocityEnu.E
		original_location.VelocityEnu.N = new_location.VelocityEnu.N
		original_location.VelocityEnu.U = new_location.VelocityEnu.U
	}

	if new_location.SpeedMps != nil {
		original_location.SpeedMps = new_location.SpeedMps
	}

	if new_location.Acceleration != nil {
		original_location.Acceleration.E = new_location.Acceleration.E
		original_location.Acceleration.N = new_location.Acceleration.N
		original_location.Acceleration.U = new_location.Acceleration.U
	}

	if new_location.AttitudeEnu != nil {
		original_location.AttitudeEnu.X = new_location.AttitudeEnu.X
		original_location.AttitudeEnu.Y = new_location.AttitudeEnu.Y
		original_location.AttitudeEnu.Z = new_location.AttitudeEnu.Z
		original_location.AttitudeEnu.W = new_location.AttitudeEnu.W
	}

	return nil
}
