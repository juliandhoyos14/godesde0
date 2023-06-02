package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TablaMultiplicar() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrese numero para calcular tabla de multiplicar: ")
	if scanner.Scan() {
		numero1, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El dato ingresado es incorrecto " + err.Error())
		}
		for i := 1; i <= 10; i++ {
			fmt.Println(numero1, "x", i, "=", numero1*i)
		}
	}
}
