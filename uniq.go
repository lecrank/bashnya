package main

import (
	"github.com/lecrank/bashnya/parse"
	"github.com/lecrank/bashnya/read_data"
	"github.com/lecrank/bashnya/unique"
	"github.com/lecrank/bashnya/write_data"
)

func main() {

	// fill the struct with arguments
	options := parse.Options{}
	flags := parse.GetFlags(options)

	// fill the file struct
	files := parse.Files{}
	files.Fill()

	// read lines from input stream
	inputData := read_data.ReadFile(files.InputFile)

	// get result
	output_data := unique.FindUnique(inputData, flags)

	write_data.WriteFile(files.OutputFile, output_data)
}
