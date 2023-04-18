package main

import (
	"log"

	"github.com/lecrank/bashnya/parse"
	"github.com/lecrank/bashnya/read_data"
	"github.com/lecrank/bashnya/unique"
	"github.com/lecrank/bashnya/write_data"
)

func checkError(err error) {
	if err != nil {
		log.Fatalf("%v", err)
	}
}

func main() {

	// fill the struct with arguments
	options, err := parse.OptionsGiven()
	checkError(err)

	// fill the file struct
	files := *(parse.NewFiles())

	// read lines from input stream
	inputData, err := read_data.ReadFile(files.InputFile)
	checkError(err)

	// get result
	output_data := unique.FindUnique(inputData, options)

	err = write_data.WriteFile(files.OutputFile, output_data)
	checkError(err)
}
