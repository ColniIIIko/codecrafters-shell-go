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

type commandInput struct {
	command string
	arg     string
}

var knownCommand map[string]command

func init() {
	knownCommand = make(map[string]command)

	knownCommand["exit"] = command{
		name: "exit",
		execute: func(arg string) error {
			os.Exit(0)
			return nil
		},
	}
	knownCommand["echo"] = command{
		name: "echo",
		execute: func(arg string) error {
			fmt.Println(arg)
			return nil
		},
	}
	knownCommand["type"] = command{
		name: "type",
		execute: func(arg string) error {
			for command := range strings.SplitSeq(arg, " ") {
				if _, exists := knownCommand[command]; exists {
					fmt.Printf("%s is a shell builtin\n", command)
				} else {
					fmt.Printf("%s: not found\n", command)
				}
			}
			return nil
		},
	}
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

func readInput(scanner *bufio.Scanner) *commandInput {
	scanner.Scan()
	input := scanner.Text()
	input = strings.Trim(input, " ")

	var command string
	var arg string

	inputArray := strings.Split(input, " ")

	if len(inputArray) == 0 {
		return nil
	}
	if len(inputArray) > 0 {
		command = inputArray[0]
		if command == "" {
			return nil
		}
	}
	if len(inputArray) > 1 {
		arg = strings.Join(inputArray[1:], " ")
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading command: %s", err)
		os.Exit(1)
	}
	return &commandInput{
		command, arg,
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		printPrompt()
		commandInput := readInput(scanner)
		if commandInput == nil {
			continue
		}
		handleCommand(commandInput.command, commandInput.arg)
	}
}
