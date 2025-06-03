package sampledata

import components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"

func SampleOntology() *components.Ontology {
	return &components.Ontology{
		Template: Pointer(components.Template(0)),
	}
}
