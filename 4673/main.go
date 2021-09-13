package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	nums := make([]int, 10001)

	for i := 0; i <= 9972; i++ {
		selfNumber := generate(i)
		if selfNumber <= 10000 {
			nums[selfNumber] = selfNumber
		}
	}

	for index, num := range nums {
		if index != 0 && num == 0 {
			fmt.Fprintln(w, index)
		}
	}
}

func generate(num int) int {
	if num < 10 {
		return num * 2
	}

	str := strconv.Itoa(num)
	for _, s := range str {
		n, _ := strconv.Atoi(string(s))
		num += n
	}

	return num
}
