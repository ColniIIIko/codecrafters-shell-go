package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type command struct {
	name    string
	execute func() error
}

var knownCommand = map[string]command{
	"exit": {
		name: "exit",
		execute: func() error {
			os.Exit(0)
			return nil
		},
	},
}

func handleCommand(command string) error {
	if cmd, exists := knownCommand[command]; exists {
		if err := cmd.execute(); err != nil {
			fmt.Printf("Error executing command: %s", err)
			return err
		}

	} else {
		fmt.Printf("%s: command not found\n", command)
	}

	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("$ ")
	for scanner.Scan() {
		command := scanner.Text()
		command = strings.Trim(command, " ")

		if err := scanner.Err(); err != nil {
			fmt.Printf("Error reading command: %s", err)
			os.Exit(1)
		}
		handleCommand(command)

		fmt.Print("$ ")
	}
}
