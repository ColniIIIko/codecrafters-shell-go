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
	"github.com/codecrafters-io/shell-starter-go/app/core"
	"github.com/codecrafters-io/shell-starter-go/app/utils"
)

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

func readInput(scanner *bufio.Scanner) *core.CommandInput {
	scanner.Scan()
	input := scanner.Text()
	input = strings.Trim(input, " ")

	ci := core.ParseInput(input)

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading command: %s", err)
		os.Exit(1)
	}

	return ci
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		printPrompt()
		commandInput := readInput(scanner)
		if commandInput == nil {
			continue
		}
		args := core.ParseArg(commandInput.Arg)
		handleCommand(commandInput.Command, args)
	}
}
