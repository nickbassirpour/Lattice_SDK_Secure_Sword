package entitymanager

import (
	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
)

func UpdatePayloads(entity *components.Entity, new_payloads *components.Payloads) error {
	if len(new_payloads.PayloadConfigurations) > 0 {
		existing_configs := make(map[string]*components.Config)
		for _, config := range entity.Payloads.PayloadConfigurations {
			existing_configs[*config.CapabilityId] = config
		}

		for _, new_payload := range new_payloads.PayloadConfigurations {
			if exists, found := existing_configs[*new_payload.CapabilityId]; found {
				err := UpdatePayloadConfiguration(exists, new_payload)
				if err != nil {
					return err
				}
			} else {
				entity.Payloads.PayloadConfigurations = append(entity.Payloads.PayloadConfigurations, new_payload)
			}
		}
	}
	return nil
}

func UpdatePayloadConfiguration(existing_payload *components.Config, new_payload *components.Config) error {
	if new_payload.Quantity != nil {
		existing_payload.Quantity = new_payload.Quantity
	}
	if new_payload.EffectiveEnvironment != nil {
		existing_payload.EffectiveEnvironment = new_payload.EffectiveEnvironment
	}
	if new_payload.PayloadOperationalState != nil {
		existing_payload.PayloadOperationalState = new_payload.PayloadOperationalState
	}
	if new_payload.PayloadDescription != nil {
		existing_payload.PayloadDescription = new_payload.PayloadDescription
	}
	return nil
}
