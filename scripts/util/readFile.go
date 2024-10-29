package util

import (
	"bufio"
	"fmt"
	"os"
)

func GetScanner (filepath string) ( *bufio.Scanner, error ) {
	file, err := os.Open(filepath)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	
	// Set max read in len

	maxIn := 5000*5000
    buf := make([]byte, 0, maxIn)
	scanner.Buffer(buf, maxIn)

	if err != nil{
		fmt.Println(err)
		return scanner, err
	}

	return scanner, nil
}

