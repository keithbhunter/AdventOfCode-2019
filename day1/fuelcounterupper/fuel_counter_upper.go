package fuelcounterupper

func FuelRequired(mass int) int {
	totalFuel := 0
	for {
		fuel := mass/3 - 2
		if fuel < 0 {
			break
		}
		totalFuel += fuel
		mass = fuel
	}
	return totalFuel
}
