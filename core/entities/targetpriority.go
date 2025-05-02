package entities

// The target prioritization associated with an entity.
type TargetPriority struct {
	HighValueTarget HighValueTarget `json:"highValueTarget"`

	Threat Threat `json:"threat"`
}

// Describes the target priority in relation to high value target lists.
type HighValueTarget struct {
	// Indicates whether the target matches any description from a high value target list.
	IsHighValueTarget bool `json:"isHighValueTarget"`

	// The priority associated with the target. If the target's description
	//appears on multiple high value target lists, the priority will be a
	// reflection of the highest priority of all of those list's target description.
	// A lower value indicates the target is of a higher priority, with 1 being the
	// highest possible priority. A value of 0 indicates there is no priority
	// associated with this target.
	TargetPriotity uint32 `json:"targetPriority"`

	TargetMatches []TargetMatch `json:"targetMatches"`
}

// All of the high value target descriptions that the target matches against.
type TargetMatch struct {
	// The ID of the high value target list that matches the target description.
	HighValueTargetListId string `json:"highValueTargetListId"`

	// The ID of the specific high value target description within a high value
	// target list that was matched against. The ID is considered to be a globally
	// unique identifier across all high value target IDs.
	HighValueTargetDescriptionId string `json:"highValueTargetDescriptionId"`

	// Indicates whether the target is a 'High Payoff Target'. Targets can be one
	// or both of high value and high payoff.
	IsHighPayoffTarget bool `json:"isHighPayoffTarget"`
}

// Describes whether the entity should be treated as a threat
type Threat struct {
	// Indicates that the entity has been determined to be a threat.
	IsThreat bool `json:"isThreat"`
}
