package entitymanager

import (
	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
)

func UpdateCorrelation(entity *components.Entity, new_correlation *components.Correlation) error {
	if new_correlation.Primary != nil && len(new_correlation.Primary.SecondaryEntityIds) > 0 {
		secondary_entity_ids := new_correlation.Primary.SecondaryEntityIds
		entity.Correlation.Primary.SecondaryEntityIds = append(entity.Correlation.Primary.SecondaryEntityIds, secondary_entity_ids...)
	}
	if new_correlation.Secondary != nil {
		if new_correlation.Secondary.PrimaryEntityId != nil {
			entity.Correlation.Secondary.PrimaryEntityId = new_correlation.Secondary.PrimaryEntityId
		}
		if new_correlation.Secondary.Metadata != nil {
			metadata := new_correlation.Secondary.Metadata
			err := UpdateMetadata(entity, metadata)
			if err != nil {
				return err
			}
		}
	}
	if new_correlation.Membership != nil {
		err := UpdateMembership(entity, new_correlation)
		if err != nil {
			return err
		}
	}
	if new_correlation.Decorrelation != nil {
		err := UpdateDecorrelation(entity, new_correlation.Decorrelation)
		if err != nil {
			return err
		}
	}
	return nil
}

func UpdateMetadata(entity *components.Entity, metadata *components.Metadata) error {
	if metadata.Provenance != nil {
		err := UpdateProvenance(entity, metadata.Provenance)
		if err != nil {
			return err
		}
	}
	if metadata.ReplicationMode != nil {
		entity.Correlation.Secondary.Metadata.ReplicationMode = metadata.ReplicationMode
	}
	if metadata.Type != nil {
		entity.Correlation.Secondary.Metadata.Type = metadata.Type
	}
	return nil
}

func UpdateMembership(entity *components.Entity, new_correlation *components.Correlation) error {
	membership := new_correlation.Membership
	if membership.CorrelationSetId != nil {
		entity.Correlation.Membership.CorrelationSetId = membership.CorrelationSetId
	}
	if membership.Primary != nil && len(membership.Primary.SecondaryEntityIds) > 0 {
		secondary_entity_ids := membership.Primary.SecondaryEntityIds
		entity.Correlation.Primary.SecondaryEntityIds = append(entity.Correlation.Primary.SecondaryEntityIds, secondary_entity_ids...)
	}
	if membership.NonPrimary != nil {
		non_primary := membership.NonPrimary
		if non_primary.PrimaryEntityId != nil {
			entity.Correlation.Membership.NonPrimary.PrimaryEntityId = non_primary.PrimaryEntityId
		}
		if non_primary.Metadata != nil {
			err := UpdateMetadata(entity, non_primary.Metadata)
			if err != nil {
				return err
			}
		}
	}
	if membership.Metadata != nil {
		err := UpdateMetadata(entity, membership.Metadata)
		if err != nil {
			return err
		}
	}
	return nil
}

func UpdateDecorrelation(entity *components.Entity, decorrelation *components.Decorrelation) error {
	if decorrelation.All != nil {
		if decorrelation.All.Metadata != nil {
			err := UpdateMetadata(entity, decorrelation.All.Metadata)
			if err != nil {
				return err
			}
		}
		if len(decorrelation.DecorrelatedEntities) > 0 {
			entity.Correlation.Decorrelation.DecorrelatedEntities = append(entity.Correlation.Decorrelation.DecorrelatedEntities, decorrelation.DecorrelatedEntities...)
		}
	}
	return nil
}
