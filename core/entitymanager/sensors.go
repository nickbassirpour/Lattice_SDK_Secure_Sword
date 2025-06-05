package entitymanager

import (
	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
)

func UpdateSensors(entity *components.Entity, new_sensors []*components.Sensor) error {
	if len(new_sensors) > 0 {
		entity.Sensors
	}
}
