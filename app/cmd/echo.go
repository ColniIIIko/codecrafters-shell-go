package cmd

import (
	"strings"
)

func Echo(args []string, ctx any) (string, error) {
	out := strings.Join(args, " ")
	// by default applying new line at the end if it not exists
	if out[len(out)-1] != '\n' {
		out += "\n"
	}
	return out, nil
}
