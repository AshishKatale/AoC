package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}

	scanner := bufio.NewScanner(file)
	reports := make([][]int, 0, 10)
	for scanner.Scan() {
		values := strings.Fields(scanner.Text())
		s := make([]int, 0, 10)
		for _, v := range values {
			n, _ := strconv.Atoi(v)
			s = append(s, n)
		}
		reports = append(reports, s)
	}

	safeCount := 0
	for _, r := range reports {
		increasing := r[0]-r[len(r)-1] < 0
		if increasing {
			i := 0
			for ; i < len(r)-1; i++ {
				if r[i] >= r[i+1] || r[i+1]-r[i] > 3 {
					break
				}
			}
			if i >= len(r)-1 {
				safeCount++
			}
		} else {
			i := 0
			for ; i < len(r)-1; i++ {
				if r[i] <= r[i+1] || r[i]-r[i+1] > 3 {
					break
				}
			}
			if i >= len(r)-1 {
				safeCount++
			}
		}
	}
	fmt.Println(safeCount)
}
