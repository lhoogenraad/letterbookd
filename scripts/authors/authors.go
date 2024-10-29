package authors

import (
	"encoding/json"
	"scripts/util"
	"strings"
	"fmt"
)

func getLineAsJSON (text string) map[string]interface{} {
	textSplit := strings.Split(text, "{")
	textSplit = textSplit[1:]
	cleaned := strings.Join(textSplit, "{")
	cleaned = "{" + cleaned
	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(cleaned), &jsonMap)

	return jsonMap
}

func ReadAndUpload () error {
	filepath := `/home/leon/Downloads/ol_dump_authors_2024-09-30.txt`
	scanner, err := util.GetScanner(filepath)

	if err != nil{
		fmt.Println(err)
		return err
	}

	i := 0
	max := 100000
	for scanner.Scan() && i < max {
		jsonMap := getLineAsJSON(scanner.Text())
		name := jsonMap["name"]
		dob := jsonMap["birth_date"]
		if name != nil && dob != nil {
			fmt.Println(name, "\n", dob, "\n", "\n\n")
		}
		i++
	}

	return nil
}

