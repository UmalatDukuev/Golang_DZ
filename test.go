package main

import (
	"fmt"
)

func main() {
	str := "one two three four"
	n := 2

	if len(str) > n {
		result := str[n:]
		fmt.Println(result)
	} else {
		fmt.Println("Недостаточно символов для вывода")
	}
}
