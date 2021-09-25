package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var count int
	fmt.Fscan(r, &count)

	sequence := make([]int, 0) // 파도반 수열
	sequence = append(sequence, 1, 1, 1)

	for i := 0; i < count; i++ {
		var n int
		fmt.Fscan(r, &n)

		sequenceLen := len(sequence)

		if n <= sequenceLen {
			fmt.Fprintln(w, sequence[n-1])
			continue
		}

		for j := sequenceLen; j < n; j++ {
			sequence = append(sequence, sequence[j-3]+sequence[j-2])
		}

		fmt.Fprintln(w, sequence[n-1])
	}
}
