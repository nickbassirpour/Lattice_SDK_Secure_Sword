package entities

import "time"

// Orbit information for space objects.
type Orbit struct {
	OrbitMeanElements OrbitMeanElements `json:"orbitMeanElements"`
}

// Orbit Mean Elements data, analogous to the Orbit Mean Elements Message in CCSDS 502.0-B-3
type OrbitMeanElements struct {
	Metadata OrbitMetadata `json:"metadata"`

	MeanKeplarianElements MeanKeplarianElements `json:"meanKeplarianElements"`

	TleParamaters TleParamaters `json:"tleParameters"`
}

type OrbitMetadata struct {
	// Creation date/time in UTC
	CreationDate time.Time `json:"creationDate"`

	// Creating agency or operator
	Originator string `json:"originator"`

	// ID that uniquely identifies a message from a given originator.
	MessageId string `json:"messageId"`

	// Reference frame, assumed to be Earth-centered
	RefFrame RefFrame `json:"refFrame"`

	// Reference frame epoch in UTC - mandatory only if not intrinsic to frame definition
	RefFrameEpoch time.Time `json:"refFrameEpoch"`

	MeanElementTheory MeanElementTheory `json:"meanElementTheory"`
}

type RefFrame int

const (
	ECI_REFERENCE_FRAME_INVALID = iota
	ECI_REFERENCE_FRAME_TEME
)

type MeanElementTheory int

const (
	MEAN_ELEMENT_THEORY_INVALID = iota
	MEAN_ELEMENT_THEORY_SGP4
)

type MeanKeplarianElements struct {
	// UTC time of validity
	Epoch time.Time `json:"epoch"`

	// Preferred: semi major axis in kilometers
	SemiMajorAxisKm float64 `json:"semiMajorAxisKm"`

	// If using SGP/SGP4, provide the Keplerian Mean Motion in revolutions per day
	MeanMotion float64 `json:"meanMotion"`

	Eccentricity float64 `json:"eccentricity"`

	// Angle of inclination in deg
	InclinationDeg float64 `json:"inclinationDeg"`

	// Right ascension of the ascending node in deg
	RaOfAscNodeDeg float64 `json:"raOfAscNodeDeg"`

	// Argument of pericenter in deg
	ArgOfPericenterDeg float64 `json:"argOfPericenterDeg"`

	// Mean anomaly in deg
	MeanAnomalyDeg float64 `json:"meanAnomalyDeg"`

	// Optional: gravitational coefficient (Gravitational Constant x central mass) in kg^3 / s^2
	Gm float64 `json:"gm"`
}

type TleParamaters struct {
	// Integer specifying TLE ephemeris type
	EphemerisType uint32 `json:"ephemerisType"`

	// User-defined free-text message classification/caveats of this TLE
	ClassificationType string `json:"classificationType"`

	// Norad catalog number: integer up to nine digits.
	NoradCatId uint32 `json:"noradCatId"`

	ElementSetNo uint32 `json:"elementSetNo"`

	// Optional: revolution number
	RevAtEpoch uint32 `json:"revAtEpoch"`

	// Drag parameter for SGP-4 in units 1 / Earth radii
	Bstar float64 `json:"bstar"`

	// Drag parameter for SGP4-XP in units m^2 / kg
	Bterm float64 `json:"bterm"`

	// First time derivative of mean motion in rev / day^2
	MeanMotionDot float64 `json:"meanMotionDot"`

	// Second time derivative of mean motion in rev / day^3. For use with SGP or PPT3.
	MeanMotionDdot float64 `json:"meanMotionDdot"`

	// Solar radiation pressure coefficient A_gamma / m in m^2 / kg. For use with SGP4-XP.
	Agom float64 `json:"agom"`
}
