package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
	// TODO: Uncomment the code below to pass the first stage

	reader := bufio.NewReader(os.Stdin)
	builtins := []string{"echo", "exit", "type"}

	for {
		fmt.Fprint(os.Stdout, "$ ")

		line, err := reader.ReadString('\n')
		if err != nil {
			os.Exit(1)
		}
		command := strings.TrimSpace(line)
		tokens := strings.Split(command, " ")

		if command == "" {
			continue
		} else if command == "exit" {
			os.Exit(0)
		} else if strings.HasPrefix(command, "echo ") {
			fmt.Fprintln(os.Stdout, command[5:])
			continue
		} else if tokens[0] == "type" && slices.Contains(builtins, tokens[1]) {
			fmt.Fprintln(os.Stdout, tokens[1]+" is a shell builtin")
			continue
		} else if tokens[0] == "type" {
			fmt.Fprintln(os.Stdout, tokens[1]+": not found")
			continue
		}
		fmt.Fprintln(os.Stdout, command+": command not found")

	}

}
