package main

import (
	"fmt"
	"github.com/hconn7/wcChallenge/sources"
	"log"
	"strings"
)

func commandWords() error {
	totalWords := []string{}
	data, err := sources.Readfile("./sources/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	myStrings := string(data)
	completeType := strings.Fields(myStrings)

	for _, word := range completeType {
		totalWords = append(totalWords, word)
	}
	fmt.Println(len(totalWords))
	return nil
}
