package readdata

import (
	"io"
	"log"
	"strings"
)

type Reader interface {
	ReadString(delim byte) (string, error)
}

func read(rd Reader) []string {
	data := make([]string, 0)
	// read lines
	for {
		line, err := rd.ReadString('\n')
		data = append(data, strings.TrimSuffix(line, "\n"))
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal("Error while reading file occurred")
		}
	}
	return data
}
