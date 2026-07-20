package mycompleter

import (
	"fmt"
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
	Builtins []string
}

func (c *MyCompleter) Do(line []rune, pos int) (newLine [][]rune, length int) {
	prefix := string(line[:pos]) // get the prefix and search for matching words
	matches := autoComplete(prefix, c.Builtins)

	if len(matches) == 0 {
		fmt.Print("\x07") // ring the bell — no matches
		return nil, 0
	}

	var result [][]rune
	for _, m := range matches {
		completion := m[len(prefix):] // only the remaining part after what's typed
		result = append(result, []rune(completion))
	}
	return result, len(prefix)
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
