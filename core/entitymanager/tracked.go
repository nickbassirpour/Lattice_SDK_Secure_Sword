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
		err := UpdateLineOfBearing(entity.Signal.LineOfBearing, new_tracked.LineOfBearing)
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

func UpdateLineOfBearing(existing_lob *components.LineOfBearing, new_lob *components.LineOfBearing) error {
	if new_lob.AngleOfArrival != nil {
		err := UpdateAngleOfArrival(existing_lob.AngleOfArrival, new_lob.AngleOfArrival)
		if err != nil {
			return err
		}
	}
	if new_lob.RangeEstimateM != nil {
		err := UpdateRangeEstimateM(existing_lob.RangeEstimateM, new_lob.RangeEstimateM)
		if err != nil {
			return err
		}
	}
	if new_lob.MaxRangeM != nil {
		err := UpdateMaxRangeM(existing_lob.MaxRangeM, new_lob.MaxRangeM)
		if err != nil {
			return err
		}
	}
	return nil
}

func UpdateAngleOfArrival(existing_angle_of_arr *components.AngleOfArrival, new_angle_of_arr *components.AngleOfArrival) error {
	if new_angle_of_arr.RelativePose != nil {
		err := UpdateRelativePose(existing_angle_of_arr.RelativePose, new_angle_of_arr.RelativePose)
		if err != nil {
			return err
		}
	}
	if new_angle_of_arr.BearingElevationCovarianceRad2 != nil {
		if new_angle_of_arr.BearingElevationCovarianceRad2.Mxx != nil {
			existing_angle_of_arr.BearingElevationCovarianceRad2.Mxx = new_angle_of_arr.BearingElevationCovarianceRad2.Mxx
		}
		if new_angle_of_arr.BearingElevationCovarianceRad2.Mxy != nil {
			existing_angle_of_arr.BearingElevationCovarianceRad2.Mxy = new_angle_of_arr.BearingElevationCovarianceRad2.Mxy
		}
		if new_angle_of_arr.BearingElevationCovarianceRad2.Myy != nil {
			existing_angle_of_arr.BearingElevationCovarianceRad2.Myy = new_angle_of_arr.BearingElevationCovarianceRad2.Myy
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

func UpdateRangeEstimateM(existing_range_est_m *components.RangeEstimateM, new_range_est_m *components.RangeEstimateM) error {

	if new_range_est_m.Value != nil {
		existing_range_est_m.Value = new_range_est_m.Value
	}
	if new_range_est_m.Sigma != nil {
		existing_range_est_m.Sigma = new_range_est_m.Sigma
	}
	return nil
}

func UpdateMaxRangeM(existing_max_range_m *components.MaxRangeM, new_max_range_m *components.MaxRangeM) error {
	if new_max_range_m.Value != nil {
		existing_max_range_m.Value = new_max_range_m.Value
	}
	if new_max_range_m.Sigma != nil {
		existing_max_range_m.Sigma = new_max_range_m.Sigma
	}
	return nil
}
