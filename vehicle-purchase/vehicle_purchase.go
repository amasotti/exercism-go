package purchase

import "fmt"

type VehicleType string

const (
	Car   VehicleType = "car"
	Truck VehicleType = "truck"
)

const (
	RelativelyNewVehicleAgeThreshold = 3
	OldVehicleAgeThreshold           = 10
	NewMultiplier                    = 0.8
	RelativelyNewMultiplier          = 0.7
	OldMultiplier                    = 0.5
)

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	vehicleKind := VehicleType(kind)
	return vehicleKind == Car || vehicleKind == Truck
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	chosen := option1
	if option2 < option1 {
		chosen = option2
	}

	return fmt.Sprintf("%s is clearly the better choice.", chosen)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	switch {
	case age < RelativelyNewVehicleAgeThreshold:
		return originalPrice * NewMultiplier
	case age < OldVehicleAgeThreshold:
		return originalPrice * RelativelyNewMultiplier
	default:
		return originalPrice * OldMultiplier
	}
}
