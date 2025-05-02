package entities

import "time"

type Overrides struct {
	Override []Override `json:"override"`
}

type Override struct {
	// override request id for an override request
	RequestId string `json:"requestId"`

	// proto field path which is the string representation of a field.
	// example: correlated.primary_entity_id would be primary_entity_id in correlated component
	FieldPath string `json:"fieldPath"`

	// new field value corresponding to field path. In the shape of an empty entity with only
	// the changed value. example: entity: { mil_view: { disposition: Disposition_DISPOSITION_HOSTILE } }
	MaskedFieldValue Entity `json:"maskedFieldValue"`

	// status of the override
	Status OverideStatus `json:"status"`

	Provenance Provenance `json:"provenance"`

	Type Type `json:"type"`

	// Timestamp of the override request. The timestamp is generated
	// by the Entity Manager instance that receives the request.
	RequestTimestamp time.Time `json:"requestTimeStamp"`
}

type OverideStatus int

const (
	OVERRIDE_STATUS_INVALID = iota
	OVERRIDE_STATUS_APPLIED
	OVERRIDE_STATUS_PENDING
	OVERRIDE_STATUS_TIMEOUT
	OVERRIDE_STATUS_REJECTED
	OVERRIDE_STATUS_DELETION_PENDING
)

type Type int

const (
	OVERRIDE_TYPE_INVALID = iota
	OVERRIDE_TYPE_LIVE
	OVERRIDE_TYPE_POST_EXPIRY
)
