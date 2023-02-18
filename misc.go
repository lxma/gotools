package gotools

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

func ReadLines(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var resultLines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		resultLines = append(resultLines, scanner.Text())
	}
	return resultLines
}

func WriteLines(filename string, lines []string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("failed writing to file %s: %s", filename, err)
	}
	writer := bufio.NewWriter(file)

	for _, line := range lines {
		_, _ = writer.WriteString(line + "\n")
	}

	writer.Flush()
	file.Close()
}

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
