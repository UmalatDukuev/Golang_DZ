package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
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
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	cnt := 1
	prevLine := lines[0]
	for i := 1; i < len(lines); i++ {
		line := lines[i]
		if line == prevLine {
			cnt++
		} else {
			if opts.c == true {
				writer.WriteString(strconv.Itoa(cnt) + " " + prevLine + "\n")
			} else {
				writer.WriteString(prevLine + "\n")
			}
			cnt = 1
		}
		prevLine = line
	}
	if opts.c == true {
		writer.WriteString(strconv.Itoa(cnt) + " " + prevLine + "\n")
	} else {
		writer.WriteString(prevLine + "\n")
	}
	writer.Flush()
}

func ParseFlags(opts Options) Options {
	flag.BoolVar(&opts.c, "c", false, "add number of lines")
	flag.BoolVar(&opts.d, "d", false, "stdout repeating lines")
	flag.BoolVar(&opts.u, "u", false, "stdout uniq lines")
	flag.IntVar(&opts.f, "f", 5, "flag 4")
	flag.IntVar(&opts.s, "s", 0, "flag 5")
	flag.BoolVar(&opts.i, "i", false, "flag 6")
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
