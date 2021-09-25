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

	tasks := make([][]int, 0) // 일정
	works := make([][]int, 0) // 작업
	highProfit := 0           // 최고 수익

	for i := 0; i < count; i++ {
		var time, value int
		fmt.Fscan(r, &time, &value)

		tasks = append(tasks, []int{time, value})

		doneDate := i + tasks[i][0]
		if doneDate <= count {
			worksLen := len(works)

			for j := 0; j < worksLen; j++ {
				if works[j][0] <= i {
					profit := works[j][1] + tasks[i][1]
					works = append(works, []int{doneDate, profit})

					if highProfit < profit {
						highProfit = profit
					}
				}
			}

			profit := tasks[i][1]
			works = append(works, []int{doneDate, tasks[i][1]})

			if highProfit < profit {
				highProfit = profit
			}
		}
	}

	fmt.Fprint(w, highProfit)
}
