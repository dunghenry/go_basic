package main
import "fmt"
func main(){
	for i:= 0; i < 10 ;i++{
		fmt.Println(i)
	}

	j := 1
	for j < 10{
		fmt.Println(j)
		j++
	}
	k := 1
	for {
		if k == 10{
			break
		}
		fmt.Println(k)
		k++
	}

	r := 1
	for r < 10 {
		if (r % 2 == 0) {
			r++
			continue;
		}
		fmt.Println(r)
		r++
	}
}