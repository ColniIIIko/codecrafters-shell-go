package cmd

import (
	"strings"
)

func Echo(args []string, ctx any) (string, error) {
	out := strings.Join(args, " ")
	return out, nil
}
