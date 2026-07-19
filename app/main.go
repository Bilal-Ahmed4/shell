package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"
)

func findPath(target string) (string, bool) {
	path, err := exec.LookPath(target)
	if err != nil {
		return "", false
	}
	return path, true
}

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
		path, ok := findPath(tokens[1])

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
		} else if tokens[0] == "type" && ok {
			fmt.Fprintln(os.Stdout, tokens[1], "is", path)
			continue
		} else if tokens[0] == "type" {
			fmt.Fprintln(os.Stdout, tokens[1]+": not found")
			continue
		}
		fmt.Fprintln(os.Stdout, command+": command not found")

	}

}
