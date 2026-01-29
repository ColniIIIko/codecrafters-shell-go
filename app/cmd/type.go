package cmd

import (
	"fmt"

	"github.com/codecrafters-io/shell-starter-go/app/utils"
)

func Type(args []string, ctx utils.Shell) error {
	for _, command := range args {
		if command == "" {
			continue
		}

		if _, exists := ctx[command]; exists {
			fmt.Printf("%s is a shell builtin\n", command)
		} else {
			fullPath, err := utils.ExecutablePath(command)

			if err != nil {
				fmt.Printf("%s: not found\n", command)
			} else {
				fmt.Printf("%s is %s\n", command, fullPath)
			}
		}
	}
	return nil
}
