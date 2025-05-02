package entities

import "time"

// The relationships between this entity and other entities in the common operational picture.
type Relationships struct {
	Relationships []Relationship `json:"relationships"`
}

type Relationship struct {
	// The entity ID to which this entity is related.
	RelatedEntityId string `json:"relatedEntityId"`

	// A unique identifier for this relationship. Allows removing or updating relationships.
	Relationshipid string `json:"relationshipId"`

	RelationshipType RelationshipType `json:"relationshipType"`
}

// The relationship type
type RelationshipType struct {
	TrackedBy TrackedBy `json:"trackedBy"`

	GroupChild GroupChild `json:"groupChild"`

	GroupParent GroupParent `json:"groupParent"`

	MergedFrom MergedFrom `json:"mergedFrom"`

	ActiveTarget ActiveTarget `json:"activeTarget"`
}

// Describes the relationship between the entity being tracked ("tracked entity")
// and the entity that is performing the tracking ("tracking entity").
type TrackedBy struct {
	// Sensor details of the tracking entity's sensors that were active and tracking
	// the tracked entity. This may be a subset of the total sensors available on the
	// tracking entity.
	ActivelyTrackingSensors Sensors `json:"activelyTrackingSensors"`

	// Latest time that any sensor in actively_tracking_sensors detected the tracked entity.
	LastMeasurementTimeStamp time.Time `json:"lastMeasurementTimeStamp"`
}

// A GroupChild relationship is a uni-directional relationship indicating that (1) this
// entity represents an Entity Group and (2) the related entity is a child member of this
// group. The presence of this relationship alone determines that the type of group is an Entity Group.
type GroupChild struct{}

// A GroupParent relationship is a uni-directional relationship indicating that this entity is
//  a member of the Entity Group represented by the related entity. The presence of this
// relationship alone determines that the type of group that this entity is a member of is an Entity Group.
type GroupParent struct{}

// A MergedFrom relationship is a uni-directional relationship indicating that this entity is a
// merged entity whose data has at least partially been merged from the related entity.
type MergedFrom struct{}

// A target relationship is the inverse of TrackedBy; a one-way relation from sensor to target,
// indicating track(s) currently prioritized by a robot.
type ActiveTarget struct{}
