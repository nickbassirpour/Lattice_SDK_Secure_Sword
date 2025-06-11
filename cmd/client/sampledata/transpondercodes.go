package sampledata

import (
	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
)

func SampleTransponderCodes() *components.TransponderCodes {
	return &components.TransponderCodes{
		Mode1:                      Pointer(uint32(4673)),
		Mode2:                      Pointer(uint32(1211)),
		Mode3:                      Pointer(uint32(580002)),
		Mode4InterrogationResponse: Pointer(components.InterrogationResponse_INTERROGATION_RESPONSE_CORRECT),
		Mode5:                      sampleMode5(),
		ModeS:                      sampleModeS(),
	}
}

func sampleMode5() *components.Mode5 {
	return &components.Mode5{
		Mode5InterrogationResponse: Pointer(components.InterrogationResponse_INTERROGATION_RESPONSE_CORRECT),
		Mode5:                      Pointer(uint32(9903302)),
		Mode5PlatformId:            Pointer(uint32(123356)),
	}
}

func sampleModeS() *components.ModeS {
	return &components.ModeS{
		Id:      Pointer("99-slkj-23990-sf"),
		Address: Pointer(uint32(880333)),
	}
}
