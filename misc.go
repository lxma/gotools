package gotools

import (
	"regexp"
	"strconv"
)

func GetIntegersInString(line string) []int {
	re := regexp.MustCompile("(-?\\d+)")
	matches := re.FindAllStringSubmatch(line, -1)
	result := make([]int, len(matches))
	for i, match := range matches {
		number, _ := strconv.Atoi(match[1])
		result[i] = number
	}
	return result
}

func Future[A any](f func() A) chan A {
	ch := make(chan A, 1)
	go func() {
		ch <- f()
		close(ch)
	}()
	return ch
}
