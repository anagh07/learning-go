package fileutil

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutPutFilePath string
}

func (fileManager FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fileManager.InputFilePath)

	if err != nil {
		return nil, errors.New("error opening file")
	}

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		file.Close()
		return nil, errors.New("error reading line from file")
	}

	file.Close()
	return lines, nil
}

func (fileManager FileManager) WriteResult(data any) error {
	file, err := os.Create(fileManager.OutPutFilePath)

	if err != nil {
		return errors.New("failed to create json file")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		return errors.New("failed to write/encode to json")
	}
	return nil
}

func New(inputFilePath, outputFilePath string) FileManager {
	return FileManager{
		InputFilePath:  inputFilePath,
		OutPutFilePath: outputFilePath,
	}
}
