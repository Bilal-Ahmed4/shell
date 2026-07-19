package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
	// TODO: Uncomment the code below to pass the first stage

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stdout, "$ ")

		line, err := reader.ReadString('\n')
		if err != nil {
			os.Exit(1)
		}
		command := strings.TrimSpace(line)

		if command == "" {
			continue
		}
		if command == "exit" {
			os.Exit(0)
		}

		fmt.Fprintln(os.Stdout, command+": command not found")
	}

}
