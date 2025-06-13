package entitymanager

import (
	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
)

func UpdateSupplies(entity *components.Entity, newSupplies *components.Supplies) error {
	if len(newSupplies.Fuel) > 0 {
		existingFuel := make(map[*string]*components.Fuel)
		for _, fuel := range entity.Supplies.Fuel {
			existingFuel[fuel.FuelId] = fuel
		}
		for _, newFuel := range newSupplies.Fuel {
			if exists, found := existingFuel[newFuel.FuelId]; found {
				err := UpdateFuel(exists, newFuel)
				if err != nil {
					return nil
				}
			} else {
				entity.Supplies.Fuel = append(entity.Supplies.Fuel, newFuel)
			}
		}
	}
	return nil
}

func UpdateFuel(existingFuel *components.Fuel, newFuel *components.Fuel) error {
	if newFuel.Name != nil {
		existingFuel.Name = newFuel.Name
	}
	if newFuel.ReportedDate != nil {
		existingFuel.ReportedDate = newFuel.ReportedDate
	}
	if newFuel.AmountGallons != nil {
		existingFuel.AmountGallons = newFuel.AmountGallons
	}
	if newFuel.MaxAuthorizedCapacityGallons != nil {
		existingFuel.MaxAuthorizedCapacityGallons = newFuel.MaxAuthorizedCapacityGallons
	}
	if newFuel.OperationalRequirementGallons != nil {
		existingFuel.OperationalRequirementGallons = newFuel.OperationalRequirementGallons
	}
	if newFuel.DataClassification != nil {
		existingFuel.DataClassification = newFuel.DataClassification
	}
	if newFuel.DataSource != nil {
		existingFuel.DataSource = newFuel.DataSource
	}
	return nil
}
