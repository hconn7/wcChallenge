package main

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"-c": {
			name:        "number of bytes",
			description: "Find the number of bytes in a text file",
			callback:    commandBytes,
		},

		"-l": {
			name:        "number of lines",
			description: "Outputs the number of lines in a text",
			callback:    commandLines,
		},

		"-m": {
			name:        "number of characters",
			description: "find the number od characters in the text",
			callback:    commandChars,
		},

		"-w": {
			name:        "number of words",
			description: "find the total words!",
			callback:    commandWords,
		},
	}
}
