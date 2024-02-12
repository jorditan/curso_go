package main

import (
	"fmt"
)

func main() {
	Saludar("Matías")
}

func Saludar(nombre string) string {
	if len(nombre) == 0 {
		return "Hola extraño"
	}

	return fmt.Sprintf("Hola %s", nombre)
}
