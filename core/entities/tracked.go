package entities

import "time"

// Available for Entities that are tracked.
type Tracked struct {
	// Quality score, 0-15, nil if none
	TrackQualityWrapper int32 `json:"trackQualityWrapper"`

	// Sensor hits aggregation on the tracked entity.
	SensorHits int32 `json:"sensorHits"`

	NumberOfObjects NumberOfObjects `json:"numberOfObjects"`

	RadarCrossSection float64 `json:"radarCrossSection"`

	LastMeasurementTime time.Time `json:"lastMeasurementTime"`

	LineOfBearing LineOfBearing `json:"lineOfBearing"`
}

// Estimated number of objects or units that are represented by this entity.
// Known as Strength in certain contexts (Link16) if UpperBound == LowerBound; (strength = LowerBound)
// If both UpperBound and LowerBound are defined; strength is between LowerBound and UpperBound
// (represented as string "Strength: 4-5") If UpperBound is defined only (LowerBound unset), Strength ≤ UpperBound
// If LowerBound is defined only (UpperBound unset), LowerBound ≤ Strength 0 indicates unset.
type NumberOfObjects struct {
	LowerBound uint32 `json:"lowerBound"`
	UpperBound uint32 `json:"upperBound"`
}

// The relative position of a track with respect to the entity that is tracking it. Used for tracks that do not yet
// have a 3D position. For this entity (A), being tracked by some entity (B), this LineOfBearing would express a
// ray from B to A.
type LineOfBearing struct {
	AngleOfArrival AngleOfArrival `json:"angleOfArrival"`
	RangeEstimateM RangeEstimateM `json:"rangeEstimateM"`
	MaxRangeM      MaxRangeM      `json:"maxRangeM"`
}

// The direction pointing from this entity to the detection
type AngleOfArrival struct {
	RelativePose                   RelativePose                   `json:"relativePose"`
	BearingElevationCovarianceRad2 BearingElevationCovarianceRad2 `json:"bearingElevationCovarianceRad2"`
}

type RelativePose struct {
	Pos    Pos         `json:"pos"`
	AttEnu AttitudeEnu `json:"attEnu"`
}

// Geospatial location defined by this Pose.
type Pos struct {
	Lon               float64           `json:"lon"`
	Lat               float64           `json:"lat"`
	Alt               float64           `json:"alt"`
	Is2D              bool              `json:"is2d"`
	AltitudeReference AltitudeReference `json:"altitudeReference"`
}

// Meaning of alt. altitude in meters above either WGS84 or EGM96, use altitude_reference to determine what zero means.
type AltitudeReference int

const (
	ALTITUDE_REFERENCE_INVALID AltitudeReference = iota
	ALTITUDE_REFERENCE_HEIGHT_ABOVE_WGS84
	ALTITUDE_REFERENCE_HEIGHT_ABOVE_EGM96
	ALTITUDE_REFERENCE_UNKNOWN
	ALTITUDE_REFERENCE_BAROMETRIC
	ALTITUDE_REFERENCE_ABOVE_SEA_FLOOR
	ALTITUDE_REFERENCE_BELOW_SEA_SURFACE
)

// Bearing/elevation covariance matrix where bearing is defined in radians CCW+ about the z-axis from the x-axis
// of FLU frame and elevation is positive down from the FL/XY plane. mxx = bearing variance in rad^2 mxy = bearing/elevation
// covariance in rad^2 myy = elevation variance in rad^2
type BearingElevationCovarianceRad2 struct {
	MXX float64 `json:"mxx"`
	MXY float64 `json:"mxy"`
	MYY float64 `json:"myy"`
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
