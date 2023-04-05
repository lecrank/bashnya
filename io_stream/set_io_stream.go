package io_stream

import (
	"bufio"
	"log"
	"os"
)

func checkFileError(err error) {
	if err != nil {
		log.Fatalf("Problem with file: \n%v", err)
	}
}

func SetReader(file string) *bufio.Reader {

	// set default reader
	reader := bufio.NewReader(os.Stdin)

	// set file reader
	if file != "" {
		readFile, err := os.Open(file)
		checkFileError(err)
		defer readFile.Close()

		reader = bufio.NewReader(readFile)
	}

	return reader
}

func SetWriter(file string) *os.File {

	writer := os.Stdout

	if file != "" {
		writer, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0600)

		defer writer.Close()

		checkFileError(err)
	}

	return writer
}

/*func SetBoth(files parse.Files) (*bufio.Reader, *os.File) {
	reader := setReader(files.InputFile)
	writer := setWriter(files.OutputFile)

	return reader, writer
}*/
