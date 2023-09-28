package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type Options struct {
	c bool
	d bool
	u bool
	f int
	s int
	i bool
}

func CollapseLines(scanner *bufio.Scanner, writer *bufio.Writer, opts Options) {
	/*prevLine := ""
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
	writer.Flush()*/
}

func ParseFlags(opts Options) Options {
	flag.BoolVar(&opts.c, "c", false, "flag 1")
	flag.BoolVar(&opts.d, "d", false, "flag 2")
	flag.BoolVar(&opts.u, "u", false, "flag 3")
	flag.IntVar(&opts.f, "f", 5, "flag 4")
	flag.IntVar(&opts.s, "s", 0, "flag 5")
	flag.BoolVar(&opts.i, "i", false, "flag 6")
	//	_, _, _, _, _, _ = c, d, u, f, s, i
	flag.Parse()
	return opts
}

func CheckInput(opts Options) {
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
	CollapseLines(scanner, writer, opts)

}

func main() {
	var opts Options
	opts = ParseFlags(opts)
	CheckInput(opts)

	return
}
