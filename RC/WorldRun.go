package main

func main() {
	initializePositions()
	main.PlaceThings()
	main.ShowWorld()
}

func initializePositions() {
	main.FarmerPosition = "West"
	main.ChickenPosition = "West"
	main.FoxPosition = "West"
	main.CornPosition = "West"
	BoatPosition = "West"
}

func moveBoat() {
	if main.BoatPosition == "West" {
		main.BoatPosition = "East"
	} else if main.BoatPosition == "East" {
		main.BoatPosition = "West"
	}
}
