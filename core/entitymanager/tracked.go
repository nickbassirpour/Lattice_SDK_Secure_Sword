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
	}
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
}

func UpdateAngleOfArrival(entity *components.Entity, new_tracked *components.Tracked) error {
	angle_of_arrival := new_tracked.LineOfBearing.AngleOfArrival
	if angle_of_arrival.RelativePose != nil {
		err := UpdateRelativePose(entity, new_tracked)
		if err != nil {
			return err
		}
	}
	if angle_of_arrival.BearingElevationCovarianceRad2 != nil {
		if angle_of_arrival.BearingElevationCovarianceRad2.Mxx != nil {
			entity.Tracked.LineOfBearing.AngleOfArrival.BearingElevationCovarianceRad2.Mxx = angle_of_arrival.BearingElevationCovarianceRad2.Mxx
		}
		if angle_of_arrival.BearingElevationCovarianceRad2.Mxy != nil {
			entity.Tracked.LineOfBearing.AngleOfArrival.BearingElevationCovarianceRad2.Mxy = angle_of_arrival.BearingElevationCovarianceRad2.Mxy
		}
		if angle_of_arrival.BearingElevationCovarianceRad2.Myy != nil {
			entity.Tracked.LineOfBearing.AngleOfArrival.BearingElevationCovarianceRad2.Myy = angle_of_arrival.BearingElevationCovarianceRad2.Myy
		}
	}
	return nil
}

func UpdateRelativePose(entity *components.Entity, new_tracked *components.Tracked) error {
	relative_pose := new_tracked.LineOfBearing.AngleOfArrival.RelativePose
	if relative_pose.Pos != nil {
		if relative_pose.Pos.Lon != nil {
			entity.Tracked.LineOfBearing.AngleOfArrival.RelativePose.Pos.Lon = relative_pose.Pos.Lon
		}
		if relative_pose.Pos.Lat != nil {
			entity.Tracked.LineOfBearing.AngleOfArrival.RelativePose.Pos.Lat = relative_pose.Pos.Lat
		}
		if relative_pose.Pos.Alt != nil {
			entity.Tracked.LineOfBearing.AngleOfArrival.RelativePose.Pos.Alt = relative_pose.Pos.Alt
		}
		if relative_pose.Pos.Is2D != nil {
			entity.Tracked.LineOfBearing.AngleOfArrival.RelativePose.Pos.Is2D = relative_pose.Pos.Is2D
		}
		if relative_pose.Pos.AltitudeReference != nil {
			entity.Tracked.LineOfBearing.AngleOfArrival.RelativePose.Pos.AltitudeReference = relative_pose.Pos.AltitudeReference
		}
	}
	if relative_pose.AttEnu != nil {
		if relative_pose.AttEnu.X != nil {
			entity.Location.AttitudeEnu.X = relative_pose.AttEnu.X
		}
		if relative_pose.AttEnu.Y != nil {
			entity.Location.AttitudeEnu.Y = relative_pose.AttEnu.Y
		}
		if relative_pose.AttEnu.Z != nil {
			entity.Location.AttitudeEnu.Z = relative_pose.AttEnu.Z
		}
		if relative_pose.AttEnu.W != nil {
			entity.Location.AttitudeEnu.W = relative_pose.AttEnu.W
		}
	}
	return nil
}

/*

// Available for Entities that are tracked.
type Tracked struct {

	LineOfBearing LineOfBearing `json:"lineOfBearing"`
}

// The relative position of a track with respect to the entity that is tracking it. Used for tracks that do not yet
// have a 3D position. For this entity (A), being tracked by some entity (B), this LineOfBearing would express a
// ray from B to A.
type LineOfBearing struct {
	RangeEstimateM RangeEstimateM `json:"rangeEstimateM"`
	MaxRangeM      MaxRangeM      `json:"maxRangeM"`
}


// The estimated distance of the detection
type RangeEstimateM struct {
	// The value of the measurement.
	Value float64 `json:"value"`

	// Estimated one standard deviation in same unit as the value.
	Sigma float64 `json:"sigma"`
}

type MaxRangeM struct {
	// The value of the measurement.
	Value float64 `json:"value"`

	// Estimated one standard deviation in same unit as the value.
	Sigma float64 `json:"sigma"`
}

*/
