package writedata

import "log"

type Writer interface {
	Write(b []byte) (n int, err error)
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("%v", err)
	}
}

func write(writer Writer, data []string) {

	for _, line := range data {
		_, err := writer.Write([]byte(line + "\n"))
		checkError(err)
	}
}
