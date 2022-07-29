package main

import "fmt"

func main() {
	fmt.Println(add(10, 10))
	fmt.Println(swap("hello", "world"))
}

func add(a int, b int) int {

	return a + b

}

func swap(a, b string) (string, string) {
	return b, a

}
