package entities

// Ontology of the entity
type Ontology struct {
	// A string that describes the entity's high-level type with natural language.
	PlatformType string `json:"platformType"`

	// A string that describes the entity's exact model or type.
	SpecificType string `json:"specificType"`

	// The template used when creating this entity. Specifies minimum required components.
	Template Template `json:"template"`
}

type Template int

const (
	TEMPLATE_INVALID = iota
	TEMPLATE_TRACK
	TEMPLATE_SENSOR_POINT_OF_INTEREST
	TEMPLATE_ASSET
	TEMPLATE_GEO
	TEMPLATE_SIGNAL_OF_INTEREST
)
