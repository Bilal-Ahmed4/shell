package main

import (
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"

	"github.com/chzyer/readline"
	"github.com/codecrafters-io/shell-starter-go/app/mycompleter"
)

// findPath searches for an executable in the system PATH.
// Returns the full path and true if found, empty string and false otherwise.
func findPath(target string) (string, bool) {
	path, err := exec.LookPath(target)
	if err != nil {
		return "", false
	}
	return path, true
}

// autoComplete completes the user input based on the available commands.
// also handles tab completion for built-in commands.
// it also handle the case if the command is not builtin and give sound feeback(invalid command)\

// main runs the shell REPL (Read-Eval-Print Loop).
// It reads user input, parses commands, and executes them.
func main() {
	// Tab completion: pressing Tab will autocomplete "echo", "exit", "type"
	// // Built-in commands that the shell handles directly (not from PATH)
	builtins := []string{"echo", "exit", "type"}
	completer := &mycompleter.MyCompleter{Builtins: builtins}

	// Create a readline instance with prompt "$ " and tab completion enabled.
	// readline handles raw terminal input, history, and cursor movement.
	reader, err := readline.NewEx(&readline.Config{
		Prompt:          "$ ",
		AutoComplete:    completer,
		InterruptPrompt: "^C",
	})
	if err != nil {
		os.Exit(1)
	}
	defer reader.Close()

	for {
		// Read a line of input from the user
		line, err := reader.Readline()
		if err != nil {
			os.Exit(1)
		}

		// Trim whitespace and split into tokens
		command := strings.TrimSpace(line)
		tokens := strings.Split(command, " ")

		// Empty input: just show the prompt again
		if command == "" {
			continue

			// "exit" command: terminate the shell
		} else if command == "exit" {
			os.Exit(0)

			// "echo" command: print everything after "echo "
		} else if strings.HasPrefix(command, "echo ") {
			fmt.Fprintln(os.Stdout, command[5:])
			continue

			// "type" command: check if a command is a builtin or found in PATH
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

			// External command: find in PATH and execute
		} else if path, ok := findPath(tokens[0]); ok {
			// exec.Command needs the full path to the executable
			// We pass the original command name as Args[0] for the tester
			cmd := exec.Command(path, tokens[1:]...)
			cmd.Args[0] = tokens[0]
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
			continue
		}

		// Command not found
		fmt.Fprintln(os.Stdout, command+": command not found")
	}
}
