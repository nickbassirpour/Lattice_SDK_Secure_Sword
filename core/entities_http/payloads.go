package entities

// List of payloads available for an entity.
type Payloads struct {
	PayloadConfigurations []Config `json:"payloadConfigurations"`
}

type Config struct {
	// Identifying ID for the capability. This ID may be used multiple times to
	// represent payloads that are the same capability but have different operational states
	CapabilityId string `json:"capabilityId"`

	// The number of payloads currently available in the configuration.
	Quantity uint32 `json:"quantity"`

	// The target environments the configuration is effective against.
	EffectiveEnvironment Environment `json:"effectiveEnvironment"`

	// The operational state of this payload.
	PayloadOperationalState PayloadOperationalState `json:"payloadOperationalState"`

	// A human readable description of the payload
	PayloadDescription string `json:"payloadDescription"`
}

type PayloadOperationalState int

const (
	PAYLOAD_OPERATIONAL_STATE_INVALID = iota
	PAYLOAD_OPERATIONAL_STATE_OFF
	PAYLOAD_OPERATIONAL_STATE_NON_OPERATIONAL
	PAYLOAD_OPERATIONAL_STATE_DEGRADED
	PAYLOAD_OPERATIONAL_STATE_OPERATIONAL
	PAYLOAD_OPERATIONAL_STATE_OUT_OF_SERVICE
	PAYLOAD_OPERATIONAL_STATE_UNKNOWN
)
