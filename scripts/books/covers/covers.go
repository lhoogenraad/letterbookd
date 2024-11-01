package covers

import (
	"os"
	"strings"
)

func getFilename(filepath string) string {
	filepathSplit := strings.Split(filepath, "/")
	filename := filepathSplit[len(filepathSplit)-1]
	return filename
}

/**
Takes in a filepath for an img file, uploads the file to the
file hosting service, then returns the URL to the file
*/
func UploadCoverAndGetURL(filepath string) (string, error) {
	filename := getFilename(filepath)
	savePath := `/home/leon/Documents/letterbookd_files/` + filename
	err := os.Rename(filepath, savePath)
	if err != nil {
		return "", nil
	}

	return savePath, nil
}
