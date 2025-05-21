package entitymanager

import (
	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
)

type Server struct {
	components.UnimplementedEntityManagerServer
}
