package problems

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func convertNumber(str string) uint8 {
	nums := []string{
		"one", "two", "three", "four", "five",
		"six", "seven", "eight", "nine",
	}

	for i, v := range nums {
		if strings.HasPrefix(str, v) {
			return uint8(i) + 1
		}
	}

	return 0
}

func getNumber(str string) int {
	var firstDigit, lastDigit uint8
	for i := 0; i < len(str); i++ {
		if str[i] > 47 && str[i] < 58 {
			firstDigit = str[i] - '0'
			break
		}

		if n := convertNumber(str[i:]); n > 0 {
			firstDigit = n
			break
		}
	}
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] > 47 && str[i] < 58 {
			lastDigit = str[i] - '0'
			break
		}

		if n := convertNumber(str[i:]); n > 0 {
			lastDigit = n
			break
		}
	}
	num := firstDigit*10 + lastDigit
	return int(num)
}

func Day1() {
	file, err := os.Open("../input/day1.txt")
	defer file.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	sum := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		num := getNumber(line)
		sum += num
		fmt.Println(num, line)
	}
	fmt.Println(sum)
}
