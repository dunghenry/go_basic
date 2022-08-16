package main

import "fmt"

func main() {
	purchases := [5]float32{19.99, 20.99, 5.99, 1.99, 14.55};
	mySlice := purchases[:]
	fmt.Println(mySlice)
	mySlice = append(mySlice, 30.99)
	fmt.Println(mySlice)

	myOtherSlice := purchases[0:3]
	myThirdSlice := purchases[2:]
	fmt.Println(myOtherSlice)
	fmt.Println(cap(myOtherSlice))
	fmt.Println(myThirdSlice)
	fmt.Println(cap(myThirdSlice))

	combind := append(myOtherSlice, myThirdSlice...)
	fmt.Println(combind)
}