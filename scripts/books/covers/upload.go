package covers

import (
	"fmt"
	"os"
	"scripts/util"
	"strings"
)

func getFilename(filepath string) string {
	filepathSplit := strings.Split(filepath, "/")
	filename := filepathSplit[len(filepathSplit)-1]
	return filename
}

func SetBookUrl(bookId int, coverURL string) error {
	updateQuery := `UPDATE books SET cover_url = ? WHERE id = ?`
	_, err := util.DB.Exec(updateQuery, coverURL, bookId)
	
	if err != nil {fmt.Println("Failed to set book url", err)}

	return err
}


/**
Takes in a filepath for an img file, uploads the file to the
file hosting service, then returns the URL to the file
*/
func UploadCoverAndGetURL(filepath string) (string, error) {
	filename := getFilename(filepath)
	savePath := `/home/leon/Documents/letterbookd/client/public/covers/` + filename
	err := os.Rename(filepath, savePath)
	if err != nil {
		return "", err
	}

	clientPath := `covers/` + filename
	return clientPath, nil
}
