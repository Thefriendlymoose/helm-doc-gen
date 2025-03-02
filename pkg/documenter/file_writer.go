package documenter

import (
	"fmt"
	"os"
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

func GenerateFile(dataToWrite string, nameOfFile string, extenstion FileExtension) {
	file, err := os.Create(fmt.Sprintf("%s.%s", nameOfFile, extenstion))
	if err != nil {
		panic(err)
	}

	defer file.Close()

	_, err = file.WriteString(dataToWrite)
	if err != nil {
		panic(err)
	}
}
