package entities

// A message describing any transponder codes associated
// with Mode 1, 2, 3, 4, 5, S interrogations.
type TransponderCodes struct {
	// The mode 1 code assigned to military assets.
	Mode1 uint32 `json:"mode1"`

	// The mode 2 code assigned to military assets.
	Mode2 uint32 `json:"mode2"`

	// The mode 3 code assigned to military assets.
	Mode3 uint32 `json:"mode3"`

	// The validity of the response from the Mode 4 interrogation.
	Mode4InterrogationResponse InterrogationResponse `json:"mode4InterrogationResponse"`

	// The Mode 5 transponder codes.
	Mode5 Mode5 `json:"mode5"`

	// The Mode S transponder codes.
	ModeS ModeS `json:"modeS"`
}

type InterrogationResponse int

const (
	INTERROGATION_RESPONSE_INVALID = iota
	INTERROGATION_RESPONSE_CORRECT
	INTERROGATION_RESPONSE_INCORRECT
	INTERROGATION_RESPONSE_NO_RESPONSE
)

type Mode5 struct {
	// The validity of the response from the Mode 5 interrogation.
	Mode5InterrogationResponse InterrogationResponse `json:"mode5InterrogationResponse"`

	// The Mode 5 code assigned to military assets.
	Mode5 uint32 `json:"mode5"`

	// The Mode 5 platform identification code.
	Mode5PlatformId uint32 `json:"mode5PlatformId"`
}

type ModeS struct {
	// Mode S identifier which comprises of 8 alphanumeric characters.
	Id string `json:"id"`

	// The Mode S ICAO aircraft address. Expected values are between 1
	// and 16777214 decimal. The Mode S address is considered unique.
	Address uint32 `json:"address"`
}
