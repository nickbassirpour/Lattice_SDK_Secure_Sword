package entities

// A component that describes an entity's signal characteristics.
type Signal struct {
	FrequencyCenter FrequencyCenter `json:"frequencyCenter"`

	FrequencyRange FrequencyRange `json:"frequencyRange"`

	// Indicates the bandwidth of a signal (Hz).
	BandwidthHz float64 `json:"bandwidthHz"`

	// Indicates the signal to noise (SNR) of this signal.
	SignalToNoiseRatio float64 `json:"signalToNoiseRatio"`

	LineOfBearing LineOfBearing `json:"lineOfBearing"`

	// A fix of a signal. No extra fields but it is expected that location should be populated when using this report.
	Fixed Fixed `json:"fixed"`

	// Emitter notations associated with this entity.
	EmitterNotations []EmitterNotation `json:"emitterNotations"`

	// length in time of a single pulse
	PulseWidthS float64 `json:"pulseWidthS"`

	PulseRepetitionInterval PulseRepetitionInterval `json:"pulseRepetitionInterval"`

	ScanCharacteristics ScanCharacteristics `json:"scanCharacteristics"`
}

type FrequencyCenter struct {
	FrequencyHz FrequencyHz `json:"frequencyHz"`
}

type FrequencyRange struct {
	MinimumFrequencyRangeHz MinimumFrequencyRangeHz `json:"minimumFrequencyRangeHz"`
	MaximumFrequencyRangeHz MaximumFrequencyRangeHz `json:"maximumFrequencyRangeHz"`
}

// A fix of a signal. No extra fields but it is expected that location should be populated when using this report.
type Fixed struct{}

type EmitterNotation struct {
	EmitterNotation string `json:"emitterNotation"`

	// confidence as a percentage that the emitter notation in this component is accurate
	Confidence float64 `json:"confidence"`
}

// length in time between the start of two pulses
type PulseRepetitionInterval struct {
	PulseRepetitionIntervalS PulseRepetitionIntervalS `json:"pulseRepetitionIntervalS"`
}

// A component that describes some measured value with error.
type PulseRepetitionIntervalS struct {
	// The value of the measurement.
	Value float64 `json:"value"`

	// Estimated one standard deviation in same unit as the value.
	Sigma float64 `json:"sigma"`
}

// describes how a signal is observing the environment
type ScanCharacteristics struct {
	ScanType ScanType `json:"scanType"`

	ScanPeriodS float64 `json:"scanPeriodS"`
}

type ScanType int

const (
	SCAN_TYPE_INVALID = iota
	SCAN_TYPE_CIRCULAR
	SCAN_TYPE_BIDIRECTIONAL_HORIZONTAL_SECTOR
	SCAN_TYPE_BIDIRECTIONAL_VERTICAL_SECTOR
	SCAN_TYPE_NON_SCANNING
	SCAN_TYPE_IRREGULAR
	SCAN_TYPE_CONICAL
	SCAN_TYPE_LOBE_SWITCHING
	SCAN_TYPE_RASTER
	SCAN_TYPE_CIRCULAR_VERTICAL_SECTOR
	SCAN_TYPE_CIRCULAR_CONICAL
	SCAN_TYPE_SECTOR_CONICAL
	SCAN_TYPE_AGILE_BEAM
	SCAN_TYPE_UNIDIRECTIONAL_VERTICAL_SECTOR
	SCAN_TYPE_UNIDIRECTIONAL_HORIZONTAL_SECTOR
	SCAN_TYPE_UNIDIRECTIONAL_SECTOR
	SCAN_TYPE_BIDIRECTIONAL_SECTOR
)
