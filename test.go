package main

import (
	"fmt"
	"strings"
)

func isEqual(str1 string, str2 string, opts bool) bool {

	if opts == false {
		fmt.Println("2   ")
		fmt.Println(str1 == str2)
		return str1 == str1
	} else {

		fmt.Println("1   ")
		fmt.Println(strings.EqualFold(str1, str2))
		return strings.EqualFold(str1, str2)
	}
}

func main() {
	str1 := "aaa"
	str2 := "AAA"
	i := false
	// fmt.Println(str1 == str2)
	// fmt.Println(strings.EqualFold(str1, str2))
	fmt.Println(isEqual(str1, str2, i))
	// if isEqual(str1, str2, i) == true {
	// 	fmt.Println(1)
	// } else {
	// 	fmt.Println(2)
	// }
}
