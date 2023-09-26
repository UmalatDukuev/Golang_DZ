package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func CollapseLines(scanner *bufio.Scanner, writer *bufio.Writer) {
	//writer.WriteString("\n")
	prevLine := ""
	if scanner.Scan() {
		prevLine = scanner.Text()
	}
	cnt := 1
	for scanner.Scan() {
		line := scanner.Text()
		if line == prevLine {
			cnt++
		} else {
			writer.WriteString(prevLine + "\n")
			cnt = 1
		}
		prevLine = line
	}
	writer.WriteString(prevLine)
	writer.Flush()
}

func ParseFlags() {
	c := flag.Bool("c", false, "flag 1")
	d := flag.Bool("d", false, "flag 2")
	u := flag.Bool("u", false, "flag 3")
	f := flag.Bool("f", false, "flag 4")
	s := flag.Bool("s", false, "flag 5")
	i := flag.Bool("i", false, "flag 6")
	_, _, _, _, _, _ = c, d, u, f, s, i

	flag.Parse()
}

func CheckInput() {
	if len(flag.Args()) > 2 {
		fmt.Println("Maximum number of arguments = 2! ")
		return
	}
	input := os.Stdin
	output := os.Stdout

	if flag.Arg(0) != "" {
		input, _ = os.Open(flag.Arg(0))
		// if err != nil {
		// 	fmt.Println("Error opening file:", err)
		// 	return
		// }
		defer input.Close()
	}

	if flag.Arg(1) != "" {
		output, _ = os.Create(flag.Arg(1))
		// if err != nil {
		// 	fmt.Println("Error finding output file:", err)
		// 	return
		// }
		defer output.Close()
	}

	scanner := bufio.NewScanner(input)
	writer := bufio.NewWriter(output)
	CollapseLines(scanner, writer)

}

func main() {
	ParseFlags()
	CheckInput()
	return
}
