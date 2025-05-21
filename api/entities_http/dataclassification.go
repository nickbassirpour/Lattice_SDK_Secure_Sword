package entities

// A component that describes an entity's security classification levels.
type DataClassification struct {
	Default Default `json:"default"`

	// TYPO
	// The set of individual field classification information which
	// should always precedence over the default classification information.
	Fields []Field `json:"fields"`
}

// The default classification information which should be assumed to apply
// to everything in the entity unless a specific field level classification is present.
type Default struct {
	// Classification level to be applied to the information in question.
	Level Level `json:"level"`

	// Caveats that may further restrict how the information can be disseminated.
	Caveats []string `json:"caveats"`
}

type Level int

const (
	CLASSIFICATION_LEVELS_INVALID = iota
	CLASSIFICATION_LEVELS_UNCLASSIFIED
	CLASSIFICATION_LEVELS_CONTROLLED_UNCLASSIFIED
	CLASSIFICATION_LEVELS_CONFIDENTIAL
	CLASSIFICATION_LEVELS_SECRET
	CLASSIFICATION_LEVELS_TOP_SECRET
)

type Field struct {
	// Proto field path which is the string representation of a field.
	// example: signal.bandwidth_hz would be bandwidth_hz in the signal component
	FieldPath string `json:"fieldPath"`

	// The information which makes up the field level classification marking.
	ClassificationInformation Default `json:"classificationInformation"`
}
