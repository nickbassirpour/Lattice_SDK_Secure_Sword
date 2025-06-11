package entitymanager

import (
	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
)

func UpdateRelationships(entity *components.Entity, newRelationships *components.Relationships) error {
	if len(newRelationships.Relationships) > 0 {
		existingRelationships := make(map[string]*components.Relationship)
		for _, relationship := range entity.Relationships.Relationships {
			existingRelationships[*relationship.RelationshipId] = relationship
		}
		for _, newRelationship := range newRelationships.Relationships {
			if exists, found := existingRelationships[*newRelationship.RelationshipId]; found {
				err := UpdateRelationship(exists, newRelationship)
				if err != nil {
					return nil
				}
			}
		}
	}
	return nil
}

func UpdateRelationship(existingRelationship *components.Relationship, newRelationship *components.Relationship) error {
	if newRelationship.RelatedEntityId != nil {
		existingRelationship.RelatedEntityId = newRelationship.RelatedEntityId
	}
	if newRelationship.RelationshipType != nil {
		err := UpdateRelationshipType(existingRelationship.RelationshipType, newRelationship.RelationshipType)
		if err != nil {
			return err
		}
	}
	return nil
}

func UpdateRelationshipType(existingRelationshipType *components.RelationshipType, newRelationshipType *components.RelationshipType) error {
	if newRelationshipType.TrackedBy != nil {
		newTrackedBy := newRelationshipType.TrackedBy
		existingTrackedBy := existingRelationshipType.TrackedBy
		if newTrackedBy.ActivelyTrackingSensors != nil {
			if newTrackedBy.ActivelyTrackingSensors.Sensors != nil && len(newTrackedBy.ActivelyTrackingSensors.Sensors) > 0 {
				newSensors, err := UpdateSensors(existingTrackedBy.ActivelyTrackingSensors.Sensors, newTrackedBy.ActivelyTrackingSensors.Sensors)
				if err != nil {
					return err
				}
				existingTrackedBy.ActivelyTrackingSensors.Sensors = newSensors
			}
		}
	}
	if newRelationshipType.GroupChild != nil {
		existingRelationshipType.GroupChild = newRelationshipType.GroupChild
	}
	if newRelationshipType.GroupParent != nil {
		existingRelationshipType.GroupParent = newRelationshipType.GroupParent
	}
	if newRelationshipType.MergedFrom != nil {
		existingRelationshipType.MergedFrom = newRelationshipType.MergedFrom
	}
	if newRelationshipType.ActiveTarget != nil {
		existingRelationshipType.ActiveTarget = newRelationshipType.ActiveTarget
	}
	return nil
}
