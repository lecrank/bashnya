package write_data

import (
	"fmt"
	"log"
	"os"
)

type Writer interface {
	Write(b []byte) (n int, err error)
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("%v", err)
	}
}

func writeLines(writer Writer, data []string) {

	for _, line := range data {
		_, err := writer.Write([]byte(line + "\n"))
		checkError(err)
	}
}

func WriteFile(file string, data []string) {

	if file != "" {
		writer, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0600)
		defer writer.Close()

		writeLines(writer, data)

		checkError(err)
	} else {
		fmt.Println("\n---OUTPUT---")
		writer := os.Stdout
		writeLines(writer, data)
	}
}
