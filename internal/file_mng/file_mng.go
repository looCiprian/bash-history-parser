package file_mng

import (
	"bufio"
	"os"
)

func FileExists(destination string) bool {
	if _, err := os.Stat(destination); err == nil {
		return true
	}
	return false
}

func OpenFile(filePath string) (*os.File, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func CloseFile(file *os.File) {
	file.Close()
}

func ScanFile(file *os.File) *bufio.Scanner {

	scanner := bufio.NewScanner(file)

	return scanner
}
