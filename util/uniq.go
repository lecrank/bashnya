package main

import (
	"flag"

	"github.com/lecrank/unique"
)

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
	options := unique.Options{}
	options.ParseOptions(c, d, u, f, s, i)

	// data input
	inputData := make([]string, 20)
	source := "stdin"

	if inputFile != "" {
		source = inputFile
	}

	inputData = readFrom(source)
	output_data := unique.FindUnique(inputData, options)

	// data output
	if outputFile != "" {
		writeInto(outputFile, output_data)
	} else {
		writeStdout(output_data)
	}
}
