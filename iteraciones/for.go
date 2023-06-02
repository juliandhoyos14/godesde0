package iteraciones

import (
	"fmt"
)

func Iterar() {

	//Iteracion forma 1
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	//Iteracion forma 2
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for i := 0; i < 100; i += 10 {
		fmt.Println(i)
	}
}
