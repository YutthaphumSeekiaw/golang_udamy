package filemanager

import (
	"bufio"
	"encoding/json"
	"os"
)

func ReadLines() ([]string, error) {
	lines := []string{}
	file, err := os.Open("pricess.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

func WriteJson(path string, data interface{}) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	json.NewEncoder(file).Encode(data)
	return nil
}
