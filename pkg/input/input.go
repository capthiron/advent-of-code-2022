package input

import (
	"aoc-2022/pkg/crypto"
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/hanson/aes"
)

func ReadInput() ([]string, error) {

	fileNameFlag := flag.String("fileName", "", "Name of the input file")
	decryptKeyFlag := flag.String("key", "", "Key to decrypt an encrypted input file")

	flag.Parse()

	var decryptionKey string
	if *decryptKeyFlag != "" {
		decryptionKey = crypto.HashKey(*decryptKeyFlag)
	}

	if *fileNameFlag == "" {
		flag.Usage()
		return nil, errors.New("fileName is required")
	}

	readInput, err := LoadFile(*fileNameFlag)
	if err != nil {
		return nil, err
	}
	defer readInput.Close()

	inputScanner := bufio.NewScanner(readInput)
	inputScanner.Split(bufio.ScanLines)

	var inputLines []string
	for inputScanner.Scan() {
		inputLines = append(inputLines, inputScanner.Text())
	}

	if decryptionKey != "" {
		inputLines = strings.Split(aes.AesDecryptCBC(inputLines[0], decryptionKey), "\n")
	}

	return inputLines, nil
}

func LoadFile(fileName string) (*os.File, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to open input file '%v'", fileName)
	}
	return file, nil
}
