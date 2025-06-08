package entitymanager

import (
	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
)

func UpdateIndicators(entity *components.Entity, new_indicators *components.Indicators) error {
	if new_indicators.Simulated != nil {
		entity.Indicators.Simulated = new_indicators.Simulated
	}
	if new_indicators.Exercise != nil {
		entity.Indicators.Exercise = new_indicators.Exercise
	}
	if new_indicators.Emergency != nil {
		entity.Indicators.Emergency = new_indicators.Emergency
	}
	if new_indicators.C2 != nil {
		entity.Indicators.C2 = new_indicators.C2
	}
	if new_indicators.Egressable != nil {
		entity.Indicators.Egressable = new_indicators.Egressable
	}
	if new_indicators.Starred != nil {
		entity.Indicators.Starred = new_indicators.Starred
	}
}
