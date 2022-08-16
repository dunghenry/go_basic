package main
import (
	"fmt"
)
func main() {
	animal := "frog"
	switch animal {
	case "cat":
		println("Meow!")
	case "dog":
		println("Woof!")
	case "frog":
		println("Ribbit!")
	case "horse":
		println("Neight!")
	default:
		fmt.Println("invalid animal")
	}

	//fallthrough switch
	var num int = 45
	switch {
	case num < 10:
		fmt.Println("num is less than 10")
		fallthrough
	case num < 50:
		fmt.Println("num is less than 50")
		fallthrough
	case num < 100:
		fmt.Println("num is less than 100")
		fallthrough
	default:
		fmt.Println("No number")
	}
	// ! result
	// num is less than 50
	// num is less than 100
	// No number
}