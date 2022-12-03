package main

import (
	"aoc-2022/lib/crypto"
	"aoc-2022/lib/input"
	"bufio"
	"flag"
	"fmt"
	"log"

	"github.com/hanson/aes"
)

func main() {

	encryptionKeyFlag := flag.String("key", "", "encryption key")
	fileNameFlag := flag.String("fileName", "", "name of file with text to be encrypted")

	flag.Parse()

	if *encryptionKeyFlag == "" || *fileNameFlag == "" {
		log.Fatalf("encryptionKeyFlag or fileNameFlag was not provided")
	}

	encryptionKey := crypto.HashKey(*encryptionKeyFlag)

	file, err := input.LoadFile(*fileNameFlag)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	inputScanner := bufio.NewScanner(file)
	inputScanner.Split(bufio.ScanLines)

	inputLines := ""
	for inputScanner.Scan() {
		inputLines += inputScanner.Text() + "\n"
	}

	inputLines = inputLines[:len(inputLines)-1]

	encryptedText := aes.AesEncryptCBC(inputLines, encryptionKey)

	fmt.Println(encryptedText)
}
