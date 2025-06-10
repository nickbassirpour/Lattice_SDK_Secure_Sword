package entitymanager

import (
	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
)

func UpdateTransponderCodes(entity *components.Entity, new_transponder_codes *components.TransponderCodes) error {
	if new_transponder_codes.Mode1 != nil {
		entity.TransponderCodes.Mode1 = new_transponder_codes.Mode1
    if new_transponder_codes.Mode2 != nil {
		entity.TransponderCodes.Mode2 = new_transponder_codes.Mode2
    if new_transponder_codes.Mode3 != nil {
		entity.TransponderCodes.Mode3 = new_transponder_codes.Mode3
    if new_transponder_codes.Mode4InterrogationResponse != nil {
		entity.TransponderCodes.Mode4InterrogationResponse = new_transponder_codes.Mode4InterrogationResponse
    if new_transponder_codes.Mode5 != nil {
		err := UpdateMode5(entity.TransponderCodes.Mode5, new_transponder_codes.Mode5)
		if err != nil {
			return err
		}
    if new_transponder_codes.ModeS != nil {

	}
	return nil
}

func UpdateMode5(existing_mode5 *components.Mode5, new_mode5 *components.Mode5) error {
	if new_mode5.Mode5InterrogationResponse != nil {
		existing_mode5.Mode5InterrogationResponse = new_mode5.Mode5InterrogationResponse
	}
	if new_mode5.Mode5 != nil {
		existing_mode5.Mode5 = new_mode5.Mode5
	}
	if new_mode5.Mode5PlatformId != nil {
		existing_mode5.Mode5PlatformId = new_mode5.Mode5PlatformId
	}
	
}