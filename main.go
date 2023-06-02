package main

import (
	"fmt"

	"github.com/juliandhoyos14/godesde0/variables"
)

func main() {
	fmt.Println("Hola mundo")
	variables.PrintIntegers()
	variables.RestoVariables()
	status, texto := variables.ConviertoATexto(3455)
	fmt.Println("status: ", status)
	fmt.Println("texto: ", texto)
}
