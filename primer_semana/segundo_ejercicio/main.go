package main

import (
	"fmt"
)

func main() {
	fmt.Println(SumaYResta(10, 5))
	_, suma := SumaYResta(10, 5)
	fmt.Print(suma)
}

func SumaYResta(x int, y int) (resta int, suma int) {
	resta = x - y
	suma = x + y

	return resta, suma
}
