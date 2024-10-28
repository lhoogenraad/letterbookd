package authors

import (
    "bufio"
    "fmt"
    "os"
)

func readFile (filepath string) ( []byte, error ){
	file, err := os.ReadFile(filepath)
	
	if err != nil{
		return nil, err
	}

	return file, nil
}
