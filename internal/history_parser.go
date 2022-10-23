package internal

import (
	"bash-history-parser/internal/file_mng"
	"bufio"
	"errors"
	"fmt"
	"path"
	"strings"
)

func Run(filePath string, starting_dir string) error {

	if !file_mng.FileExists(filePath) {
		errorStr := fmt.Sprintf("File \"%s\" not exists\n", filePath)
		return errors.New(errorStr)
	}

	file, err := file_mng.OpenFile(filePath)

	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(file)
	file_content := parseHistory(scanner, starting_dir)

	file_mng.CloseFile(file)

	fmt.Println(file_content)

	return nil
}

func parseHistory(scanner *bufio.Scanner, starting_dir string) string {

	var file_content string
	var home_dir = starting_dir
	var current_dir = starting_dir

	for scanner.Scan() {
		line := scanner.Text()

		// change directory found
		if strings.HasPrefix(line, "cd") {
			current_dir = detectCurrentDir(line, home_dir, current_dir)
			file_content += fmt.Sprintf(line+"\t--> %s\n", current_dir)
		} else {
			file_content += line + "\n"
		}
	}

	return file_content

}

func detectCurrentDir(cdArg string, home_dir string, currDir string) string {

	cdArg = strings.TrimSpace(cdArg)
	// classic "cd"
	if len(cdArg) == 2 {
		return home_dir
	}

	// if longer than 2, means that there is arg, remove the first space
	if len(cdArg) > 2 {
		cdArg = cdArg[2:len(cdArg)]
		cdArg = strings.TrimSpace(cdArg)
	}

	// easy scenarios
	switch cdArg {
	case "~":
		return home_dir
	case "/":
		return "/"
	case ".":
		return currDir
	}

	// command starts with /, return the cdarg path
	if strings.HasPrefix(cdArg, "/") {
		return cdArg
	}

	// command starts with ~, create the full path
	if strings.HasPrefix(cdArg, "~") {
		cdArg = cdArg[1:len(cdArg)]
		return path.Clean(path.Join(home_dir, cdArg))
	}

	// otherwise join cd arg with current path
	tpath := path.Join(currDir, cdArg)
	return path.Clean(tpath)
}
