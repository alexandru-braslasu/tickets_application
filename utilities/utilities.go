package utilities

import (
	"bufio"
	"encoding/json"
	"errors"
	"math/rand"
	"os"
	"time"
)

type Utility struct {
	InputFilePath string
	OutputFilePath string
	From string
	To string
}

var arrayFrom = []string{"Paris", "Bucharest", "Berlin", "Athens"}
var arrayTo = []string{"Helsinki", "Brussels", "Dublin", "Amsterdam"}

func (fm Utility) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)

	if err != nil {
		return nil, errors.New("Failed to open file.")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		return nil, errors.New("Failed to read line in file.")
	}

	return lines, nil
}

func (fm Utility) WriteResult(data any) error {
	file, err := os.Create(fm.OutputFilePath)
	
	if err != nil {
		return errors.New("Failed to create file.")
	}

	defer file.Close()

	time.Sleep(3 * time.Second)

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		return errors.New("Failed to convert data to JSON.")
	}

	return nil
}

func New(inputPath, outputPath string) Utility {
	return Utility {
		InputFilePath: inputPath,
		OutputFilePath: outputPath,
		From: pickRandomCity(arrayFrom),
		To: pickRandomCity(arrayTo),
	}
}

func pickRandomCity(capitals []string) string {
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(capitals))
	randomCapital := capitals[randomIndex]
	return randomCapital
}

func (fm Utility) GetCapitalFrom() string {
	return fm.From
}

func (fm Utility) GetCapitalTo() string {
	return fm.To
}