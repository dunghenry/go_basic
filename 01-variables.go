package main

import "fmt"

func main() {
	//! DECLARING AND ASSIGNING
	var favBook = "Harry Potter"
	var favBook2 string = "Harry Potter 2"

	fmt.Println(favBook)      // Harry Potter
	fmt.Printf("%T", favBook) //string
	fmt.Println()

	fmt.Println(favBook2)      // Harry Potter2
	fmt.Printf("%T", favBook2) //string
	fmt.Println()

	//? REASSIGNING
	favBook = "Power of Habits"
	fmt.Println(favBook) //Power of Habit

	var otherBook string = "Bad Blood"
	fmt.Println(otherBook) //Bad Blood

	var thirdFavBook string
	thirdFavBook = "Diary of Wimpy Kid"
	fmt.Println(thirdFavBook) //Diary of Wimpy Kid

	var myAge int    //default = 0
	var amICool bool //default = false

	fmt.Println(myAge)   //0
	fmt.Println(amICool) //false

	//var favNumber = 22
	// var favChocolate = "Kitkat"
	// var favTeam = "MU"

	//todo COMPOUND CREATION
	// var favNumber, favChocolate, favTeam = 22, "Kitkat", "MU"

	//BLOCK CREATION
	var (
		favNumber    = 22
		favChocolate = "Kitkat"
		favTeam      = "MU"
	)

	fmt.Println(favNumber, favChocolate, favTeam)

	favAnimal := "Tiger"
	fmt.Println(favAnimal) //"Tiger"

	pet1, pet2, pet3 := "cat", "dog", "rat"
	fmt.Println(pet1, pet2, pet3) // cat dog rat
	pet2 = "parrot"
	//!  pet2 := "cat" //no new variables on left side of :=compiler
	fmt.Println(pet2) //parrot
	pet4, pet3 := "parrot", "whale"
	fmt.Println(pet1, pet2, pet3, pet4)

	//? CONSTANTS
	const myName = "Tran Dung"
	//! myName = "Tran Dung" //cannot assign to myName
}
