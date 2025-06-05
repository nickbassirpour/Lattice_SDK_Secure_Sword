package sampledata

import (
	"time"

	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func UpdateTracked() *components.Tracked {
	return &components.Tracked{
		TrackQualityWrapper: Pointer(int32(34)),
		SensorHits:          Pointer(int32(67)),
		NumberOfObjects:     SampleNumberOfObjects(),
		RadarCrossSection:   Pointer(float64(356.3467)),
		LastMeasurementTime: timestamppb.New(time.Now()),
		LineOfBearing:       SampleLineOfBearing(),
	}
}

func SampleNumberOfObjects() *components.NumberOfObjects {
	return &components.NumberOfObjects{
		LowerBound: Pointer(uint32(13)),
		UpperBound: Pointer(uint32(16)),
	}
}

func SampleLineOfBearing() *components.LineOfBearing {
	return &components.LineOfBearing{
		AngleOfArrival: SampleAngelOfArrival(),
		RangeEstimateM: &components.RangeEstimateM{
			Value: Pointer(float64(3488.2)),
			Sigma: Pointer(float64(1991.2)),
		},
		MaxRangeM: &components.MaxRangeM{
			Value: Pointer(float64(3488.2)),
			Sigma: Pointer(float64(1991.2)),
		},
	}
}

func SampleAngelOfArrival() *components.AngleOfArrival {
	return &components.AngleOfArrival{
		RelativePose: &components.RelativePose{
			Pos: SamplePos(),
			AttEnu: &components.AttitudeEnu{
				X: Pointer(float64(902.2)),
				Y: Pointer(float64(102.6)),
				Z: Pointer(float64(52.34)),
				W: Pointer(float64(672.1)),
			},
		},
		BearingElevationCovarianceRad2: &components.BearingElevationCovarianceRad2{
			Mxx: Pointer(float64(745.23)),
			Mxy: Pointer(float64(5467.2)),
			Myy: Pointer(float64(16745.1)),
		},
	}
}

func SamplePos() *components.Pos {
	return &components.Pos{
		Lon:               Pointer(float64(462.2)),
		Lat:               Pointer(float64(122.6)),
		Alt:               Pointer(float64(521.34)),
		Is2D:              Pointer(false),
		AltitudeReference: components.AltitudeReference_ALTITUDE_REFERENCE_ABOVE_SEA_FLOOR.Enum(),
	}
}
