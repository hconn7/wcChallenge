package sources

import (
	"log"
	"os"
)

func Readfile(fileName string) ([]byte, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return data, nil
}
