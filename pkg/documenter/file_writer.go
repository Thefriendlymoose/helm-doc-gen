package documenter

import (
	"fmt"
	"os"
	"path/filepath"
)

type FileExtension int

const (
	MarkDown FileExtension = iota
	HTML
)

var FileExtensionName = map[FileExtension]string{
	MarkDown: "md",
	HTML:     "html",
}

func (fe FileExtension) String() string {
	return FileExtensionName[fe]
}

func GenerateFile(outputDir string, dataToWrite string, nameOfFile string, extenstion FileExtension) {

	absPath, err := filepath.Abs(outputDir)
	if err != nil {
		fmt.Println("Error resolving absolute path:", err)
		return
	}

	err = os.MkdirAll(absPath, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	outputFilePath := filepath.Join(absPath, fmt.Sprintf("%s.%s", nameOfFile, extenstion))

	file, err := os.Create(outputFilePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	defer file.Close()

	_, err = file.WriteString(dataToWrite)
	if err != nil {
		panic(err)
	}
}
