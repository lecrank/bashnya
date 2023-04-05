package main

import (
	"github.com/lecrank/bashnya/io_stream"
	"github.com/lecrank/bashnya/parse"
	"github.com/lecrank/bashnya/unique"
)

func main() {

	// init the struct of files
	files := parse.Files{}
	files.Fill()

	reader := io_stream.SetReader(files.InputFile)
	writer := io_stream.SetWriter(files.OutputFile)

	// fill the struct with arguments
	options := parse.Options{}
	flags := parse.GetFlags(options)

	// read lines from input stream
	inputData := readdata.read(reader)

	// get result
	output_data := unique.FindUnique(inputData, flags)

	writedata.write(writer, output_data)
}
