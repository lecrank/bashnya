package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func check_error(err error) {
	if err != nil {
		panic(err)
	}
}

func read_from(file string) []string {
	var data []string

	// set default reader
	reader := bufio.NewReader(os.Stdin)

	// set file reader
	if file != "stdin" {
		read_file, err := os.Open(file)
		if err != nil {
			panic("There is no such file")
		}
		defer read_file.Close()

		reader = bufio.NewReader(read_file)
	}
	// read lines
	for {
		line, err := reader.ReadString('\n')
		data = append(data, strings.TrimSuffix(line, "\n"))
		if err == io.EOF {
			break
		} else {
			check_error(err)
		}
	}
	return data
}

func write_into(file string, data []string) {
	write_file, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0600)

	defer write_file.Close()

	check_error(err)

	for _, line := range data {
		_, err := write_file.Write([]byte(line + "\n"))
		check_error(err)
	}
}

/*func write_into(file string, data []string) {
	// set default reader
	writer := bufio.NewWriter(os.Stdout)

	// set file reader
	if file != "stdout" {
		write_file, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0600)
		defer write_file.Close()
		check_error(err)

		writer = bufio.NewWriter(write_file)
	}
	// read lines
	for _, line := range data {
		_, err := writer.WriteString(line)
		check_error(err)
	}
}*/

func write_stdout(data []string) {
	fmt.Println("----OUTPUT----")
	for _, line := range data {
		fmt.Println(line)
	}
}
