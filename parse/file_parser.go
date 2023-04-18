package parse

import (
	"flag"
)

type Files struct {
	InputFile  string
	OutputFile string
}

func (filestruct *Files) getFiles() {
	// check files
	files := flag.Args()

	if len(files) == 1 {
		filestruct.InputFile = files[0]
	} else if len(files) == 2 {
		filestruct.InputFile = files[0]
		filestruct.OutputFile = files[1]
	}
}

func NewFiles() *Files {
	files := Files{}
	files.getFiles()
	return &files
}
