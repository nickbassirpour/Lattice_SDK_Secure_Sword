package entitymanager

import (
	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
)

func UpdateTaskCatalog(entity *components.Entity, newTaskCatalog *components.TaskCatalog) error {
	if newTaskCatalog.TaskDefinitions != nil && len(newTaskCatalog.TaskDefinitions) > 0 {
		entity.TaskCatalog.TaskDefinitions = append(entity.TaskCatalog.TaskDefinitions, newTaskCatalog.TaskDefinitions...)
	}
	return nil
}
