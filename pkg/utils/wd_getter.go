package utils

import (
	"fmt"
	"os"
)

func GetWorkingDir() (string, error) {
	wd, err := os.Getwd()

	if err != nil {
		fmt.Println("Error: ", err)
		return "", fmt.Errorf("getWorkingDir: %s", err.Error())
	}

	return wd, nil
}
