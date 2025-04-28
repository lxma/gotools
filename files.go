package gotools

import (
    "bufio"
    "encoding/json"
    "errors"
    "fmt"
    "io"
    "os"
)

// Tests if a file exists
func FileExists(filename string) bool {
    if _, err := os.Stat(filename); err == nil {
        return true
    } else if errors.Is(err, os.ErrNotExist) {
        return false
    } else {
        panic(fmt.Errorf("cannot find out if file %s exists: %v", filename, err))
    }
}

// ReadLines reads a text file named filename and
// returns its contents a slice of strings
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

// ReadBytes reads binary contents from a file and returns it as bytes.
func ReadBytes(filename string) ([]byte, error) {
    const cBlockSize = 1048576
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var contents []byte
    buffer := make([]byte, cBlockSize)
    bufferedReader := bufio.NewReader(file)
    for {
        n, _ := bufferedReader.Read(buffer)
        contents = append(contents, buffer[:n]...)
        if n < cBlockSize {
            break
        }
    }
    return contents, nil
}

// WriteLines writes a slice of stings as text file
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

// ReadJsonFile reads a JSON file into an existing object
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

// WriteJsonFile writes an object into a JSON file
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
