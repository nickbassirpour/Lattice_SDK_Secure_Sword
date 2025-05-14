package entities

// Catalog of supported tasks.
type TaskCatalog struct {
	TaskDefinitions []TaskDefinition `json:"taskDefinitions"`
}

type TaskDefinition struct {
	// Url path must be prefixed with type.googleapis.com/.
	TaskSpecificationUrl string `json:"taskSpecificationUrl"`
}
