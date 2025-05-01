package entities

// Available for Entities that have a single or primary Location.
type Location struct {
	Position     Position     `json:"position"`
	VelocityEnu  VelocityEnu  `json:"velocityEnu"`
	SpeedMps     float64      `json:"speedMps"`
	Acceleration Acceleration `json:"acceleration"`
	AttitudeEnu  AttitudeEnu  `json:"attitudeEnu"`
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
