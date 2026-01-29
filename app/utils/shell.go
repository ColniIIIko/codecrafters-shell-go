package utils

type Command struct {
	Name    string
	Execute func(args []string, ctx Shell) error
}

type Shell map[string]Command
