package sampledata

import components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"

func SamplePowerState() *components.PowerState {
	return &components.PowerState{
		SourceIdToState: SampleSourceIdToState(),
	}
}

func SampleSourceIdToState() *components.SourceIdToState {
	return &components.SourceIdToState{
		Key:   Pointer("source_key"),
		Value: SampleValue(),
	}
}

func SampleValue() *components.Value {
	return &components.Value{
		PowerStatus: Pointer(components.PowerStatus_POWER_STATUS_DISABLED),
		PowerType:   Pointer(components.PowerType_POWER_TYPE_BATTERY),
		PowerLevel:  SamplePowerLevel(),
		Messages:    []string{"Power Level sample message"},
		Offloadable: Pointer(true),
	}
}

func SamplePowerLevel() *components.PowerLevel {
	return &components.PowerLevel{
		Capacity:             Pointer(float32(356)),
		Remaining:            Pointer(float32(233)),
		PercentRemaining:     Pointer(float32(89)),
		Voltage:              Pointer(float64(346457.457)),
		CurrentAmps:          Pointer(float64(34654.457)),
		RunTimeToEmptyMins:   Pointer(float64(234.24)),
		ConsumptionRateLPerS: Pointer(float64(211.67)),
	}
}
