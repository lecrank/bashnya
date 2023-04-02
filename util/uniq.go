package main

import (
	"files/unique"
	"flag"
)

type Options struct {
	c           bool
	d           bool
	u           bool
	i           bool
	f           int
	s           int
	input_file  string
	output_file string
}

func (opt *Options) parse_options(c, d, u bool, f, s int, i bool) {

	// check if -c -d -u used at the same time
	count := 0
	for _, flag := range []bool{c, d, u} {
		if flag {
			count++
		}
	}

	if count > 1 {
		panic("Wrong usage!\n\nCorrect: uniq [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]")
	}

	opt.c = c
	opt.d = d
	opt.u = u
	opt.i = i
	opt.f = f
	opt.s = s
}

func main() {
	var c, d, u, i bool
	var f, s int
	var input_file, output_file string

	// set flags
	flag.BoolVar(&c, "c", false, "occurrence count")
	flag.BoolVar(&d, "d", false, "non-unique only")
	flag.BoolVar(&u, "u", false, "unique only")
	flag.IntVar(&f, "f", 0, "amount of ignored fields")
	flag.IntVar(&s, "s", 0, "amount of ignored chars")
	flag.BoolVar(&i, "i", false, "ignore register")
	flag.Parse()

	// check files
	files := flag.Args()

	if len(files) == 1 {
		input_file = files[0]
	} else {
		if len(files) == 2 {
			input_file = files[0]
			output_file = files[1]
		}
	}

	// fill struct with arguments
	options := Options{}
	options.parse_options(c, d, u, f, s, i)

	// data input
	input_data := make([]string, 20)
	source := "stdin"

	if input_file != "" {
		source = input_file
	}

	input_data = read_from(source)
	output_data := unique.GetOutput(input_data, options.c, options.d, options.u, options.f, options.s, options.i)

	// data output
	if output_file != "" {
		write_into(output_file, output_data)
	} else {
		write_stdout(output_data)
	}
}
