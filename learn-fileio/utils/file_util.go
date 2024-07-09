package utils

import (
	"fmt"
	"os"
)

func ReadFromFile(fileName string) string {
	data, err := os.ReadFile(fileName)
    if err != nil {
        fmt.Println("Error: file not found!")
        return ""
    }
	dataString := string(data)
	return dataString
}

func WriteToFile(fileName, text string) {
	data := fmt.Sprint(text)
	os.WriteFile(fileName, []byte(data), 0644)
}
