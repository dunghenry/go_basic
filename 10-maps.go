package main

import "fmt"

func main() {
	cart := make(map[string]int)
	cart["laptop"] = 1
	cart["lamp"] = 3
	lampQuantity, pound1:= cart["lamp"]
	bookQuantity, pound2:= cart["book"]
	fmt.Println(cart)
	fmt.Println(cart["laptop"])
	fmt.Println(cart["lamp"])
	fmt.Println(lampQuantity, pound1) //3 true
	fmt.Println(bookQuantity, pound2) //0 false

	//delete
	// delete(cart, "book")

	n := map[string]string{"name": "Dung"}
	fmt.Println(n)
}