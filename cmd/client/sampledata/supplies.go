package sampledata

import (
	"time"

	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func SampleSupplies() *components.Supplies {
	return &components.Supplies{
		Fuel: sampleFuel(),
	}
}

func sampleFuel() []*components.Fuel {
	fuelList := []*components.Fuel{}
	fuel := &components.Fuel{
		FuelId:                        Pointer("991018"),
		Name:                          Pointer("HP-5"),
		ReportedDate:                  timestamppb.New(time.Now()),
		AmountGallons:                 Pointer(uint32(300)),
		MaxAuthorizedCapacityGallons:  Pointer(uint32(350)),
		OperationalRequirementGallons: Pointer(uint32(200)),
		DataClassification:            SampleDataClassification(),
		DataSource:                    Pointer("USS Bunker Hill"),
	}
	fuelList = append(fuelList, fuel)
	return fuelList
}
