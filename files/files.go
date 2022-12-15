package files

import (
	"bufio"
	"log"
	"os"
	"path"
	"strings"
)

var (
	filesFolder = "input-files"
)

func buildPath(filename string) string {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("Unable to get current filepath: ", err)
	}
	filepath := path.Join(wd, filesFolder, filename)

	return filepath
}

// ReadFile reads a file from the "input-files" folder and returns its content
// allowing to easily get input data
func ReadFile(filename string) string {
	filepath := buildPath(filename)
	content, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal("Failed to read file content: ", err)
	}

	return strings.TrimSpace(string(content))
}

// ReadFileByLine reads a file line by line from the "input-folder"
// and returns a scanner to iterate over each line
func ReadFileByLine(filename string) (*os.File, *bufio.Scanner) {
	filepath := buildPath(filename)
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal("Failed to open the file: ", err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	return file, scanner
}
