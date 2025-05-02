package entities

import "time"

// Represents the state of supplies associated with an entity
// (available but not in condition to use immediately)
type Supplies struct {
	Fuel []Fuel `json:"fuel"`
}

type Fuel struct {
	// unique fuel identifier
	FieldId string `json:"fueldId"`

	// long form name of the fuel source.
	Name string `json:"name"`

	// timestamp the information was reported
	ReportedDate time.Time `json:"reportedDate"`

	// amount of gallons on hand
	AmountGallons uint32 `json:"amountGallons"`

	// how much the asset is allowed to have available (in gallons)
	MaxAuthorizedCapacityGallons uint32 `json:"maxAuthorizedCapacityGallons"`

	// minimum required for operations (in gallons)
	OperationalRequirementGallons uint32 `json:"operationalRequirementGallons"`

	// fuel in a single asset may have different levels of classification use case:
	// fuel for a SECRET asset while diesel fuel may be UNCLASSIFIED
	DataClassification DataClassifiction `json:"dataClassification"`

	// source of information
	DataSource string `json:"dataSource"`
}
