package main

import "fmt"

func main() {

	if add(10, 20) <= 10 {
		fmt.Println("Variable a is smaller than 10")
	} else {
		fmt.Println(add(10, 20), "is bigger than 10")
	}

}
func add(a, b int) (i int) {
	return a + b
}
