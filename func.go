package main

import (
	"bufio"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	// Читаем строки до тех пор, пока не будет пустой строкой
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		//fmt.Println("Введено:", text)
	}
	return
}
