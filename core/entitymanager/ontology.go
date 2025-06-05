package entitymanager

import (
	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
)

func UpdateOntology(entity *components.Entity, new_ontology *components.Ontology) error {
	if new_ontology.PlatformType != nil {
		entity.Ontology.PlatformType = new_ontology.PlatformType
	}
	if new_ontology.SpecificType != nil {
		entity.Ontology.SpecificType = new_ontology.SpecificType
	}
	if new_ontology.Template != nil {
		entity.Ontology.Template = new_ontology.Template
	}
	return nil
}
