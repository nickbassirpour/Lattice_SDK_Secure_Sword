package sampledata

import (
	"time"

	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func SampleSensors() *components.Sensors {
	return &components.Sensors{
		Sensors: GenerateSensors(3),
	}
}

func GenerateSensors(num_of_sensors int) []*components.Sensor {
	sensors := []*components.Sensor{}
	for i := 0; i < num_of_sensors; i++ {
		sensors = append(sensors, SampleSensor())
	}
	return sensors
}

func SampleSensor() *components.Sensor {
	return &components.Sensor{
		SensorId:               Pointer("325-019-wf912"),
		OperationalState:       Pointer(components.OperationalState_OPERATIONAL_STATE_DEGRADED),
		SensorType:             Pointer(components.SensorType_SENSOR_TYPE_GPS),
		SensorDescription:      Pointer("A gps system"),
		RfConfiguration:        SampleRfConfiguration(),
		LastDetectionTimestamp: timestamppb.New(time.Now()),
		FieldsOfView:           SampleFieldsOfView(),
	}
}

func SampleRfConfiguration() *components.RFConfiguration {
	return &components.RFConfiguration{
		FrequencyRangeHz: SampleFreqRangeHz(),
		BandwidthRangeHz: SampleBandwidthRangeHz(),
	}
}

func SampleFreqRangeHz() []*components.FrequencyRangeHz {
	freq_range_hz_list := []*components.FrequencyRangeHz{}

	freq_range_hz := &components.FrequencyRangeHz{}
	freq_range_hz.MinimumFrequencyRangeHz = SampleMinimumFrequencyRangeHz()
	freq_range_hz.MaximumFrequencyRangeHz = SampleMaximumFrequencyRangeHz()

	freq_range_hz_list = append(freq_range_hz_list, freq_range_hz)

	return freq_range_hz_list
}

func SampleMinimumFrequencyRangeHz() *components.MinimumFrequencyRangeHz {
	return &components.MinimumFrequencyRangeHz{
		FrequencyHz: &components.FrequencyHz{
			Value: Pointer(float64(235.68)),
			Sigma: Pointer(float64(1990.67)),
		},
	}
}

func SampleMaximumFrequencyRangeHz() *components.MaximumFrequencyRangeHz {
	return &components.MaximumFrequencyRangeHz{
		FrequencyHz: &components.FrequencyHz{
			Value: Pointer(float64(305.68)),
			Sigma: Pointer(float64(2003.67)),
		},
	}
}

func SampleBandwidthRangeHz() []*components.BandwidthRangeHz {
	bandwidth_range_hz_list := []*components.BandwidthRangeHz{}

	bandwidth_range_hz := &components.BandwidthRangeHz{}
	bandwidth_range_hz.MinimumBandwidth = SampleMinimumBandwidth()
	bandwidth_range_hz.MaximumBandwidth = SampleMaximumBandwidth()

	bandwidth_range_hz_list = append(bandwidth_range_hz_list, bandwidth_range_hz)

	return bandwidth_range_hz_list
}

func SampleMinimumBandwidth() *components.MinimumBandwidth {
	return &components.MinimumBandwidth{
		BandwidthHz: Pointer(float64(235.68)),
	}
}

func SampleMaximumBandwidth() *components.MaximumBandwidth {
	return &components.MaximumBandwidth{
		BandwidthHz: Pointer(float64(265.68)),
	}
}

func SampleFieldsOfView() []*components.FieldOfView {
	fields_of_view := []*components.FieldOfView{}

	field_of_view := &components.FieldOfView{
		FovId:              Pointer(int32(230987)),
		MountId:            Pointer("89435"),
		ProjectFrustum:     SampleProjectFrustum(),
		ProjectedCenterRay: SamplePosition(),
		CenterRayPose:      SampleRelativePose(),
		HorizontalFov:      Pointer(float32(235.6546)),
		VerticalFov:        Pointer(float32(135.115)),
		Range:              Pointer(float32(12200.90)),
		Mode:               components.Mode_SENSOR_MODE_AUTOSENSOR_MODE_MUTE.Enum(),
	}

	fields_of_view = append(fields_of_view, field_of_view)
	return fields_of_view
}

func SampleProjectFrustum() *components.ProjectFrustum {
	return &components.ProjectFrustum{
		UpperLeft:   SamplePosition(),
		UpperRight:  SamplePosition(),
		BottomRight: SamplePosition(),
		BottomLeft:  SamplePosition(),
	}
}

/*

// Multiple fields of view for a single sensor component
type FieldOfView struct {

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

*/
