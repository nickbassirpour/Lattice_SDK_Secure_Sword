package entities

import (
	"time"
)

// A track represents any entity tracked by another asset or integration source.
// Tracks are not directly under the control of friendly forces.
// This includes aircraft tracks from radar or sensor hits, signal detections,
// and vehicles, people, or animals detected through cameras.
type Track struct {

	// Unique string identifier. Can be a Globally Unique Identifier (GUID).
	EntityId string `json:"entityId" validate:"required""`

	// A human-readable entity description that's helpful for debugging purposes and
	// human traceability. If this field is empty, the Entity Manager API generates one for you.
	Description string `json:"description"`

	// Boolean that when true, creates or updates the entity.
	// If false and the entity is still live, triggers a DELETE event.
	IsLive bool `json:"isLive" validate:"required""`

	// The time when the entity was first known to the entity producer. If this field is empty,
	// the Entity Manager API uses the current timestamp of when the entity is first received.
	// For example, when a drone is first powered on, it might report its startup time as the
	// created time. The timestamp doesn't change for the lifetime of an entity.
	CreatedTime time.Time `json:"createdTime"`

	// Expiration time that must be greater than the current time and less than 30 days in the future.
	// The Entities API will reject any entity update with an expiry_time in the past.
	// When the expiry_time has passed, the Entities API will delete the entity from the COP and send a DELETE event.
	ExpiryTime time.Time `json:"expiryTime" validate:"required""`

	// The primary data source provenance for this entity.
	Provenance Provenance `json:"provenance" validate:"required"`

	Location Location `json:"location"`
}
