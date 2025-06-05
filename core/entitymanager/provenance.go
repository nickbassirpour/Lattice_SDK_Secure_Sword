package entitymanager

import (
	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
)

func UpdateProvenance(entity *components.Entity, new_provenance *components.Provenance) error {
	if new_provenance.IntegrationName != nil {
		entity.Provenance.IntegrationName = new_provenance.IntegrationName
	}
	if new_provenance.DataType != nil {
		entity.Provenance.DataType = new_provenance.DataType
	}
	if new_provenance.SourceId != nil {
		entity.Provenance.SourceId = new_provenance.SourceId
	}
	if new_provenance.SourceUpdateTime != nil {
		entity.Provenance.SourceUpdateTime = new_provenance.SourceUpdateTime
	}
	if new_provenance.SourceDescription != nil {
		entity.Provenance.SourceDescription = new_provenance.SourceDescription
	}
	return nil
}
