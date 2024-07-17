package fileutil

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

func ReadLinesFromTextFile(path string) ([]string, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, errors.New("Error opening file")
	}

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		file.Close()
		return nil, errors.New("Error reading line from file")
	}

	file.Close()
	return lines, nil
}

func WriteJson(path string, data map[string]string) error {
	file, err := os.Create(path)

	if err != nil {
		return errors.New("Failed to create json file")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		return errors.New("Failed to write/encode to json")
	}
	return nil
}
