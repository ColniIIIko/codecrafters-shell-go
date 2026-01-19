package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
)

const (
	PATH_ENV = "PATH"
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

func executablePath(executable string) (string, error) {
	if isExecutable(executable) {
		return executable, nil
	}

	pathEnvValue, exists := os.LookupEnv(PATH_ENV)

	if !exists {
		return "", fmt.Errorf("PATH env not found")
	}

	for pathValue := range strings.SplitSeq(pathEnvValue, string(os.PathListSeparator)) {
		fullPath := path.Join(pathValue, executable)

		if isExecutable(fullPath) {
			return fullPath, nil
		}
	}

	return "", fmt.Errorf("%s: not found", executable)
}

func isExecutable(path string) bool {
	info, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false
	}

	mode := info.Mode()
	return mode.IsRegular() && (mode.Perm()&0111 != 0)
}

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
				if command == "" {
					continue
				}

				if _, exists := knownCommand[command]; exists {
					fmt.Printf("%s is a shell builtin\n", command)
				} else {
					fullPath, err := executablePath(command)

					if err != nil {
						fmt.Printf("%s: not found\n", command)
					} else {
						fmt.Printf("%s is %s\n", command, fullPath)
					}
				}
			}
			return nil
		},
	}

	knownCommand["pwd"] = command{
		name: "pwd",
		execute: func(arg string) error {
			pwd, err := os.Getwd()

			if err != nil {
				fmt.Printf("pwd error: %s\n", err)
				return err
			}

			fmt.Printf("%s\n", pwd)

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
	} else if execPath, err := executablePath(command); err == nil {
		execCommand := execPath

		if path.IsAbs(execPath) {
			execCommand = command
		}

		cmd := exec.Command(execCommand, strings.Split(arg, " ")...)
		out, err := cmd.Output()
		if err != nil {
			fmt.Printf("Error running %s: %s\nOutput: %s", command, err, string(out))
			return err
		}

		fmt.Print(string(out))
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
