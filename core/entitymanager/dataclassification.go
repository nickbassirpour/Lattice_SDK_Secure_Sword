package entitymanager

import (
	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
)

func UpdateDataClassification(entity *components.Entity, newDataClassification *components.DataClassification) error {
	if newDataClassification.Default != nil {
		newDefault := newDataClassification.Default
		existingDefault := entity.DataClassification.Default
		if newDefault.Level != nil {
			existingDefault.Level = newDefault.Level
		}
		if newDefault.Caveats != nil && len(newDefault.Caveats) > 0 {
			existingDefault.Caveats = append(existingDefault.Caveats, newDefault.Caveats...)
		}
	}
	if newDataClassification.Fields != nil {
		if len(newDataClassification.Fields) > 0 {
			entity.DataClassification.Fields = append(entity.DataClassification.Fields, newDataClassification.Fields...)
		}
	}
	return nil
}
