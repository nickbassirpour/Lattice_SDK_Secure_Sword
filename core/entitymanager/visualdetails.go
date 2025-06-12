package entitymanager

import (
	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
)

func UpdateVisualDetails(entity *components.Entity, newVisualDetails *components.VisualDetails) error {
	if newVisualDetails.RangeRings != nil {
		newRangeRings := newVisualDetails.RangeRings
		if newRangeRings.MinDistanceM != nil {
			entity.VisualDetails.RangeRings.MinDistanceM = newRangeRings.MinDistanceM
		}
		if newRangeRings.MaxDistanceM != nil {
			entity.VisualDetails.RangeRings.MaxDistanceM = newRangeRings.MaxDistanceM
		}
		if newRangeRings.RingCount != nil {
			entity.VisualDetails.RangeRings.RingCount = newRangeRings.RingCount
		}
		if newRangeRings.RingLineColor != nil {
			existingRingLineColor := entity.VisualDetails.RangeRings.RingLineColor
			err := UpdateRingLingColor(existingRingLineColor, newRangeRings.RingLineColor)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func UpdateRingLingColor(existingRingLineColor *components.RingLineColor, newRingLingColor *components.RingLineColor) error {
	if newRingLingColor.Red != nil {
		existingRingLineColor.Red = newRingLingColor.Red
	}
	if newRingLingColor.Blue != nil {
		existingRingLineColor.Blue = newRingLingColor.Blue
	}
	if newRingLingColor.Green != nil {
		existingRingLineColor.Green = newRingLingColor.Green
	}
	if newRingLingColor.Alpha != nil {
		existingRingLineColor.Alpha = newRingLingColor.Alpha
	}
	return nil
}
