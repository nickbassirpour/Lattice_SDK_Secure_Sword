package entities

import "time"

// The primary data source provenance for this entity.
type Provenance struct {

	// Name of the integration that produced this entity
	IntegrationName string `json:"integrationName"`

	// Source data type of this entity. Examples: ADSB, Link16, etc.
	DataType string `json:"dataType"`

	// An ID that allows an element from a source to be uniquely identified
	SourceId string `json:"sourceId"`

	// The time, according to the source system, that the data in the entity
	// was last modified. Generally, this should be the time that the source-reported
	// time of validity of the data in the entity. This field must be updated with
	// every change to the entity or else Entity Manager will discard the update.
	SourceUpdateTime time.Time `json:"sourceUpdateTime"`

	// Description of the modification source. In the case of a user this is the email address.
	SourceDescription string `json:"sourceDescription"`
}
