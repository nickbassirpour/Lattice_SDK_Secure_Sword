package entitymanager

import (
	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
)

func UpdateLocationUncertainty(entity *components.Entity, new_location_uncertainty *components.LocationUncertainty) error {
	if new_location_uncertainty.PositionEnuCov != nil {
		if new_location_uncertainty.PositionEnuCov.Mxx != nil {
			entity.LocationUncertainty.PositionEnuCov.Mxx = new_location_uncertainty.PositionEnuCov.Mxx
		}
		if new_location_uncertainty.PositionEnuCov.Mxy != nil {
			entity.LocationUncertainty.PositionEnuCov.Mxy = new_location_uncertainty.PositionEnuCov.Mxy
		}
		if new_location_uncertainty.PositionEnuCov.Mxz != nil {
			entity.LocationUncertainty.PositionEnuCov.Mxz = new_location_uncertainty.PositionEnuCov.Mxz
		}
		if new_location_uncertainty.PositionEnuCov.Myy != nil {
			entity.LocationUncertainty.PositionEnuCov.Myy = new_location_uncertainty.PositionEnuCov.Myy
		}
		if new_location_uncertainty.PositionEnuCov.Myz != nil {
			entity.LocationUncertainty.PositionEnuCov.Myz = new_location_uncertainty.PositionEnuCov.Myz
		}
		if new_location_uncertainty.PositionEnuCov.Mzz != nil {
			entity.LocationUncertainty.PositionEnuCov.Mzz = new_location_uncertainty.PositionEnuCov.Mzz
		}
	}

	if new_location_uncertainty.VelocityEnuCov != nil {
		if new_location_uncertainty.VelocityEnuCov.Mxx != nil {
			entity.LocationUncertainty.VelocityEnuCov.Mxx = new_location_uncertainty.VelocityEnuCov.Mxx
		}
		if new_location_uncertainty.VelocityEnuCov.Mxy != nil {
			entity.LocationUncertainty.VelocityEnuCov.Mxy = new_location_uncertainty.VelocityEnuCov.Mxy
		}
		if new_location_uncertainty.VelocityEnuCov.Mxz != nil {
			entity.LocationUncertainty.VelocityEnuCov.Mxz = new_location_uncertainty.VelocityEnuCov.Mxz
		}
		if new_location_uncertainty.VelocityEnuCov.Myy != nil {
			entity.LocationUncertainty.VelocityEnuCov.Myy = new_location_uncertainty.VelocityEnuCov.Myy
		}
		if new_location_uncertainty.VelocityEnuCov.Myz != nil {
			entity.LocationUncertainty.VelocityEnuCov.Myz = new_location_uncertainty.VelocityEnuCov.Myz
		}
		if new_location_uncertainty.VelocityEnuCov.Mzz != nil {
			entity.LocationUncertainty.VelocityEnuCov.Mzz = new_location_uncertainty.VelocityEnuCov.Mzz
		}
	}

	if new_location_uncertainty.PositionErrorEllipse != nil {
		if new_location_uncertainty.PositionErrorEllipse.Probability != nil {
			entity.LocationUncertainty.PositionErrorEllipse.Probability = new_location_uncertainty.PositionErrorEllipse.Probability
		}
		if new_location_uncertainty.PositionErrorEllipse.SemiMajorAxisM != nil {
			entity.LocationUncertainty.PositionErrorEllipse.SemiMajorAxisM = new_location_uncertainty.PositionErrorEllipse.SemiMajorAxisM
		}
		if new_location_uncertainty.PositionErrorEllipse.SemiMinorAxisM != nil {
			entity.LocationUncertainty.PositionErrorEllipse.SemiMinorAxisM = new_location_uncertainty.PositionErrorEllipse.SemiMinorAxisM
		}
		if new_location_uncertainty.PositionErrorEllipse.OrientationD != nil {
			entity.LocationUncertainty.PositionErrorEllipse.OrientationD = new_location_uncertainty.PositionErrorEllipse.OrientationD
		}
	}
	return nil
}
