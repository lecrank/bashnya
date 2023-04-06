package parse

import (
	"flag"
)

type Files struct {
	InputFile  string
	OutputFile string
}

func (filestruct *Files) Fill() {
	var inputFile, outputFile string

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
	filestruct.InputFile = inputFile
	filestruct.OutputFile = outputFile
}
