package main

import "fmt"

func zeroptr(iptr *int) int{
	return *iptr
}
func main() {
	i := 5
	fmt.Println(zeroptr(&i))
}