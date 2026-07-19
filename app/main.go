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
	reader := bufio.NewReader(os.Stdin)
	builtins := []string{"echo", "exit", "type"}

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

		tokens := strings.Split(command, " ")
		cmd := tokens[0]
		args := tokens[1:]

		switch {
		case cmd == "exit":
			os.Exit(0)

		case strings.HasPrefix(command, "echo "):
			fmt.Fprintln(os.Stdout, command[5:])

		case cmd == "type":
			if len(args) == 0 {
				continue
			}
			target := args[0]

			if slices.Contains(builtins, target) {
				fmt.Fprintln(os.Stdout, target+" is a shell builtin")
			} else if path, ok := findPath(target); ok {
				fmt.Fprintln(os.Stdout, target+" is "+path)
			} else {
				fmt.Fprintln(os.Stdout, target+": not found")
			}

		default:
			fmt.Fprintln(os.Stdout, command+": command not found")
		}
	}
}
