package sampledata

import components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"

func SampleVisualDetails() *components.VisualDetails {
	return &components.VisualDetails{
		RangeRings: sampleRangeRings(),
	}
}

func sampleRangeRings() *components.RangeRings {
	return &components.RangeRings{
		MinDistanceM:  Pointer(float64(100235.13)),
		MaxDistanceM:  Pointer(float64(900235.13)),
		RingCount:     Pointer(uint32(6)),
		RingLineColor: sampleRingLineColor(),
	}
}

func sampleRingLineColor() *components.RingLineColor {
	return &components.RingLineColor{
		Red:   Pointer(float32(600034.2354)),
		Blue:  Pointer(float32(742300.234)),
		Green: Pointer(float32(990345.346)),
		Alpha: Pointer(float32(32095.234235)),
	}
}
