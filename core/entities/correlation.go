package entities

// Available for Entities that are a correlated (N to 1) set of entities.
// This will be present on each entity in the set.
type Correlation struct {
	Primary Primary `json:"primary"`

	Secondary Secondary `json:"secondary"`

	Membership Membership `json:"membership"`
}

// This entity is the primary of a correlation meaning that it serves as
// the representative entity of the correlation set.
type Primary struct {
	// This entity is the primary of a correlation meaning that it serves
	// as the representative entity of the correlation set.
	SecondaryEntityIds []string `json:"secondaryEntityIds"`
}

// This entity is a secondary of a correlation meaning that it will be
// represented by the primary of the correlation set.
type Secondary struct {
	// The primary of this correlation
	PrimaryEntityId string `json:"primaryEntityId"`

	// Metadata about the correlation
	Metadata Metadata `json:"metadata"`
}

type Metadata struct {
	Provenance Provenance `json:"provenance"`

	ReplicationMode ReplicationMode `json:"replicationMode"`

	Type CorrelationType `json:"type"`
}

// TYPO
// Indicates how the correlation will be distributed. Because a correlation is
// composed of multiple secondaries, each of which may have been correlated with
// different replication modes, the distribution of the correlation is composed
// of distributions of the individual entities within the correlation set. For example,
// if there are two secondary entities A and B correlated against a primary C, with A
// having been correlated globally and B having been correlated locally, then the
// correlation set that is distributed globally than what is known locally in the node.
type ReplicationMode int

const (
	CORRELATION_REPLICATION_MODE_INVALID = iota
	CORRELATION_REPLICATION_MODE_LOCAL
	CORRELATION_REPLICATION_MODE_GLOBAL
)

// What type of (de)correlation was this entity added with.
type CorrelationType int

const (
	CORRELATION_TYPE_INVALID = iota
	CORRELATION_TYPE_MANUAL
	CORRELATION_TYPE_AUTOMATED
)

type Membership struct {
	CorrelationSetId string `json:"correlationSetId"`

	Primary Primary `json:"primary"`

	// No object specified in SDK docs
	NonPrimary Secondary `json:"nonPrimary"`

	Metadata Metadata `json:"metadata"`
}

// If present, this entity was explicitly decorrelated from one or more entities.
// An entity can be both correlated and decorrelated as long as they are disjoint
// sets. An example would be if a user in the UI decides that two tracks are not
// actually the same despite an automatic correlator having correlated them. The
// user would then decorrelate the two tracks and this decorrelation would be preserved
// preventing the correlator from re-correlating them at a later time.
type Decorrelation struct {
	All All `json:"all"`

	DecorrelatedEntities []DecorrelatedEntity `json:"decorrelatedEntities"`
}

// This will be specified if this entity was decorrelated against all other entities.
type All struct {
	Metadata Metadata `json:"metadata"`
}

type DecorrelatedEntity struct {
	EntityId string `json:"entityId"`

	Metadata Metadata `json:"metadata"`
}
