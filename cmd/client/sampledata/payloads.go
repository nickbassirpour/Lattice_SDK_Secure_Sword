package sampledata

import components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"

func SamplePayloads() *components.Payloads {
	return &components.Payloads{
		PayloadConfigurations: SampleConfig(),
	}
}

func SampleConfig() []*components.Config {
	configs_list := []*components.Config{}

	sample_config := &components.Config{
		CapabilityId:            Pointer("23780975"),
		Quantity:                Pointer(uint32(120)),
		EffectiveEnvironment:    Pointer(components.Environment_ENVIRONMENT_AIR),
		PayloadOperationalState: Pointer(components.PayloadOperationalState_PAYLOAD_OPERATIONAL_STATE_DEGRADED),
		PayloadDescription:      Pointer("Big Guns"),
	}

	configs_list = append(configs_list, sample_config)

	return configs_list
}
