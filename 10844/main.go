package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	dp := [100][10]int{}

	for i := 1; i <= 9; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][1] % 1000000000
		dp[i][9] = dp[i-1][8] % 1000000000
		for j := 1; j < 9; j++ {
			dp[i][j] = (dp[i-1][j-1] + dp[i-1][j+1]) % 1000000000
		}
	}

	count := 0
	for i := 0; i < 10; i++ {
		count = (count + dp[n-1][i]) % 1000000000
	}
	fmt.Print(count)
}
