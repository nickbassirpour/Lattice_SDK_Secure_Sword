package entitymanager

import (
	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
)

func UpdateOrbit(entity *components.Entity, newOrbit *components.Orbit) error {
	if newOrbit.OrbitMeanElements != nil {
		newOrbitMeanEle := newOrbit.OrbitMeanElements
		existingOrbitMeanEle := entity.Orbit.OrbitMeanElements
		if newOrbitMeanEle.Metadata != nil {
			err := updateOrbitMetadata(existingOrbitMeanEle.Metadata, newOrbitMeanEle.Metadata)
			if err != nil {
				return nil
			}
		}
		if newOrbitMeanEle.MeanKeplarianElements != nil {
			newMeanKeplEle := newOrbit.OrbitMeanElements.MeanKeplarianElements
			existingMeanKeplEle := entity.Orbit.OrbitMeanElements.MeanKeplarianElements
			err := updateOrbitMeanKeplarianElements(existingMeanKeplEle, newMeanKeplEle)
			if err != nil {
				return nil
			}
		}
		if newOrbitMeanEle.TleParameters != nil {
			newTleParams := newOrbit.OrbitMeanElements.TleParameters
			existingTleParams := entity.Orbit.OrbitMeanElements.TleParameters
			err := updateTleParameters(existingTleParams, newTleParams)
			if err != nil {
				return nil
			}
		}
		return nil
	}
}

func updateOrbitMetadata(existingOrbitMetadata *components.OrbitMetadata, newOrbitMetadata *components.OrbitMetadata) error {
	if newOrbitMetadata.CreationDate != nil {
		existingOrbitMetadata.CreationDate = newOrbitMetadata.CreationDate
	}
	if newOrbitMetadata.Originator != nil {
		existingOrbitMetadata.Originator = newOrbitMetadata.Originator
	}
	if newOrbitMetadata.MessageId != nil {
		existingOrbitMetadata.MessageId = newOrbitMetadata.MessageId
	}
	if newOrbitMetadata.RefFrame != nil {
		existingOrbitMetadata.RefFrame = newOrbitMetadata.RefFrame
	}
	if newOrbitMetadata.RefFrameEpoch != nil {
		existingOrbitMetadata.RefFrameEpoch = newOrbitMetadata.RefFrameEpoch
	}
	if newOrbitMetadata.MeanElementTheory != nil {
		existingOrbitMetadata.MeanElementTheory = newOrbitMetadata.MeanElementTheory
	}
	return nil
}

func updateOrbitMeanKeplarianElements(existingMeanKeplEle *components.MeanKeplarianElements, newMeanKeplEle *components.MeanKeplarianElements) error {
	if newMeanKeplEle.Epoch != nil {
		existingMeanKeplEle.Epoch = newMeanKeplEle.Epoch
	}
	if newMeanKeplEle.SemiMajorAxisKm != nil {
		existingMeanKeplEle.SemiMajorAxisKm = newMeanKeplEle.SemiMajorAxisKm
	}
	if newMeanKeplEle.MeanMotion != nil {
		existingMeanKeplEle.MeanMotion = newMeanKeplEle.MeanMotion
	}
	if newMeanKeplEle.Eccentricity != nil {
		existingMeanKeplEle.Eccentricity = newMeanKeplEle.Eccentricity
	}
	if newMeanKeplEle.InclinationDeg != nil {
		existingMeanKeplEle.InclinationDeg = newMeanKeplEle.InclinationDeg
	}
	if newMeanKeplEle.RaOfAscNodeDeg != nil {
		existingMeanKeplEle.RaOfAscNodeDeg = newMeanKeplEle.RaOfAscNodeDeg
	}
	if newMeanKeplEle.ArgOfPericenterDeg != nil {
		existingMeanKeplEle.ArgOfPericenterDeg = newMeanKeplEle.ArgOfPericenterDeg
	}
	if newMeanKeplEle.MeanAnomalyDeg != nil {
		existingMeanKeplEle.MeanAnomalyDeg = newMeanKeplEle.MeanAnomalyDeg
	}
	if newMeanKeplEle.Gm != nil {
		existingMeanKeplEle.Gm = newMeanKeplEle.Gm
	}
	return nil
}

func updateTleParameters(existingTleParams *components.TleParamaters, newTleParams *components.TleParamaters) error {
	if newTleParams.EphemerisType != nil {
		existingTleParams.EphemerisType = newTleParams.EphemerisType
	}
	if newTleParams.ClassificationType != nil {
		existingTleParams.ClassificationType = newTleParams.ClassificationType
	}
	if newTleParams.NoradCatId != nil {
		existingTleParams.NoradCatId = newTleParams.NoradCatId
	}
	if newTleParams.ElementSetNo != nil {
		existingTleParams.ElementSetNo = newTleParams.ElementSetNo
	}
	if newTleParams.RevAtEpoch != nil {
		existingTleParams.RevAtEpoch = newTleParams.RevAtEpoch
	}
	if newTleParams.Bstar != nil {
		existingTleParams.Bstar = newTleParams.Bstar
	}
	if newTleParams.Bterm != nil {
		existingTleParams.Bterm = newTleParams.Bterm
	}
	if newTleParams.MeanMotionDot != nil {
		existingTleParams.MeanMotionDot = newTleParams.MeanMotionDot
	}
	if newTleParams.MeanMotionDdot != nil {
		existingTleParams.MeanMotionDdot = newTleParams.MeanMotionDdot
	}
	if newTleParams.Agom != nil {
		existingTleParams.Agom = newTleParams.Agom
	}
	return nil
}
