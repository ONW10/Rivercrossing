package mainworld

import (
	"fmt"
)

//Visualisering av Verden
//W----[+][+][+][+]--\\___//__________________--[+][+][+][+]----E

var human string = ""
var chick string = ""
var fox string = ""
var corn string = ""
var boat string = ""

var HumanP string = ""
var ChickP string = ""
var FoxP string = ""
var CornP string = ""
var BoatP string = ""

var Slot1w string = ""
var Slot2w string = ""
var Slot3w string = ""
var Slot4w string = ""
var Slot1b string = ""
var Slot2b string = ""
var Slot1e string = ""
var Slot2e string = ""
var Slot3e string = ""
var Slot4e string = ""

var SlotLW string = ""
var SlotLE string = ""

func main() {
	boat = "W----\\_[" + Slot1b + "]" + "__" + "[" + Slot2b + "]_//----E"
	fmt.Print(boat)
}

func ShowWorld() {
	fmt.Print(Slot1w + "-" + Slot2w + "-" + Slot3w + "-" + Slot4w + "-W|" + SlotLW + "\\____________//" + SlotLE + "|E-" +
		Slot1e + "-" + Slot2e + "-" + Slot3e + "-" + Slot4e)
}

func placeBoat() {
	boat = ("\\_" + Slot1b + "_" + Slot2b + "_/")
	if BoatP == "West" {
		SlotLW = boat
		SlotLE = ""
	} else if BoatP == "East" {
		SlotLE = boat
		SlotLW = ""
	}
}

func PlaceThings() {
	position()
	placeBoat()
}

func position() {
	if HumanP == "West" {
		goToWestSlot("Human")
	} else if HumanP == "Boat" {
		goToBoatSlot("Human")
	} else if HumanP == "East" {
		goToEastSlot("Human")
	}

	if ChickP == "West" {
		goToWestSlot("Chick")
	} else if ChickP == "Boat" {
		goToBoatSlot("Chick")
	} else if ChickP == "East" {
		goToEastSlot("Chick")
	}

	if FoxP == "West" {
		goToWestSlot("Fox")
	} else if FoxP == "Boat" {
		goToBoatSlot("Fox")
	} else if FoxP == "East" {
		goToEastSlot("Fox")
	}

	if CornP == "West" {
		goToWestSlot("Corn")
	} else if CornP == "Boat" {
		goToBoatSlot("Corn")
	} else if CornP == "East" {
		goToEastSlot("Corn")
	}
}

func goToWestSlot(item string) {
	if Slot1w == "" {
		Slot1w = item
	} else if Slot2w == "" {
		Slot2w = item
	} else if Slot3w == "" {
		Slot3w = item
	} else if Slot4w == "" {
		Slot4w = item
	}
}

func goToBoatSlot(item string) {
	if Slot1b == "_" {
		Slot1b = item
	} else if Slot2b == "_" {
		Slot2b = item
	}
}

func goToEastSlot(item string) {
	if Slot1e == "" {
		Slot1e = item
	} else if Slot2e == "" {
		Slot2e = item
	} else if Slot3e == "" {
		Slot3e = item
	} else if Slot4e == "" {
		Slot4e = item
	}
}
