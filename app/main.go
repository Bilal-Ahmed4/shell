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

		if command == "" {
			continue
		} else if command == "exit" {
			os.Exit(0)
		} else if strings.HasPrefix(command, "echo ") {
			fmt.Fprintln(os.Stdout, command[5:])
			continue
		} else if tokens[0] == "type" {
			if len(tokens) < 2 {
				continue
			}
			target := tokens[1]
			path, ok := findPath(target)
			if slices.Contains(builtins, target) {
				fmt.Fprintln(os.Stdout, target+" is a shell builtin")
			} else if ok {
				fmt.Fprintln(os.Stdout, target, "is", path)
			} else {
				fmt.Fprintln(os.Stdout, target+": not found")
			}
			continue
		}
		fmt.Fprintln(os.Stdout, command+": command not found")
	}
}
