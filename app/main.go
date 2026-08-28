package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
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
var _DEBUG = false

func debug(format string, a ...any) {
	debugPrefix := "[DEBUG] "

	if _DEBUG {
		fmt.Printf(debugPrefix+format, a...)
	}
}

func init() {
	knownCommand = make(utils.Shell)

	knownCommand["exit"] = utils.Command{
		Name: "exit",
		Execute: func(args []string, ctx utils.Shell) (string, error) {
			return shell.Exit(args, ctx)
		},
	}
	knownCommand["echo"] = utils.Command{
		Name: "echo",
		Execute: func(args []string, ctx utils.Shell) (string, error) {
			return shell.Echo(args, ctx)
		},
	}
	knownCommand["type"] = utils.Command{
		Name: "type",
		Execute: func(args []string, ctx utils.Shell) (string, error) {
			return shell.Type(args, ctx)
		},
	}
	knownCommand["pwd"] = utils.Command{
		Name: "pwd",
		Execute: func(args []string, ctx utils.Shell) (string, error) {
			return cmd.Pwd(args, ctx)
		},
	}
	knownCommand["cd"] = utils.Command{
		Name: "cd",
		Execute: func(args []string, ctx utils.Shell) (string, error) {
			return cmd.Cd(args, ctx)
		},
	}
}

func handleCommand(command string, args []string) (string, error) {
	debug("command=%s args=%s\n", command, strings.Join(args, ", "))

	if cmd, exists := knownCommand[command]; exists {
		return cmd.Execute(args, knownCommand)
	} else if execPath, err := utils.ExecutablePath(command); err == nil {
		execCommand := execPath

		if path.IsAbs(execPath) {
			execCommand = command
		}

		cmd := exec.Command(execCommand, args...)
		stderr, _ := cmd.StderrPipe()
		stdout, _ := cmd.StdoutPipe()

		err := cmd.Start()

		if err != nil {
			debug("Error running %s: %s\n", command, err)
		}

		out, _ := io.ReadAll(stdout)
		errOut, _ := io.ReadAll(stderr)

		err = cmd.Wait()

		if err != nil {
			debug("Error running %s: %s\n", command, err)
		}

		return string(out), errors.New(string(errOut))
	}

	return fmt.Sprintf("%s: command not found", command), nil
}

func printPrompt() {
	fmt.Print("$ ")
}

func readInput(scanner *bufio.Scanner) *core.CommandInput {
	scanner.Scan()
	input := scanner.Text()
	debug("raw input: %s\n", input)
	input = strings.Trim(input, " ")

	ci := core.ParseInput(input)

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading command: %s", err)
		os.Exit(1)
	}

	return ci
}

func redirectOutput(output string, redirect core.RedirectOutput, to core.RedirectConsumer) {
	path := utils.ResolvePath(string(to))
	debug("Redirect path=%s\n", path)
	err := os.WriteFile(path, []byte(output), 0644)

	if err != nil {
		fmt.Printf("%s\n", err)
	}
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "--debug" {
		_DEBUG = true
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		printPrompt()
		commandInput := readInput(scanner)
		if commandInput == nil {
			continue
		}
		out, err := handleCommand(commandInput.Command, commandInput.Args)

		debug("Command Input %s\n", commandInput)
		debug("Command Output out=%s, err=%s\n", out, err)

		if commandInput.Redirect != "" {
			redirectOutput(out, commandInput.Redirect, commandInput.RedirectTo)
		} else {
			fmt.Print(out)
		}

		if err != nil {
			fmt.Println(err)
		}
	}
}
