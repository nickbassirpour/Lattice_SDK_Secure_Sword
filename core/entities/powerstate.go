package entities

// TYPO IN HTTP DOCS
// Represents the state of power sources connected to this entity.
type PowerState struct {
	SourceIdToState SourceIdToState `json:"sourceIdToState"`
}

// This is a map where the key is a unique id of the power source
// and the value is additional information about the power source.
type SourceIdToState struct {
	Key string `json:"key"`

	Value Value `json:"value"`
}

type Value struct {
	// Status of the power source.
	PowerStatus PowerStatus `json:"powerStatus"`

	// Used to determine the type of power source.
	PowerType PowerType `json:"powerType"`

	// Power level of the system. If absent, the power level is assumed to be unknown.
	PowerLevel PowerLevel `json:"powerLevel"`

	// Set of human-readable messages with status of the power system. Typically this
	// would be used in an error state to provide additional error information. This
	// can also be used for informational messages.
	Messages []string `json:"messages"`

	// Whether the power source is offloadable. If the value is missing (as opposed to false)
	// then the entity does not report whether the power source is offloadable.
	Offloadable bool `json:"offloadable"`
}

type PowerStatus int

const (
	POWER_STATUS_INVALID = iota
	POWER_STATUS_UNKNOWN
	POWER_STATUS_NOT_PRESENT
	POWER_STATUS_OPERATING
	POWER_STATUS_DISABLED
	POWER_STATUS_ERROR
)

type PowerType int

const (
	POWER_TYPE_INVALID = iota
	POWER_TYPE_UNKNOWN
	POWER_TYPE_GAS
	POWER_TYPE_BATTERY
)

type PowerLevel struct {
	//Total power capacity of the system.
	Capacity float32 `json:"capacity"`

	// Remaining power capacity of the system.
	Remaining float32 `json:"remaining"`

	// Percent of power remaining.
	PercentRemaining float32 `json:"percentRemaining"`

	// Voltage of the power source subsystem, as reported by the power
	// source. If the source does not report this value this field will be null.
	Voltage float64 `json:"voltage"`

	// Current in amps of the power source subsystem, as reported by the power
	// source. If the source does not report this value this field will be null.
	CurrentAmps float64 `json:"currentAmps"`

	// Estimated minutes until empty. Calculated with consumption at the moment,
	// as reported by the power source. If the source does not report this value
	// this field will be null.
	RunTimeToEmptyMins float64 `json:"runTimeToEmptyMins"`

	// Fuel consumption rate in liters per second.
	ConsumptionRateLPerS float64 `json:"consumptionRateLPerS"`
}
