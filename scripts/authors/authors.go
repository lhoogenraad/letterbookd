package authors

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadAndUpload () error {
	filepath := `/home/leon/Downloads/ol_dump_authors_2024-09-30.txt`
	file, err := os.Open(filepath)
	scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)
	
	if err != nil{
		fmt.Println(err)
		return err
	}

	i := 0
	max := 100
	for scanner.Scan() && i < max {
		text := scanner.Text()
		textSplit := strings.Split(text, "{")
		textSplit = textSplit[1:]
		cleaned := strings.Join(textSplit, "{")
		fmt.Println(i, "\n", text, "\n", cleaned, "\n\n")
		i++
	}

	return nil
}

