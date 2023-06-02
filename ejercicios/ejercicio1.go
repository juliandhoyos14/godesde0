package ejercicios

import "strconv"

func ConvertirAEntero(numero string) (int, string) {
	numeroToInt, err := strconv.Atoi(numero)
	if err != nil {
		return -1, "Hubo un error" + err.Error()
	} else if numeroToInt >= 100 {
		return numeroToInt, "Es mayor o igual a 100"
	} else {
		return numeroToInt, "Es menor o igual a 100"
	}
}
