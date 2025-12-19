package services

import (
	"encoding/csv"
	"fmt"
	"os"
)

func exportCSV(filePath string, data [][]string) error {
	file, err := os.Create(filePath)

	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", filePath, err)
	}

	defer file.Close()

	writer := csv.NewWriter(file)

	defer writer.Flush()

	return writer.WriteAll(data)
}

func importCSV(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)

	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", filePath, err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	return reader.ReadAll()
}
