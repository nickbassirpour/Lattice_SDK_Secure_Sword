package entities

// Details related to grouping for this entity
type GroupDetails struct {
	Team Team `json:"team"`

	Echelon Echelon `json:"echelon"`
}

// Describes a Team group type. Comprised of autonomous
// entities where an entity in a Team can only be a part of
// a single Team at a time.
type Team struct{}

// Describes a Echelon group type. Comprised of entities which are members
// of the same unit or echelon. Ex: A group of tanks within a armored
// company or that same company as a member of a battalion.
type Echelon struct {
	ArmyEchelon ArmyEchelon `json:"armyEchelon"`
}

type ArmyEchelon int

const (
	ARMY_ECHELON_INVALID = iota
	ARMY_ECHELON_FIRE_TEAM
	ARMY_ECHELON_SQUAD
	ARMY_ECHELON_PLATOON
	ARMY_ECHELON_COMPANY
	ARMY_ECHELON_BATTALION
	ARMY_ECHELON_REGIMENT
	ARMY_ECHELON_BRIGADE
	ARMY_ECHELON_DIVISION
	ARMY_ECHELON_CORPSARMY_ECHELON_ARMY
)
