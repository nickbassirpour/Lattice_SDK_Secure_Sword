package entitymanager

import (
	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
)

func UpdatePosition(entity *components.Entity, new_position *components.Position) error {
	if new_position.LatitudeDegrees != nil {
		entity.Location.Position.LatitudeDegrees = new_position.LatitudeDegrees
	}
	if new_position.LongitudeDegrees != nil {
		entity.Location.Position.LongitudeDegrees = new_position.LongitudeDegrees
	}
	if new_position.AltitudeHaeMeters != nil {
		entity.Location.Position.AltitudeHaeMeters = new_position.AltitudeHaeMeters
	}
	if new_position.AltitudeAglMeters != nil {
		entity.Location.Position.AltitudeAglMeters = new_position.AltitudeAglMeters
	}
	if new_position.AltitudeAsfMeters != nil {
		entity.Location.Position.AltitudeAsfMeters = new_position.AltitudeAsfMeters
	}
	if new_position.PressureDepthMeters != nil {
		entity.Location.Position.PressureDepthMeters = new_position.PressureDepthMeters
	}
	return nil
}
