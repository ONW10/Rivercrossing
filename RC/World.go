package main

import "fmt"

//Visualisering av Verden
//W----[+][+][+][+]--\\___//__________________--[+][+][+][+]----E

var human string = ""
var chick string = ""
var fox string = ""
var corn string = ""
var boat string = ""

var humanP string = ""
var chickP string = ""
var foxP string = ""
var cornP string = ""
var boatP string = ""

var slot1w string = ""
var slot2w string = ""
var slot3w string = ""
var slot4w string = ""
var slot1b string = ""
var slot2b string = ""
var slot1e string = ""
var slot2e string = ""
var slot3e string = ""
var slot4e string = ""

func World() {
	boat = "\\_[" + slot1b + "]" + "__" + "[" + slot2b + "]_//"
	fmt.Print(boat)
}
