package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func readFrom(file string) []string {
	var data []string

	// set default reader
	reader := bufio.NewReader(os.Stdin)

	// set file reader
	if file != "stdin" {
		readFile, err := os.Open(file)
		if err != nil {
			panic("There is no such file")
		}
		defer readFile.Close()

		reader = bufio.NewReader(readFile)
	}
	// read lines
	for {
		line, err := reader.ReadString('\n')
		data = append(data, strings.TrimSuffix(line, "\n"))
		if err == io.EOF {
			break
		} else {
			checkError(err)
		}
	}
	return data
}

func writeInto(file string, data []string) {
	writeFile, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0600)

	defer writeFile.Close()

	checkError(err)

	for _, line := range data {
		_, err := writeFile.Write([]byte(line + "\n"))
		checkError(err)
	}
}

func writeStdout(data []string) {
	fmt.Println("\n----OUTPUT----")
	for _, line := range data {
		fmt.Println(line)
	}
}
