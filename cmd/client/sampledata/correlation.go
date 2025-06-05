package sampledata

import components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"

func SampleCorrelation() *components.Correlation {
	return &components.Correlation{
		Primary: &components.Primary{
			SecondaryEntityIds: []string{
				"12490245",
			},
		},
		Secondary: &components.Secondary{
			PrimaryEntityId: Pointer("3927590"),
			Metadata:        SampleMetadata(),
		},
		Membership: &components.Membership{
			CorrelationSetId: Pointer("10035"),
			Primary: &components.Primary{
				SecondaryEntityIds: []string{
					"12490245",
				},
			},
			NonPrimary: &components.Secondary{
				PrimaryEntityId: Pointer("3927590"),
				Metadata:        SampleMetadata(),
			},
			Metadata: SampleMetadata(),
		},
		Decorrelation: SampleDecorrelation(),
	}
}

func SampleMetadata() *components.Metadata {
	return &components.Metadata{
		Provenance:      SampleProvenance(),
		ReplicationMode: Pointer(components.ReplicationMode_CORRELATION_REPLICATION_MODE_INVALID),
		Type:            Pointer(components.CorrelationType_CORRELATION_TYPE_AUTOMATED),
	}
}

func SampleDecorrelation() *components.Decorrelation {
	return &components.Decorrelation{
		All: &components.All{
			Metadata: SampleMetadata(),
		},
		DecorrelatedEntities: []*components.DecorrelatedEntity{
			{
				EntityId: Pointer("993250"),
				Metadata: SampleMetadata(),
			},
		},
	}
}
