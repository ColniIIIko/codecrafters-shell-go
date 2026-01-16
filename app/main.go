package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("$ ")
	for scanner.Scan() {
		command := scanner.Text()

		if err := scanner.Err(); err != nil {
			fmt.Print("fail to read input")
			os.Exit(1)
		}

		fmt.Printf("%s: command not found\n", command)
		fmt.Print("$ ")
	}
}
