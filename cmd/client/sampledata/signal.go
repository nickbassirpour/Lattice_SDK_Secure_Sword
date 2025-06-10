package sampledata

import (
	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
)

func SampleSignal() *components.Signal {
	return &components.Signal{
		FrequencyCenter:         sampleFrequencyCenter(),
		FrequencyRange:          SampleFrequencyRange(),
		BandwidthHz:             Pointer(float64(330209.25)),
		SignalToNoiseRatio:      Pointer(float64(20056.23)),
		LineOfBearing:           SampleLineOfBearing(),
		Fixed:                   nil,
		EmitterNotations:        sampleEmitterNotations(),
		PulseWidthS:             Pointer(float64(199002.324325)),
		PulseRepetitionInterval: samplePulseRepetitionInterval(),
		ScanCharacteristics:     sampleScanCharacteristics(),
	}
}

func sampleFrequencyCenter() *components.FrequencyCenter {
	return &components.FrequencyCenter{
		FrequencyHz: SampleFrequencyHz(),
	}
}

func SampleFrequencyHz() *components.FrequencyHz {
	return &components.FrequencyHz{
		Value: Pointer(float64(2356.346)),
		Sigma: Pointer(float64(1005.36)),
	}
}

func SampleFrequencyRange() *components.FrequencyRange {
	return &components.FrequencyRange{
		MinimumFrequencyRangeHz: &components.MinimumFrequencyRangeHz{
			FrequencyHz: SampleFrequencyHz(),
		},
		MaximumFrequencyRangeHz: &components.MaximumFrequencyRangeHz{
			FrequencyHz: SampleFrequencyHz(),
		},
	}
}

func sampleEmitterNotations() []*components.EmitterNotation {
	emitterNotations := []*components.EmitterNotation{}
	emitterNotation := &components.EmitterNotation{
		EmitterNotation: Pointer("Sample emitter notation"),
		Confidence:      Pointer(float64(2330.2)),
	}
	emitterNotations = append(emitterNotations, emitterNotation)
	return emitterNotations
}

func samplePulseRepetitionInterval() *components.PulseRepetitionInterval {
	return &components.PulseRepetitionInterval{
		PulseRepetitionIntervalS: &components.PulseRepetitionIntervalS{
			Value: Pointer(float64(1239803456.346)),
			Sigma: Pointer(float64(1003885.346)),
		},
	}
}

func sampleScanCharacteristics() *components.ScanCharacteristics {
	return &components.ScanCharacteristics{
		ScanType:    Pointer(components.ScanType_SCAN_TYPE_AGILE_BEAM),
		ScanPeriodS: Pointer(float64(213990.3245)),
	}
}
