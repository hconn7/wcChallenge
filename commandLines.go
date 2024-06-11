package main

import (
	"bytes"
	"fmt"
	"github.com/hconn7/wcChallenge/sources"
	"log"
)

func commandLines() error {

	data, err := sources.Readfile("./sources/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := bytes.Split(data, []byte("\n"))

	lineCount := len(lines) - 1

	fmt.Println("The numbers of lines are:", lineCount)
	return nil
}
