package sampledata

import components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"

func SampleMilView() *components.MilView {
	return &components.MilView{
		Disposition: components.Disposition_DISPOSITION_ASSUMED_FRIENDLY.Enum(),
		Environment: components.Environment_ENVIRONMENT_AIR.Enum(),
		Nationality: components.Nationality_NATIONALITY_ARGENTINA.Enum(),
	}
}
