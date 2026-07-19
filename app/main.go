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
		} else if command == "exit" {
			os.Exit(0)
		} else if strings.HasPrefix(command, "echo ") {
			fmt.Fprintln(os.Stdout, command[5:])
			continue
		} else if strings.HasPrefix(command, "type ") {
			if strings.HasSuffix(command, "echo") {
				fmt.Fprintln(os.Stdout, "echo is a shell builtin")
				continue
			} else if strings.HasSuffix(command, "exit") {
				fmt.Fprintln(os.Stdout, "exit is a shell builtin")
				continue
			} else if strings.HasSuffix(command, "type") {
				fmt.Fprintln(os.Stdout, "type is a shell builtin")
				continue
			} else {
				fmt.Fprintln(os.Stdout, command[5:]+": not found")
				continue
			}
		}
		fmt.Fprintln(os.Stdout, command+": command not found")

	}

}
