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
