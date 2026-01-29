package cmd

import "os"

func Exit(args []string, ctx any) error {
	os.Exit(0)
	return nil
}
