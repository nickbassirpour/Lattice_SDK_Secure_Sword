package entitymanager

import (
	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
)

func UpdateLocationUncertainty(original_location_uncertainty *components.LocationUncertainty, new_location_uncertainty *components.LocationUncertainty) error {
	if new_location_uncertainty.PositionEnuCov != nil {
		if new_location_uncertainty.PositionEnuCov.Mxx != nil {
			original_location_uncertainty.PositionEnuCov.Mxx = new_location_uncertainty.PositionEnuCov.Mxx
		}
		if new_location_uncertainty.PositionEnuCov.Mxy != nil {
			original_location_uncertainty.PositionEnuCov.Mxy = new_location_uncertainty.PositionEnuCov.Mxy
		}
		if new_location_uncertainty.PositionEnuCov.Mxz != nil {
			original_location_uncertainty.PositionEnuCov.Mxz = new_location_uncertainty.PositionEnuCov.Mxz
		}
		if new_location_uncertainty.PositionEnuCov.Myy != nil {
			original_location_uncertainty.PositionEnuCov.Myy = new_location_uncertainty.PositionEnuCov.Myy
		}
		if new_location_uncertainty.PositionEnuCov.Myz != nil {
			original_location_uncertainty.PositionEnuCov.Myz = new_location_uncertainty.PositionEnuCov.Myz
		}
		if new_location_uncertainty.PositionEnuCov.Mzz != nil {
			original_location_uncertainty.PositionEnuCov.Mzz = new_location_uncertainty.PositionEnuCov.Mzz
		}
	}

	if new_location_uncertainty.VelocityEnuCov != nil {
		if new_location_uncertainty.VelocityEnuCov.Mxx != nil {
			original_location_uncertainty.VelocityEnuCov.Mxx = new_location_uncertainty.VelocityEnuCov.Mxx
		}
		if new_location_uncertainty.VelocityEnuCov.Mxy != nil {
			original_location_uncertainty.VelocityEnuCov.Mxy = new_location_uncertainty.VelocityEnuCov.Mxy
		}
		if new_location_uncertainty.VelocityEnuCov.Mxz != nil {
			original_location_uncertainty.VelocityEnuCov.Mxz = new_location_uncertainty.VelocityEnuCov.Mxz
		}
		if new_location_uncertainty.VelocityEnuCov.Myy != nil {
			original_location_uncertainty.VelocityEnuCov.Myy = new_location_uncertainty.VelocityEnuCov.Myy
		}
		if new_location_uncertainty.VelocityEnuCov.Myz != nil {
			original_location_uncertainty.VelocityEnuCov.Myz = new_location_uncertainty.VelocityEnuCov.Myz
		}
		if new_location_uncertainty.VelocityEnuCov.Mzz != nil {
			original_location_uncertainty.VelocityEnuCov.Mzz = new_location_uncertainty.VelocityEnuCov.Mzz
		}
	}

	if new_location_uncertainty.PositionErrorEllipse != nil {
		if new_location_uncertainty.PositionErrorEllipse.Probability != nil {
			original_location_uncertainty.PositionErrorEllipse.Probability = new_location_uncertainty.PositionErrorEllipse.Probability
		}
		if new_location_uncertainty.PositionErrorEllipse.SemiMajorAxisM != nil {
			original_location_uncertainty.PositionErrorEllipse.SemiMajorAxisM = new_location_uncertainty.PositionErrorEllipse.SemiMajorAxisM
		}
		if new_location_uncertainty.PositionErrorEllipse.SemiMinorAxisM != nil {
			original_location_uncertainty.PositionErrorEllipse.SemiMinorAxisM = new_location_uncertainty.PositionErrorEllipse.SemiMinorAxisM
		}
		if new_location_uncertainty.PositionErrorEllipse.OrientationD != nil {
			original_location_uncertainty.PositionErrorEllipse.OrientationD = new_location_uncertainty.PositionErrorEllipse.OrientationD
		}
	}
	return nil
}
