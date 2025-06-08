package sampledata

import (
	"time"

	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func SampleOverrides() *components.Overrides {
	return &components.Overrides{
		Overrides: SampleOverrideList(),
	}
}

func SampleOverrideList() []*components.Override {
	overrides := []*components.Override{}
	masked_field_value, err := WrapEntityAsAny(&components.Entity{
		MilView: &components.MilView{
			Disposition: components.Disposition_DISPOSITION_ASSUMED_FRIENDLY.Enum(),
		},
	})
	if err != nil {
		panic(err)
	}

	override := &components.Override{
		RequestId:        Pointer("890723"),
		FieldPath:        Pointer("disposition"),
		MaskedFieldValue: masked_field_value,
		Status:           Pointer(components.OverrideStatus_OVERRIDE_STATUS_APPLIED),
		Provenance:       SampleProvenance(),
		Type:             Pointer(components.Type_OVERRIDE_TYPE_LIVE),
		RequestTimeStamp: timestamppb.New(time.Now()),
	}

	overrides = append(overrides, override)

	return overrides
}
