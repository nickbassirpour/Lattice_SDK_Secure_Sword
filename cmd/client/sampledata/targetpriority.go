package sampledata

import components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"

func SampleTargetPriority() *components.TargetPriority {
	return &components.TargetPriority{
		HighValueTarget: &components.HighValueTarget{
			IsHighValueTarget: Pointer(true),
			TargetPriority:    Pointer(uint32(123540)),
			TargetMatches:     SampleTargetMatches(),
		},
		Threat: &components.Threat{
			IsThreat: Pointer(true),
		},
	}
}

func SampleTargetMatches() []*components.TargetMatch {
	target_match_list := []*components.TargetMatch{}
	target_match := &components.TargetMatch{
		HighValueTargetListId:        Pointer("235-sdg-88567"),
		HighValueTargetDescriptionId: Pointer("12-567-kdg"),
		IsHighPayoffTarget:           Pointer(true),
	}

	target_match_list = append(target_match_list, target_match)

	return target_match_list
}
