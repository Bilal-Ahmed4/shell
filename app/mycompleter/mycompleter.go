package mycompleter

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

//	type AutoCompleter interface {
//	    Do(line []rune, pos int) (newLine [][]rune, length int)
//	}
//
// we have to overwrite Do method to implement AutoCompleter interface
// line — everything the user has typed so far (as runes, not a plain string)
// pos — the cursor position

// And expects back:

// newLine — the possible completions (as a list of rune-slices)
// length — how many characters of the current input should be considered "replaced" by the completion

type MyCompleter struct {
	Builtins   []string
	lastPrefix string
	tabPressed bool
}

func (c *MyCompleter) Do(line []rune, pos int) (newLine [][]rune, length int) {
	prefix := string(line[:pos]) // get the prefix and search for matching words
	matches := autoComplete(prefix, c.Builtins)
	matches = append(matches, autoCompleteExecutables(prefix)...) //... unpacks the slice into individual elements
	//removing duplicates
	matches = removeDuplicates(matches)

	// sort the matches alphabetically
	sort.Strings(matches)

	if len(matches) == 0 {
		fmt.Print("\x07") // ring the bell — no matches
		c.lastPrefix = ""
		c.tabPressed = false
		return nil, 0
	}

	//when there are exactly one match add a trailing space to indicate completion
	if len(matches) == 1 {
		completion := matches[0][len(prefix):] + " " // ← add trailing space
		c.lastPrefix = ""
		c.tabPressed = false
		return [][]rune{[]rune(completion)}, len(prefix)
	}

	if len(matches) >= 2 {
		if c.tabPressed && c.lastPrefix == prefix {
			fmt.Println()
			for _, match := range matches {
				fmt.Fprint(os.Stdout, match, "  ")
			}
			fmt.Println()
			fmt.Fprint(os.Stdout, "$ ", c.lastPrefix)
			c.lastPrefix = ""
			c.tabPressed = false
			return nil, 0
		} else {
			fmt.Print("\x07")
			c.lastPrefix = prefix
			c.tabPressed = true
			return nil, 0
		}

	}
	return nil, 0
}

func autoComplete(line string, builtins []string) []string {
	var words []string
	for _, word := range builtins {
		if strings.HasPrefix(word, line) {
			words = append(words, word)
		}
	}

	return words
}

func removeDuplicates(words []string) []string {
	seen := make(map[string]bool)
	var result []string
	for _, word := range words {
		if !seen[word] {
			seen[word] = true
			result = append(result, word)
		}
	}
	return result
}

func autoCompleteExecutables(line string) []string {
	PATH := os.Getenv("PATH")
	pathDirs := strings.Split(PATH, string(os.PathListSeparator))
	var executables []string
	for _, dir := range pathDirs {
		entries, err := os.ReadDir(dir)
		if err != nil {
			continue
		}
		for _, entry := range entries {
			if !strings.HasPrefix(entry.Name(), line) {
				continue
			}
			info, err := entry.Info()
			if err != nil {
				continue
			}
			if info.Mode()&0111 != 0 {
				executables = append(executables, entry.Name())
			}
		}
	}
	return executables
}
