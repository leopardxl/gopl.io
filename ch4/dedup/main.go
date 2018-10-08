package main

import (
	"bufio"
	"fmt"
	"os"
)

var m = make(map[string]int)

func main() {
	seen := make(map[string]bool) // a set of strings

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}

func k(list []string) string { return fmt.Sprintf("%q", list) }

// Add increments the number of times a particular []string has been seen
func Add(list []string) { m[k(list)]++ }

// Count returns the number of times a particular []string has been seen
func Count(list []string) int { return m[k(list)] }
