package entitymanager

import (
	"fmt"

	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
)

func UpdateTransponderCodes(entity *components.Entity, newTransponderCodes *components.TransponderCodes) error {
	if newTransponderCodes.Mode1 != nil {
		entity.TransponderCodes.Mode1 = newTransponderCodes.Mode1
	}
	if newTransponderCodes.Mode2 != nil {
		entity.TransponderCodes.Mode2 = newTransponderCodes.Mode2
	}
	if newTransponderCodes.Mode3 != nil {
		entity.TransponderCodes.Mode3 = newTransponderCodes.Mode3
	}
	if newTransponderCodes.Mode4InterrogationResponse != nil {
		entity.TransponderCodes.Mode4InterrogationResponse = newTransponderCodes.Mode4InterrogationResponse
	}
	if newTransponderCodes.Mode5 != nil {
		err := UpdateMode5(entity.TransponderCodes.Mode5, newTransponderCodes.Mode5)
		if err != nil {
			return err
		}
	}
	if newTransponderCodes.ModeS != nil {
		err := UpdateModeS(entity.TransponderCodes.ModeS, newTransponderCodes.ModeS)
		if err != nil {
			return err
		}
	}
	return nil
}

func UpdateMode5(existingMode5 *components.Mode5, newMode5 *components.Mode5) error {
	if newMode5.Mode5InterrogationResponse != nil {
		existingMode5.Mode5InterrogationResponse = newMode5.Mode5InterrogationResponse
	}
	if newMode5.Mode5 != nil {
		existingMode5.Mode5 = newMode5.Mode5
	}
	if newMode5.Mode5PlatformId != nil {
		existingMode5.Mode5PlatformId = newMode5.Mode5PlatformId
	}
	return nil
}

func UpdateModeS(existingModeS *components.ModeS, newModeS *components.ModeS) error {
	if newModeS.Id != nil {
		existingModeS.Id = newModeS.Id
	}
	if newModeS.Address != nil {
		if *newModeS.Address > uint32(16777214) {
			return fmt.Errorf("Mode S address %d exceeds limit of %d", *newModeS.Address, 16777214)
		} else {
			existingModeS.Address = newModeS.Address
		}
	}
	return nil
}
