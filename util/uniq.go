package main

import (
	"flag"

	"github.com/lecranck/unique"
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

func (opt *Options) ParseOptions(c, d, u bool, f, s int, i bool) {

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
	var inputFile, outputFile string

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
		inputFile = files[0]
	} else {
		if len(files) == 2 {
			inputFile = files[0]
			outputFile = files[1]
		}
	}

	// fill struct with arguments
	options := Options{}
	options.ParseOptions(c, d, u, f, s, i)

	// data input
	inputData := make([]string, 20)
	source := "stdin"

	if inputFile != "" {
		source = inputFile
	}

	inputData = readFrom(source)
	output_data := unique.GetOutput(inputData, options.c, options.d, options.u, options.f, options.s, options.i)

	// data output
	if outputFile != "" {
		writeInto(outputFile, output_data)
	} else {
		writeStdout(output_data)
	}
}
