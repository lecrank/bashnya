package write_data

import (
	"os"
)

func writeLines(writer *os.File, data []string) error {

	for _, line := range data {
		_, err := writer.Write([]byte(line + "\n"))
		if err != nil {
			return err
		}
	}
	return nil
}

func WriteFile(file string, data []string) error {
	var err error

	writer := os.Stdout
	if file != "" {
		writer, err = os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0600)
		defer writer.Close()

		if err != nil {
			return err
		}
		err = writeLines(writer, data)
		return err
	}
	err = writeLines(writer, data)
	return err
}
