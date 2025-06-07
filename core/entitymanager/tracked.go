package entitymanager

import (
	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
)

func UpdateTracked(entity *components.Entity, new_tracked *components.Tracked) error {
	if new_tracked.TrackQualityWrapper != nil {
		entity.Tracked.TrackQualityWrapper = new_tracked.TrackQualityWrapper
	}
	if new_tracked.SensorHits != nil {
		entity.Tracked.SensorHits = new_tracked.SensorHits
	}
	if new_tracked.NumberOfObjects != nil {
		err := UpdateNumberOfObjects(entity, new_tracked)
		if err != nil {
			return err
		}
	}
	if new_tracked.RadarCrossSection != nil {
		entity.Tracked.RadarCrossSection = new_tracked.RadarCrossSection
	}
	if new_tracked.LastMeasurementTime != nil {
		entity.Tracked.LastMeasurementTime = new_tracked.LastMeasurementTime
	}
	if new_tracked.LineOfBearing != nil {
		err := UpdateLineOfBearing(entity, new_tracked)
		if err != nil {
			return err
		}
	}
	return nil
}

func UpdateNumberOfObjects(entity *components.Entity, new_tracked *components.Tracked) error {
	if new_tracked.NumberOfObjects.LowerBound != nil {
		entity.Tracked.NumberOfObjects.LowerBound = new_tracked.NumberOfObjects.LowerBound
	}
	if new_tracked.NumberOfObjects.UpperBound != nil {
		entity.Tracked.NumberOfObjects.UpperBound = new_tracked.NumberOfObjects.UpperBound
	}
	return nil
	// Write logic to define strength calculation with upper/lower data
}

func UpdateLineOfBearing(entity *components.Entity, new_tracked *components.Tracked) error {
	if new_tracked.LineOfBearing.AngleOfArrival != nil {
		err := UpdateAngleOfArrival(entity, new_tracked)
		if err != nil {
			return err
		}
	}
	if new_tracked.LineOfBearing.RangeEstimateM != nil {
		err := UpdateRangeEstimateM(entity, new_tracked)
		if err != nil {
			return err
		}
	}
	if new_tracked.LineOfBearing.MaxRangeM != nil {
		err := UpdateMaxRangeM(entity, new_tracked)
		if err != nil {
			return err
		}
	}
	return nil
}

func UpdateAngleOfArrival(entity *components.Entity, new_tracked *components.Tracked) error {
	new_angle_of_arrival := new_tracked.LineOfBearing.AngleOfArrival
	existing_angle_of_arrival := entity.Tracked.LineOfBearing.AngleOfArrival
	if new_angle_of_arrival.RelativePose != nil {
		err := UpdateRelativePose(existing_angle_of_arrival.RelativePose, new_angle_of_arrival.RelativePose)
		if err != nil {
			return err
		}
	}
	if new_angle_of_arrival.BearingElevationCovarianceRad2 != nil {
		if new_angle_of_arrival.BearingElevationCovarianceRad2.Mxx != nil {
			entity.Tracked.LineOfBearing.AngleOfArrival.BearingElevationCovarianceRad2.Mxx = new_angle_of_arrival.BearingElevationCovarianceRad2.Mxx
		}
		if new_angle_of_arrival.BearingElevationCovarianceRad2.Mxy != nil {
			entity.Tracked.LineOfBearing.AngleOfArrival.BearingElevationCovarianceRad2.Mxy = new_angle_of_arrival.BearingElevationCovarianceRad2.Mxy
		}
		if new_angle_of_arrival.BearingElevationCovarianceRad2.Myy != nil {
			entity.Tracked.LineOfBearing.AngleOfArrival.BearingElevationCovarianceRad2.Myy = new_angle_of_arrival.BearingElevationCovarianceRad2.Myy
		}
	}
	return nil
}

func UpdateRelativePose(existing_rel_pose *components.RelativePose, new_rel_pose *components.RelativePose) error {
	if new_rel_pose.Pos != nil {
		if new_rel_pose.Pos.Lon != nil {
			existing_rel_pose.Pos.Lon = new_rel_pose.Pos.Lon
		}
		if new_rel_pose.Pos.Lat != nil {
			existing_rel_pose.Pos.Lat = new_rel_pose.Pos.Lat
		}
		if new_rel_pose.Pos.Alt != nil {
			existing_rel_pose.Pos.Alt = new_rel_pose.Pos.Alt
		}
		if new_rel_pose.Pos.Is2D != nil {
			existing_rel_pose.Pos.Is2D = new_rel_pose.Pos.Is2D
		}
		if new_rel_pose.Pos.AltitudeReference != nil {
			existing_rel_pose.Pos.AltitudeReference = new_rel_pose.Pos.AltitudeReference
		}
	}
	if new_rel_pose.AttEnu != nil {
		if new_rel_pose.AttEnu.X != nil {
			existing_rel_pose.AttEnu.X = new_rel_pose.AttEnu.X
		}
		if new_rel_pose.AttEnu.Y != nil {
			existing_rel_pose.AttEnu.Y = new_rel_pose.AttEnu.Y
		}
		if new_rel_pose.AttEnu.Z != nil {
			existing_rel_pose.AttEnu.Z = new_rel_pose.AttEnu.Z
		}
		if new_rel_pose.AttEnu.W != nil {
			existing_rel_pose.AttEnu.W = new_rel_pose.AttEnu.W
		}
	}
	return nil
}

func UpdateRangeEstimateM(entity *components.Entity, new_tracked *components.Tracked) error {
	range_estimate := new_tracked.LineOfBearing.RangeEstimateM
	if range_estimate.Value != nil {
		entity.Tracked.LineOfBearing.RangeEstimateM.Value = range_estimate.Value
	}
	if range_estimate.Sigma != nil {
		entity.Tracked.LineOfBearing.RangeEstimateM.Sigma = range_estimate.Sigma
	}
	return nil
}

func UpdateMaxRangeM(entity *components.Entity, new_tracked *components.Tracked) error {
	max_range_m := new_tracked.LineOfBearing.MaxRangeM
	if max_range_m.Value != nil {
		entity.Tracked.LineOfBearing.MaxRangeM.Value = max_range_m.Value
	}
	if max_range_m.Sigma != nil {
		entity.Tracked.LineOfBearing.MaxRangeM.Sigma = max_range_m.Sigma
	}
	return nil
}
