package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"uniq"
)

func CheckInput(opts uniq.Options) {

	if len(flag.Args()) > 2 {
		fmt.Println("Maximum number of arguments = 2! ")
		return
	}
	input := os.Stdin
	output := os.Stdout
	var err error
	if flag.Arg(0) != "" {
		input, err = os.Open(flag.Arg(0))
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer input.Close()
	}
	if flag.Arg(1) != "" {
		output, err = os.Create(flag.Arg(1))
		if err != nil {
			fmt.Println("Error finding output file:", err)
			return
		}
		defer output.Close()
	}
	scanner := bufio.NewScanner(input)
	writer := bufio.NewWriter(output)
	uniq.CollapseLines(scanner, writer, opts)
}

func main() {
	var opts uniq.Options
	opts = uniq.ParseFlags(opts)
	CheckInput(opts)
	return
}
