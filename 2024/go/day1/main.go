package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}

	var arr1, arr2 []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		values := strings.Fields(scanner.Text())

		val1, _ := strconv.Atoi(values[0])
		val2, _ := strconv.Atoi(values[1])
		arr1 = append(arr1, val1)
		arr2 = append(arr2, val2)
	}

	slices.Sort(arr1)
	slices.Sort(arr2)

	sum := 0
	for i := 0; i < len(arr1); i++ {
		d := arr1[i] - arr2[i]
		if d < 0 {
			d *= -1
		}
		sum += d
	}

	fmt.Println(sum)
}
