package entitymanager

import (
	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
)

func UpdatePowerState(entity *components.Entity, new_power_state *components.PowerState) error {
	if new_power_state.SourceIdToState != nil {
		if new_power_state.SourceIdToState.Key != nil {
			entity.PowerState.SourceIdToState.Key = new_power_state.SourceIdToState.Key
		}
		if new_power_state.SourceIdToState.Value != nil {
			existing_value := entity.PowerState.SourceIdToState.Value
			new_value := new_power_state.SourceIdToState.Value

			err := UpdateValue(existing_value, new_value)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func UpdateValue(existing_value *components.Value, new_value *components.Value) error {
	if new_value.PowerStatus != nil {
		existing_value.PowerStatus = new_value.PowerStatus
	}
	if new_value.PowerType != nil {
		existing_value.PowerType = new_value.PowerType
	}
	if new_value.PowerLevel != nil {
		UpdatePowerLevel(existing_value.PowerLevel, new_value.PowerLevel)
	}
	if len(new_value.Messages) > 0 {
		existing_value.Messages = append(existing_value.Messages, new_value.Messages...)
	}
	if new_value.Offloadable != nil {
		existing_value.Offloadable = new_value.Offloadable
	}
	return nil
}

func UpdatePowerLevel(existing_power_level *components.PowerLevel, new_power_level *components.PowerLevel) error {
	if new_power_level.Capacity != nil {
		existing_power_level.Capacity = new_power_level.Capacity
	}
	if new_power_level.Remaining != nil {
		existing_power_level.Remaining = new_power_level.Remaining
	}
	if new_power_level.PercentRemaining != nil {
		existing_power_level.PercentRemaining = new_power_level.PercentRemaining
	}
	if new_power_level.Voltage != nil {
		existing_power_level.Voltage = new_power_level.Voltage
	}
	if new_power_level.CurrentAmps != nil {
		existing_power_level.CurrentAmps = new_power_level.CurrentAmps
	}
	if new_power_level.RunTimeToEmptyMins != nil {
		existing_power_level.RunTimeToEmptyMins = new_power_level.RunTimeToEmptyMins
	}
	if new_power_level.ConsumptionRateLPerS != nil {
		existing_power_level.ConsumptionRateLPerS = new_power_level.ConsumptionRateLPerS
	}
	return nil
}
