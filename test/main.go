package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	// b := append(a[:0], a[1:]...)
	fmt.Println(append(a[:2], a[3:]...))
}
