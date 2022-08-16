package main
import "fmt"
func main() {
	type Animal struct {
		class  string
		age    int
		gender bool
	}
	var teddy = Animal{
		class: "bear",
		age: 2,
		gender: true,
	}
	fmt.Println(teddy)
	fmt.Println(teddy.class)
	teddy.age = teddy.age + 1
	fmt.Println(teddy.age)

	var leo = Animal{"lion", 2, false}
	fmt.Println(leo)

	var lalo = Animal{}
	lalo.class = "lion"
	lalo.age = 4
	fmt.Println(lalo)

	tuco := struct{
		class string
		age int
		gender bool
	}{
		class: "sadas",
		age: 12,
		gender: true,
	}

	fmt.Println(tuco)
}