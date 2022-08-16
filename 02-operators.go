package main
import "fmt"
func main() {
	//ARITHMETIC OPERATOR
	// + - * / %
	var num1 = 4
	var num2 = 3
	var sum =  num1 + num2
	var difference = num1 - num2
	var quotient = num1 / num2
	var product = num1 * num2
	var remainder = num1 % 2
	fmt.Println(sum)
	fmt.Println(difference)
	fmt.Println(quotient)
	fmt.Println(product)
	fmt.Println(remainder)

	//RELETIONAL OPERATORS >, >=, <, <=, == , !=
	var result = num1 > num2
	fmt.Println(result) //true

	//LOGICAL OPERATORS
	const name = "Dung"
	var age = 22

	var invitedToParty = age > 20 && name == "Dung"
	fmt.Println(invitedToParty) //true

	var invitedToParty1 = !(name == "Dung") && age > 20
	fmt.Println(invitedToParty1) //false

	var invitedToParty2 = age > 100 || name == "Dung"
	fmt.Println(invitedToParty2) //true

	var invitedToParty3 = !(name == "Dung") || age < 100
	fmt.Println(invitedToParty3) //true
}