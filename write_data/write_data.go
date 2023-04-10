package write_data

import (
	"log"
	"os"
)

type Writer interface {
	Write(b []byte) (n int, err error)
}

func CheckError(err error) {
	if err != nil {
		log.Fatalf("%v", err)
	}
}

func writeLines(writer Writer, data []string) {

	for _, line := range data {
		_, err := writer.Write([]byte(line + "\n"))
		CheckError(err)
	}
}

func WriteFile(file string, data []string) error {

	writer := os.Stdout
	if file != "" {
		writer, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0600)
		defer writer.Close()

		writeLines(writer, data)
		return err
	}
	writeLines(writer, data)
	return nil
}
