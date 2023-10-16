package uniq

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Options struct {
	c bool
	d bool
	u bool
	f int
	s int
	i bool
}

func isEqual(str1 string, str2 string, opts Options) bool {
	if opts.i == true {
		return strings.EqualFold(str1, str2)
	} else {
		return str1 == str2
	}
}

func ParseFlags(opts Options) Options {
	flag.BoolVar(&opts.c, "c", false, "add number of lines")
	flag.BoolVar(&opts.d, "d", false, "stdout repeating lines")
	flag.BoolVar(&opts.u, "u", false, "stdout uniq lines")
	flag.IntVar(&opts.f, "f", 0, "don't consider first num fields")
	flag.IntVar(&opts.s, "s", 0, "don't consider first num symbols")
	flag.BoolVar(&opts.i, "i", false, "")
	flag.Parse()
	return opts
}

func CheckFlags(opts Options) {

	cnt := 0
	if opts.c == true {
		cnt++
	}
	if opts.d == true {
		cnt++
	}
	if opts.u == true {
		cnt++
	}
	if cnt > 1 {
		fmt.Println("You can use only one of c/d/u arguments")
		os.Exit(0)
	}
}

func CollapseLines(lines []string, opts Options) []string {
	result := make([]string, 0)
	CheckFlags(opts)
	if opts.i == true {
		for i := 0; i < len(lines); i++ {
			lines[i] = strings.ToLower(lines[i])
		}
	}
	for i := 0; i < len(lines); i++ {
		if opts.f != 0 {
			words := strings.Fields(lines[i])
			if len(words) > opts.f {
				result := strings.Join(words[opts.f:], " ")
				lines[i] = result
			} else {
				fmt.Println("Недостаточно слов в строке для вывода")
				os.Exit(0)
			}
		}
	}

	for i := 0; i < len(lines); i++ {
		if opts.s != 0 {
			str := lines[i]
			if len(str) > opts.s {
				result := str[opts.s:]
				lines[i] = result
			} else {
				fmt.Println("Недостаточно символов в строке для вывода")
				os.Exit(0)
			}

		}
	}

	cnt := 1
	prevLine := lines[0]
	for i := 1; i < len(lines); i++ {
		line := lines[i]
		if isEqual(line, prevLine, opts) == true {
			cnt++
		} else {
			if opts.c == true {
				result = append(result, strconv.Itoa(cnt)+" "+prevLine+"\n")
			} else {
				if opts.d == true {
					if cnt > 1 {
						result = append(result, prevLine+"\n")
					}
				} else {
					if opts.u == true {
						if cnt == 1 {
							result = append(result, prevLine+"\n")
						}
					} else {
						result = append(result, prevLine+"\n")
					}
				}

			}
			cnt = 1
		}
		prevLine = line
	}
	if opts.c == true {
		result = append(result, strconv.Itoa(cnt)+" "+prevLine+"\n")
	} else {
		if opts.d == true {
			if cnt > 1 {
				result = append(result, prevLine+"\n")
			}
		} else {
			if opts.u == true {
				if cnt == 1 {
					result = append(result, prevLine+"\n")
				}
			} else {
				result = append(result, prevLine+"\n")
			}
		}
	}
	return result
}
