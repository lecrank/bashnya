package main

import (
	"github.com/lecrank/bashnya/parse"
	"github.com/lecrank/bashnya/read_data"
	"github.com/lecrank/bashnya/unique"
	"github.com/lecrank/bashnya/write_data"
)

func main() {

	// fill the struct with arguments
	options := parse.OptionsGiven()

	// fill the file struct
	files := parse.GivenFiles()

	// read lines from input stream
	inputData, err := read_data.ReadFile(files.InputFile)
	write_data.CheckError(err)

	// get result
	output_data := unique.FindUnique(inputData, options)

	err = write_data.WriteFile(files.OutputFile, output_data)
	write_data.CheckError(err)
}
