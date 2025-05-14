package entities

// Indicators to describe entity to consumers.
type Indicators struct {
	Simulated bool `json:"simulated"`

	Exercise bool `json:"exercise"`

	Emergency bool `json:"emergency"`

	C2 bool `json:"c2"`

	// Indicates the Entity should be egressed to external sources.
	// Integrations choose how the egressing happens (e.g. if an Entity needs fuzzing).
	Egressable bool `json:"egressable"`

	// A signal of arbitrary importance such that the entity should be globally marked for all users
	Starred bool `json:"starred"`
}
