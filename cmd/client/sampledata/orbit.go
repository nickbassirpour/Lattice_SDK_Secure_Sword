package sampledata

import components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"

func SampleOrbit() *components.Orbit {
	return &components.Orbit{
		OrbitMeanElements: sampleOrbitMeanElements(),
	}
}

func sampleOrbitMeanElements() *components.OrbitMeanElements {
	return &components.OrbitMeanElements{
		Metadata:              sampleOrbitMetadata(),
		MeanKeplarianElements: sampleMeanKeplarianElements(),
		TleParameters:         sampleTleParameters(),
	}
}

func sampleOrbitMetadata() *components.OrbitMetadata {
	return &components.OrbitMetadata{
		CreationDate:      Pointer("March 15, 2025"),
		Originator:        Pointer("USS Boxer"),
		MessageId:         Pointer("199-kwgjk1-938"),
		RefFrame:          Pointer(components.RefFrame_ECI_REFERENCE_FRAME_TEME),
		RefFrameEpoch:     Pointer("Sample RefFrame Epoch"),
		MeanElementTheory: Pointer(components.MeanElementTheory_MEAN_ELEMENT_THEORY_SGP4),
	}
}

func sampleMeanKeplarianElements() *components.MeanKeplarianElements {
	return &components.MeanKeplarianElements{
		Epoch:              Pointer(string("end times")),
		SemiMajorAxisKm:    Pointer(float64(19902.3266)),
		MeanMotion:         Pointer(float64(87.346)),
		Eccentricity:       Pointer(float64(6690.3465)),
		InclinationDeg:     Pointer(float64(991035.2)),
		RaOfAscNodeDeg:     Pointer(float64(99012.236)),
		ArgOfPericenterDeg: Pointer(float64(1257.346)),
		MeanAnomalyDeg:     Pointer(float64(902335.253)),
		Gm:                 Pointer(float64(1240253.235)),
	}
}

func sampleTleParameters() *components.TleParamaters {
	return &components.TleParamaters{
		EphemerisType:      Pointer(uint32(23456)),
		ClassificationType: Pointer(string("Floating")),
		NoradCatId:         Pointer(uint32(77568)),
		ElementSetNo:       Pointer(uint32(12567)),
		RevAtEpoch:         Pointer(uint32(12229)),
		Bstar:              Pointer(float64(92350.45374678)),
		Bterm:              Pointer(float64(99230.235)),
		MeanMotionDot:      Pointer(float64(334246.2356)),
		MeanMotionDdot:     Pointer(float64(99230.235)),
		Agom:               Pointer(float64(344001.235)),
	}
}
