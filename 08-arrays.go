package main

import "fmt"

func main() {
	purchases := [2]float32{19.9, 20.99}
	purchase2 := purchases[1]
	fmt.Println(purchases)
	fmt.Println(purchase2)

	var sales [3]float32
	sales[0] = 1.05
	fmt.Println(sales)

	numbers := [...]float32{1}
	fmt.Println(numbers)

	for i := 0; i < len(purchases); i++ {
		fmt.Println(purchases[i])
	}

	for i, item := range purchases {
		fmt.Println(i)
		fmt.Println(item)
	}
}