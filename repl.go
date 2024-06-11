package main

import (
	"bufio"
	"fmt"
	"os"
)

func repl() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		command := scanner.Text()

		if len(command) == 0 {
			continue

		}

		commandName := command
		cmd, ok := getCommands()[commandName]

		if ok {
			err := cmd.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unkown command")
			continue
		}
	}
}
