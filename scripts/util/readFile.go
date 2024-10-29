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
	
	if err != nil{
		fmt.Println(err)
		return scanner, err
	}


	return scanner, nil
}

