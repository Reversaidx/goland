package main

import (
	"fmt"
	"os"
)

func main() {
	str_1 := os.Args[1]

	lenght := len(str_1)
	slice := make([]byte, lenght)
	for i := 0; i < lenght/2+lenght%2; i++ {
		tmp := str_1[i]
		slice[i] = byte(str_1[lenght-i-1])
		slice[lenght-i-1] = tmp
	}
	fmt.Printf(string(slice))
}
