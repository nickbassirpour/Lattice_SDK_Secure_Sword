package entitymanager

import (
	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
)

func UpdateOverrides(entity *components.Entity, new_overrides *components.Overrides) error {
	if len(new_overrides.Overrides) > 0 {
		existing_overrides := make(map[*string]*components.Override)
		for _, override := range entity.Overrides.Overrides {
			existing_overrides[override.RequestId] = override
		}

		for i := 0; i < len(new_overrides.Overrides); i++ {
			if exists, found := existing_overrides[new_overrides.Overrides[i].RequestId]; found {
				err := UpdateOverride(exists, new_overrides.Overrides[i])
				if err != nil {
					return err
				}
			} else {
				entity.Overrides.Overrides = append(entity.Overrides.Overrides, new_overrides.Overrides[i])
			}
		}
	}
	return nil
}

func UpdateOverride(existing_overrides *components.Override, new_override *components.Override) error {
	if new_override.FieldPath != nil {
		existing_overrides.FieldPath = new_override.FieldPath
	}
	if new_override.MaskedFieldValue != nil {
		existing_overrides.MaskedFieldValue = new_override.MaskedFieldValue
	}
	if new_override.Status != nil {
		existing_overrides.Status = new_override.Status
	}
	if new_override.Provenance != nil {
		existing_overrides.Provenance = new_override.Provenance
	}
	if new_override.Type != nil {
		existing_overrides.Type = new_override.Type
	}
	if new_override.RequestTimeStamp != nil {
		existing_overrides.RequestTimeStamp = new_override.RequestTimeStamp
	}
	return nil
}
