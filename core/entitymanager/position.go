package entitymanager

import (
	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
)

func UpdatePosition(existing_position *components.Position, new_position *components.Position) error {
	if new_position.LatitudeDegrees != nil {
		existing_position.LatitudeDegrees = new_position.LatitudeDegrees
	}
	if new_position.LongitudeDegrees != nil {
		existing_position.LongitudeDegrees = new_position.LongitudeDegrees
	}
	if new_position.AltitudeHaeMeters != nil {
		existing_position.AltitudeHaeMeters = new_position.AltitudeHaeMeters
	}
	if new_position.AltitudeAglMeters != nil {
		existing_position.AltitudeAglMeters = new_position.AltitudeAglMeters
	}
	if new_position.AltitudeAsfMeters != nil {
		existing_position.AltitudeAsfMeters = new_position.AltitudeAsfMeters
	}
	if new_position.PressureDepthMeters != nil {
		existing_position.PressureDepthMeters = new_position.PressureDepthMeters
	}
	return nil
}
