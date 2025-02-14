package main

import "fmt"

func variables() {
	// Declarar variables
	var a string = "inicial"
	fmt.Println(a)

	// Declarar multiples variables
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go infiere el tipo de la variable
	var d = true
	fmt.Println(d)

	// Declarar e inicializar una variable
	f := "short"
	fmt.Println(f)
}
