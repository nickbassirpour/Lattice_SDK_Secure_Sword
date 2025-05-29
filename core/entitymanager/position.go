package entitymanager

import (
	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
)

func UpdatePosition(original_position *components.Position, new_position *components.Position) error {
	if new_position.LatitudeDegrees != nil {
		original_position.LatitudeDegrees = new_position.LatitudeDegrees
	}
	if new_position.LongitudeDegrees != nil {
		original_position.LongitudeDegrees = new_position.LongitudeDegrees
	}
	if new_position.AltitudeHaeMeters != nil {
		original_position.AltitudeHaeMeters = new_position.AltitudeHaeMeters
	}
	if new_position.AltitudeAglMeters != nil {
		original_position.AltitudeAglMeters = new_position.AltitudeAglMeters
	}
	if new_position.AltitudeAsfMeters != nil {
		original_position.AltitudeAsfMeters = new_position.AltitudeAsfMeters
	}
	if new_position.PressureDepthMeters != nil {
		original_position.PressureDepthMeters = new_position.PressureDepthMeters
	}
	return nil
}
