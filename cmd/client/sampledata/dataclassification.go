package sampledata

import (
	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
)

func SampleDataClassification() *components.DataClassification {
	return &components.DataClassification{
		Default: sampleDefault(),
		Fields:  sampleField(),
	}
}

func sampleDefault() *components.Default {
	return &components.Default{
		Level:   Pointer(components.Level_CLASSIFICATION_LEVELS_CONFIDENTIAL),
		Caveats: sampleCaveats(),
	}
}

func sampleCaveats() []string {
	caveats := []string{"May not be confidential"}
	return caveats
}

func sampleField() []*components.Field {
	field := &components.Field{
		FieldPath:                 Pointer("bandwidth_hz"),
		ClassificationInformation: sampleDefault(),
	}
	fields := []*components.Field{field}
	return fields
}
