package files

import (
	"log"
	"os"
	"path"
)

var filesFolder = "input-files"

// ReadFile reads a file from the "input-files" folder and returns its content
// allowing to easily get input data
func ReadFile(filename string) []byte {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("Unable to get current filepath: ", err)
	}
	filepath := path.Join(wd, filesFolder, filename)
	content, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal("Failed to read file content: ", err)
	}

	return content
}

// ReadFileByLine reads a file line by line from the "input-folder"
func ReadFileByLine(filename string) {}
