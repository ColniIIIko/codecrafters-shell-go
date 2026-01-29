package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/app/cmd"
	shell "github.com/codecrafters-io/shell-starter-go/app/cmd"
	"github.com/codecrafters-io/shell-starter-go/app/utils"
)

const QUOTE = '\''
const DOUBLE_QUOTE = '"'

type commandInput struct {
	command string
	arg     string
}

var knownCommand utils.Shell

func init() {
	knownCommand = make(utils.Shell)

	knownCommand["exit"] = utils.Command{
		Name: "exit",
		Execute: func(args []string, ctx utils.Shell) error {
			return shell.Exit(args, ctx)
		},
	}
	knownCommand["echo"] = utils.Command{
		Name: "echo",
		Execute: func(args []string, ctx utils.Shell) error {
			return shell.Echo(args, ctx)
		},
	}
	knownCommand["type"] = utils.Command{
		Name: "type",
		Execute: func(args []string, ctx utils.Shell) error {
			return shell.Type(args, ctx)
		},
	}

	knownCommand["pwd"] = utils.Command{
		Name: "pwd",
		Execute: func(args []string, ctx utils.Shell) error {
			return cmd.Pwd(args, ctx)
		},
	}

	knownCommand["cd"] = utils.Command{
		Name: "cd",
		Execute: func(args []string, ctx utils.Shell) error {
			return cmd.Cd(args, ctx)
		},
	}
}

func isQuote(char byte) bool {
	return char == QUOTE || char == DOUBLE_QUOTE
}

func parseArg(arg string) []string {
	quoteType := byte(QUOTE)
	insideQuotes := false
	groups := make([]string, 0)

	group := ""

	index := 0

	for index < len(arg) {
		if isQuote(arg[index]) && index+1 != len(arg) && arg[index+1] == arg[index] {
			index += 2
			continue
		}

		if !insideQuotes && isQuote(arg[index]) {
			quoteType = arg[index]
			insideQuotes = true
		} else if insideQuotes && isQuote(arg[index]) && arg[index] == quoteType {
			if group != "" {
				groups = append(groups, group)
				group = ""
			}
			insideQuotes = false
		} else if insideQuotes || arg[index] != ' ' {
			group += string(arg[index])
		} else if !insideQuotes && arg[index] == ' ' && len(group) > 0 {
			groups = append(groups, group)
			group = ""
		}

		index += 1
	}

	if group != " " && group != "" {
		groups = append(groups, group)
	}

	if insideQuotes {
		// has only opening quote
		return []string{arg}
	}

	return groups
}

func handleCommand(command string, args []string) error {
	if cmd, exists := knownCommand[command]; exists {

		if err := cmd.Execute(args, knownCommand); err != nil {
			fmt.Printf("Error executing command: %s", err)
			return err
		}
	} else if execPath, err := utils.ExecutablePath(command); err == nil {
		execCommand := execPath

		if path.IsAbs(execPath) {
			execCommand = command
		}

		cmd := exec.Command(execCommand, args...)
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
		args := parseArg(commandInput.arg)
		handleCommand(commandInput.command, args)
	}
}
