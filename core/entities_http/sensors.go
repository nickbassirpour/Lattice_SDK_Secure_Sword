package entities

import "time"

// List of sensors available for an entity
type Sensors struct {
	Sensors []Sensor `json:"sensor"`
}

type Sensor struct {
	// This generally is used to indicate a specific type at a more detailed granularity. E.g. COMInt or LWIR
	SensorId string `json:"sensorId"`

	OperationalState OperationalState `json:"operationalState"`

	// The type of sensor
	SensorType SensorType `json:"sensorType"`

	// A human readable description of the sensor
	SensorDescription string `json:"sensorDescription"`

	//RF configuration details of the sensor
	RFConfiguration RFConfiguration `json:"rfConfiguration"`

	LastDetectionTimestamp time.Time `json:"lastDetectionTimestamp"`

	FieldsOfView []FieldOfView `json:"fieldsOfView"`
}

type SensorType int

const (
	SENSOR_TYPE_INVALID = iota
	SENSOR_TYPE_RADAR
	SENSOR_TYPE_CAMERA
	SENSOR_TYPE_TRANSPONDER
	SENSOR_TYPE_RF
	SENSOR_TYPE_GPS
	SENSOR_TYPE_PTU_POS
	SENSOR_TYPE_PERIMETER
	SENSOR_TYPE_SONAR
)

type OperationalState int

const (
	OPERATIONAL_STATE_INVALID = iota
	OPERATIONAL_STATE_OFF
	OPERATIONAL_STATE_NON_OPERATIONAL
	OPERATIONAL_STATE_DEGRADED
	OPERATIONAL_STATE_OPERATIONAL
	OPERATIONAL_STATE_DENIED
)

type RFConfiguration struct {
	FrequencyRangeHz []FrequencyRangeHz `json:"frequencyRangeHz"`
	BandwidthRangeHz []BandwidthRangeHz `json:"bandwidthRangeHz"`
}

// Frequency ranges that are available for this sensor.
type FrequencyRangeHz struct {
	MinimumFrequencyRangeHz MinimumFrequencyRangeHz `json:"minimumFrequencyRangeHz"`
	MaximumFrequencyRangeHz MaximumFrequencyRangeHz `json:"maximumFrequencyRangeHz"`
}

// Indicates the lowest measured frequency of a signal (Hz).
type MinimumFrequencyRangeHz struct {
	FrequencyHz FrequencyHz `json:"frequencyHz"`
}

// Indicates the maximum measured frequency of a signal (Hz).
type MaximumFrequencyRangeHz struct {
	FrequencyHz FrequencyHz `json:"frequencyHz"`
}

// Indicates a frequency of a signal (Hz) with its standard deviation.
type FrequencyHz struct {
	// The value of the measurement.
	Value float64 `json:"value"`

	// Estimated one standard deviation in same unit as the value.
	Sigma float64 `json:"sigma"`
}

// Bandwidth ranges that are available for this sensor.
type BandwidthRangeHz struct {
	MinimumBandwidth MinimumBandwidth `json:"minimumBandwidth"`
	MaximumBandwidth MaximumBandwidth `json:"maximumBandwidth"`
}

type MinimumBandwidth struct {
	BandwidthHz float64 `json:"bandwithHz"`
}

type MaximumBandwidth struct {
	BandwidthHz float64 `json:"bandwithHz"`
}

// Multiple fields of view for a single sensor component
type FieldOfView struct {
	// The Id for one instance of a FieldOfView, persisted across multiple updates
	// to provide continuity during smoothing. This is relevant for sensors where the
	// dwell schedule is on the order of milliseconds, making multiple FOVs a
	// requirement for proper display of search beams.
	FovId int32 `json:"fovId"`

	// The Id of the mount the sensor is on.
	MountId string `json:"mountId"`

	ProjectFrustum ProjectFrustum `json:"projectFrustum"`

	// Center ray of the frustum projected onto the ground.
	ProjectedCenterRay Position `json:"projectedCenterRay"`

	// The origin and direction of the center ray for this sensor relative to the ENU frame.
	// A ray which is aligned with the positive X axis in the sensor frame will be transformed
	// into the ray along the sensor direction in the ENU frame when transformed by the
	// quaternion contained in this pose.
	CenterRayPose RelativePose `json:"centerRayPose"`

	// Horizontal field of view in radians
	HorizontalFov float32 `json:"horizontalFov"`

	// Vertical field of view in radians
	VerticalFov float32 `json:"verticalFov"`

	// Sensor range in meters
	Range float32 `json:"range"`

	// The mode that this sensor is currently in, used to display for context in the UI.
	// Some sensors can emit multiple sensor field of views with different modes, for
	// example a radar can simultaneously search broadly and perform tighter bounded tracking.
	Mode Mode `json:"mode"`
}

// The field of view the sensor projected onto the ground.
type ProjectFrustum struct {
	Upperleft   Position `json:"upperLeft"`
	UpperRight  Position `json:"upperRight"`
	BottomRight Position `json:"bottomRight"`
	BottomLeft  Position `json:"bottomLeft"`
}

type Mode int

const (
	SENSOR_MODE_INVALID = iota
	SENSOR_MODE_SEARCH
	SENSOR_MODE_TRACK
	SENSOR_MODE_WEAPON_SUPPORT
	SENSOR_MODE_AUTOSENSOR_MODE_MUTE
)
