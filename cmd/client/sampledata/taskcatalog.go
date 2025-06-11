package sampledata

import (
	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
)

func SampleTaskCatalog() *components.TaskCatalog {
	taskDefinitions := []*components.TaskDefinition{{TaskSpecificationUrl: Pointer("fakeurl.com")}}
	return &components.TaskCatalog{
		TaskDefinitions: taskDefinitions,
	}
}
