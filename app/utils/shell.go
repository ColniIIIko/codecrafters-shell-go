package utils

type Command struct {
	Name    string
	Execute func(args []string, ctx Shell) (string, error)
}

type Shell map[string]Command
