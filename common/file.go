package common

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ServeFile(fileDirectory string) []byte {
	file, err := os.Open(fileDirectory)
	defer file.Close()

	if err != nil {
		return []byte{}
	}

	byteContent, err := ioutil.ReadAll(file)
	if err != nil {
		return []byte{}
	}

	return []byte(byteContent)
}

func CreateFile(fileDirectory string, fileContents []byte) bool {
	if len(ServeFile(fileDirectory)) != 0 {
		os.Remove(fileDirectory)
	}

	err := os.WriteFile(fileDirectory, fileContents, 0755)

	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

func RemoveFile(fileDirectory string) {
	os.Remove(fileDirectory)
}
