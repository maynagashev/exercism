package purchase

import "fmt"

var licensedVehicles = []string{"car", "truck"}

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	// check if kind is in licensedVehicles
	for _, vehicle := range licensedVehicles {
		if vehicle == kind {
			return true
		}
	}
	return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	var selectedOption = option1
	if option2 < option1 {
		selectedOption = option2
	}
	return fmt.Sprintf("%s is clearly the better choice.", selectedOption)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age < 3 {
		return originalPrice * 0.8
	}
	if age < 10 {
		return originalPrice * 0.7
	}
	return originalPrice * 0.5
}
