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

	var n, m int
	fmt.Fscan(r, &n, &m)

	candies := make([][]int, n)

	for x := 0; x < n; x++ {
		for y := 0; y < m; y++ {
			var count int
			fmt.Fscan(r, &count)

			values := [3]int{0}

			if x != 0 {
				values[0] = candies[x-1][y]
			}
			if y != 0 {
				values[1] = candies[x][y-1]
			}
			if x != 0 && y != 0 {
				values[2] = candies[x-1][y-1]
			}

			count += max(values)

			candies[x] = append(candies[x], count)
		}
	}

	fmt.Fprint(w, candies[n-1][m-1])
}

func max(values [3]int) int {
	var t int

	for _, v := range values {
		if t < v {
			t = v
		}
	}

	return t
}
