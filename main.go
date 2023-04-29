package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Command struct {
	operation int
	value     int
}

func interpretSmiley(smileyCode string) {
	var commands []Command
	var commandBuffer string

	for _, c := range smileyCode {
		if c == 'ツ' {
			commandBuffer += string(c)
		} else if c == 'ッ' {
			commandBuffer += string(c)
		} else {
			if len(commandBuffer) > 0 {
				commands = append(commands, Command{
					operation: strings.Count(commandBuffer, "ツ"),
					value:     strings.Count(commandBuffer, "ッ"),
				})
				commandBuffer = ""
			}
		}
	}

	executeCommands(commands)
}

func executeCommands(commands []Command) {

	for _, command := range commands {
		switch command.operation {
		case 3:
			fmt.Print(string(rune(command.value)))
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: smiley-lang <filename>.smiley")
		os.Exit(1)
	}

	filename := os.Args[1]

	smileyCodeBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	smileyCode := string(smileyCodeBytes)
	interpretSmiley(smileyCode)
}
