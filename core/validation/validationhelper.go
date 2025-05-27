package validation

import (
	"errors"

	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
)

func ValidateEntity(entity *components.Entity) error {
	if entity == nil {
		return errors.New("entity is nil")
	}
	if entity.EntityId == "" {
		return errors.New("missing EntityId")
	}
	if entity.ExpiryTime == nil || entity.ExpiryTime.AsTime().IsZero() {
		return errors.New("missing ExpiryTime")
	}
	if entity.CreatedTime == nil || entity.CreatedTime.AsTime().IsZero() {
		return errors.New("missing CreatedTime")
	}
	if !entity.IsLive {
		return errors.New("missing IsLive")
	}
	if entity.Provenance != nil {
		if entity.Provenance.IntegrationName == "" {
			return errors.New("missing Provenance IntegrationName")
		}
		if entity.Provenance.DataType == "" {
			return errors.New("missing Provenance DataType")
		}
		if entity.Provenance.SourceUpdateTime == nil || entity.Provenance.SourceUpdateTime.AsTime().IsZero() {
			return errors.New("missing SourceUpdateTime")
		}
	} else {
		return errors.New("missing Provenance")
	}
	if entity.Aliases != nil && entity.Aliases.Name == "" {
		return errors.New("missing Alias Name")
	}
	if entity.Ontology != nil && entity.Ontology.Template == 0 {
		return errors.New("missing Ontology Template")
	}
	return nil
}
