package variables

import "fmt"

func PrintIntegers() {
	var commonInt int
	int32 := int32(10)
	int64 := int64(100)

	fmt.Println("Common Int = ", commonInt)
	fmt.Println("Int 32 = ", int32)
	fmt.Println("Int 64 = ", int64)
}
