package cmd

import (
	"fmt"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/app/utils"
)

func Type(args []string, ctx utils.Shell) (string, error) {
	outputs := make([]string, 0)

	for _, command := range args {
		if command == "" {
			continue
		}

		if _, exists := ctx[command]; exists {
			outputs = append(outputs, fmt.Sprintf("%s is a shell builtin", command))
		} else {
			fullPath, err := utils.ExecutablePath(command)

			if err != nil {
				outputs = append(outputs, fmt.Sprintf("%s: not found", command))
			} else {
				outputs = append(outputs, fmt.Sprintf("%s is %s", command, fullPath))
			}
		}
	}
	out := strings.Join(outputs, "\n")

	if len(out) <= 0 {
		return "", nil
	}

	return out, nil
}
