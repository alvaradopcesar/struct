package main

import (
	"fmt"

	a "struct/tda"
)

func main() {
	var lista a.Lista
	lista.Agregar_Primero("l")
	fmt.Println(lista.Primero)
}
