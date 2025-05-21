package entities

// A component that describes a geo-entity.
type GeoDetails struct {
	Type        GeoType     `json:"type"`
	ControlArea ControlArea `json:"controlArea"`
	Acm         Acm         `json:"acm"`
}

type GeoType int

const (
	GEO_TYPE_INVALID GeoType = iota
	GEO_TYPE_GENERAL
	GEO_TYPE_HAZARD
	GEO_TYPE_EMERGENCY
	GEO_TYPE_ENGAGEMENT_ZONE
	GEO_TYPE_CONTROL_AREA
	GEO_TYPE_BULLSEYE
	GEO_TYPE_ACM
)

// Determines the type of control area being represented by the geo-entity,
// in which an asset can, or cannot, operate.
type ControlArea struct {
	Type ControlAreaType `json:"type"`
}

type ControlAreaType int

const (
	CONTROL_AREA_TYPE_INVALID ControlAreaType = iota
	CONTROL_AREA_TYPE_KEEP_IN_ZONE
	CONTROL_AREA_TYPE_KEEP_OUT_ZONE
	CONTROL_AREA_TYPE_DITCH_ZONE
	CONTROL_AREA_TYPE_LOITER_ZONE
)

type Acm struct {
	AcmType AcmType `json:"acmType"`

	// Used for loosely typed associations, such as assignment to a specific fires unit. Limit to 150 characters.
	AcmDescription string `json:"acmDescription"`
}

type AcmType int

const (
	ACM_DETAIL_TYPE_INVALID AcmType = iota
	ACM_DETAIL_TYPE_LANDING_ZONE
)
