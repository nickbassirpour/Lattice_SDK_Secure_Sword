package sampledata

import (
	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
)

func SampleLocationUncertainty() *components.LocationUncertainty {
	return &components.LocationUncertainty{
		PositionEnuCov: &components.PositionEnuCov{
			Mxx: Pointer(float32(923.15)),
			Mxy: Pointer(float32(13.22)),
			Mxz: Pointer(float32(523.234)),
			Myy: Pointer(float32(326.27)),
			Myz: Pointer(float32(97.445)),
			Mzz: Pointer(float32(93.2)),
		},
		VelocityEnuCov: &components.VelocityEnuCov{
			Mxx: Pointer(float32(23.25)),
			Mxy: Pointer(float32(723.25)),
			Mxz: Pointer(float32(93.25)),
			Myy: Pointer(float32(923.85)),
			Myz: Pointer(float32(928.75)),
			Mzz: Pointer(float32(25.15)),
		},
	}
}
