package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type command struct {
	name    string
	execute func(arg string) error
}

var knownCommand = map[string]command{
	"exit": {
		name: "exit",
		execute: func(arg string) error {
			os.Exit(0)
			return nil
		},
	},
	"echo": {
		name: "echo",
		execute: func(arg string) error {
			fmt.Println(arg)
			return nil
		},
	},
}

func handleCommand(command string, arg string) error {
	if cmd, exists := knownCommand[command]; exists {
		if err := cmd.execute(arg); err != nil {
			fmt.Printf("Error executing command: %s", err)
			return err
		}
	} else {
		fmt.Printf("%s: command not found\n", command)
	}

	return nil
}

func printPrompt() {
	fmt.Print("$ ")
}

func readInput(scanner *bufio.Scanner) (string, string) {
	scanner.Scan()
	input := scanner.Text()
	input = strings.Trim(input, " ")

	var command string
	var arg string

	sepIndex := strings.Index(input, " ")
	if sepIndex == -1 {
		command = input
		arg = ""
	} else {
		command = input[:sepIndex]
		arg = input[(sepIndex + 1):]
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading command: %s", err)
		os.Exit(1)
	}
	return command, arg
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		printPrompt()
		command, arg := readInput(scanner)
		handleCommand(command, arg)
	}
}
