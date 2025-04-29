package entities

// Available for Entities that have a single or primary Location.
type Location struct {
	Position     Position     `json:"position"`
	VelocityEnu  VelocityEnu  `json:"velocityEnu"`
	SpeedMps     float64      `json:"speedMps"`
	Acceleration Acceleration `json:"acceleration"`
	AttitudeEnu  AttitudeEnu  `json:"attitudeEnu"`
}

type Position struct {
	// WGS84 geodetic latitude in decimal degrees.
	LatitudeDegrees float64 `json:"latitudeDegrees"`

	// WGS84 longitude in decimal degrees.
	LongitudeDegrees float64 `json:"longitudeDegrees"`

	// Altitude as height above ellipsoid (WGS84) in meters.
	// float64 Value wrapper is used to distinguish optional from default 0.
	AltitudeHaeMeters float64 `json:"altitudeHaeMeters"`

	// Altitude as AGL (Above Ground Level) if the upstream data source has this value set.
	// This value represents the entity's height above the terrain. This is typically measured
	// with a radar altimeter or by using a terrain tile set lookup. If the value is not set
	// from the upstream, this value is not set.
	AltitudeAglMeters float64 `json:"altitudeAglMeters"`

	// Altitude as ASF (Above Sea Floor) if the upstream data source has this value set.
	// If the value is not set from the upstream, this value is not set.
	AltitudeAsfMeters float64 `json:"altitudeAsfMeters"`

	// The depth of the entity from the surface of the water through sensor measurements
	// based on differential pressure between the interior and exterior of the vessel.
	// If the value is not set from the upstream, this value is not set.
	PressureDepthMeters float64 `json:"pressureDepthMeters"`
}

// Velocity in an ENU reference frame centered on the corresponding position. All units are meters per second.
type VelocityEnu struct {
	E float64 `json:"e"`
	N float64 `json:"n"`
	U float64 `json:"u"`
}

// The entity's acceleration in meters/s^2.
type Acceleration struct {
	E float64 `json:"e"`
	N float64 `json:"n"`
	U float64 `json:"u"`
}

// quaternion to translate from entity body frame to it's ENU frame
type AttitudeEnu struct {
	// x, y, z are vector portion, w is scalar
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
	W float64 `json:"w"`
}
