package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var count int
	fmt.Fscan(r, &count)

	words := make([]string, 0)
	for i := 0; i < count; i++ {
		var word string
		fmt.Fscan(r, &word)
		if !contains(words, word) {
			words = append(words, word)
		}
	}

	sort.Slice(words, func(a, b int) bool {
		if len(words[a]) < len(words[b]) {
			return true
		} else if len(words[a]) > len(words[b]) {
			return false
		} else {
			return words[a] <= words[b]
		}
	})

	for _, word := range words {
		fmt.Fprint(w, word, "\n")
	}
}

func contains(arr []string, target string) bool {
	for _, elmt := range arr {
		if elmt == target {
			return true
		}
	}

	return false
}
