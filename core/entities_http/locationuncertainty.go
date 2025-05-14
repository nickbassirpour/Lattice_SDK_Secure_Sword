package entities

// Uncertainty of entity position and velocity, if available.
type LocationUncertainty struct {
	PositionEnuCov PositionEnuCov `json:"positionEnuCov"`
	VelocityEnuCov VelocityEnuCov `json:"velocityEnuCov"`
}

// Positional covariance represented by the upper triangle of the
// covariance matrix. It is valid to populate only the diagonal
// of the matrix if the full covariance matrix is unknown.
type PositionEnuCov struct {
	MXX float32 `json:"mxx"`
	MXY float32 `json:"mxy"`
	MXZ float32 `json:"mxz"`
	MYY float32 `json:"myy"`
	MYZ float32 `json:"myz"`
	MZZ float32 `json:"mzz"`
}

// Velocity covariance represented by the upper triangle of the
// covariance matrix. It is valid to populate only the diagonal
// of the matrix if the full covariance matrix is unknown.
type VelocityEnuCov struct {
	MXX float32 `json:"mxx"`
	MXY float32 `json:"mxy"`
	MXZ float32 `json:"mxz"`
	MYY float32 `json:"myy"`
	MYZ float32 `json:"myz"`
	MZZ float32 `json:"mzz"`
}

// An ellipse that describes the certainty probability and error
// boundary for a given geolocation.
type PositionErrorEllipse struct {
	// Defines the probability in percentage
	// that an entity lies within the given ellipse: 0-1.
	Probability float64 `json:"probability"`

	// Defines the distance from the center point of the
	// ellipse to the furthest distance on the perimeter in meters.
	SemiMajorAxisM float64 `json:"semiMajorAxisM"`

	// Defines the distance from the center point of the ellipse
	// to the shortest distance on the perimeter in meters.
	SemiMinorAxisM float64 `json:"semiMinorAxisM"`

	// The orientation of the semi-major relative to true north
	// in degrees from clockwise: 0-180 due to symmetry across the semi-minor axis.
	OrientationD float64 `json:"orientationD"`
}
