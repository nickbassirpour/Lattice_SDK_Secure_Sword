package entities

// Available for Entities that have a single or primary Location.
type Location struct {
	Position types.Position `json:"position"`
}
