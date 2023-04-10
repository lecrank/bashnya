package read_data

import (
	"bufio"
	"io"
	"os"
	"strings"
)

type Reader interface {
	ReadString(delim byte) (string, error)
}

func readLines(rd Reader) ([]string, error) {
	lines := make([]string, 0)
	// read lines
	for {
		line, err := rd.ReadString('\n')

		lines = append(lines, strings.TrimSuffix(line, "\n"))

		if err == io.EOF {
			break
		} else if err != nil {
			return lines, err
		}
	}
	return lines, nil
}

func ReadFile(file string) ([]string, error) {

	// set default reader
	reader := bufio.NewReader(os.Stdin)

	// set file reader
	if file != "" {
		readFile, err := os.Open(file)
		defer readFile.Close()

		if err != nil {
			return []string{""}, err
		}

		reader = bufio.NewReader(readFile)
	}
	return readLines(reader)
}
