package sampledata

import components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"

func UpdateTracked() *components.Tracked {
	return &components.Tracked{
		TrackQualityWrapper: Pointer(int32(34)),
		SensorHits:          Pointer(int32(67)),
		NumberOfObjects:     SampleNumberOfObjects(),
		RadarCrossSection:   Pointer(float64(356.3467)),
		LastMeasurementTime: SampleLastMeasurementTime(),
		LineOfBearing:       SampleLineOfBearing(),
	}
}

func SampleNumberOfObjects() *components.NumberOfObjects {
	return &components.NumberOfObjects{
		LowerBound: Pointer(uint32(13)),
		UpperBound: Pointer(uint32(16)),
	}
}
