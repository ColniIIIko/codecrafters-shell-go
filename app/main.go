package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("$ ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	command := scanner.Text()

	if err := scanner.Err(); err != nil {
		fmt.Print("fail to read input")
		return
	}

	fmt.Printf("%s: command not found\n", command)
}
