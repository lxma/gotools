package gotools

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
)

func FileExists(filename string) bool {
	if _, err := os.Stat(filename); err == nil {
		return true
	} else if errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		panic(fmt.Errorf("cannot find out if file %s exists: %v", filename, err))
	}
}

func ReadLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var resultLines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		resultLines = append(resultLines, scanner.Text())
	}
	return resultLines, nil
}

func WriteLines(filename string, lines []string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(fmt.Errorf("failed writing to file %s: %s", filename, err))
	}
	writer := bufio.NewWriter(file)

	for _, line := range lines {
		_, _ = writer.WriteString(line + "\n")
	}

	writer.Flush()
	file.Close()
}

func ReadJsonFile(filename string, data any) error {
	jsonFile, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	bytes, _ := io.ReadAll(jsonFile)
	err = json.Unmarshal(bytes, data)
	return err
}

func WriteJsonFile(filename string, data any) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(fmt.Errorf("failed to marshal JSON content: %v", err))
	}
	jsonFile, err := os.Create(filename)
	if err != nil {
		panic(fmt.Errorf("failed to open JSON file %s for writing: %v", filename, err))
	}
	jsonFile.Write(bytes)
	jsonFile.Close()
}
