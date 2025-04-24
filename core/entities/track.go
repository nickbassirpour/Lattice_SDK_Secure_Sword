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
	EntityId string `json:"entityId"`

	// Expiration time that must be greater than the current time and less than 30 days in the future.
	// The Entities API will reject any entity update with an expiry_time in the past.
	// When the expiry_time has passed, the Entities API will delete the entity from the COP and send a DELETE event.
	ExpiryTime time.Time `json:"expiryTime"`

	// Boolean that when true, creates or updates the entity.
	// If false and the entity is still live, triggers a DELETE event.
	IsLive bool `json:"isLive"`

	// The primary data source provenance for this entity.
	Provenance Provenance `json:"provenance"`

	Description string    `json:"description"`
	CreatedTime time.Time `json:"createdTime"`

	Location Location `json:"location"`
}
