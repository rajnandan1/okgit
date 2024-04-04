package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/rajnandan1/okgit/models"
)

const (
	directoryName = ".conventionalcommits"
	fileName      = "branches.json"
)

func CreateDirectoryAndFileIfNotExist() error {

	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}

	directoryPath := filepath.Join(currentDir, directoryName)
	filePath := filepath.Join(directoryPath, fileName)

	// Check if the directory already exists
	if _, err := os.Stat(directoryPath); os.IsNotExist(err) {
		// Directory does not exist, create it
		err := os.MkdirAll(directoryPath, 0755)
		if err != nil {
			return err
		}
	}

	// Check if the file already exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {

		err = ioutil.WriteFile(filePath, []byte("{}"), 0644)
		if err != nil {
			return err
		}
	}

	return nil
}

func ReadBranchesFile() (map[string][]models.Commit, error) {

	currentDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	filePath := filepath.Join(currentDir, directoryName, fileName)

	// Check if the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("file " + directoryName + "/" + fileName + " does not exist")
	}

	// Read the file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON data into branches slice
	var branches map[string][]models.Commit

	err = json.Unmarshal(data, &branches)
	if err != nil {
		return nil, err
	}

	// Create the map of commits

	return branches, nil
}

func AddCommitToBranchFile(branchName string, cmt models.Commit) error {
	branches, err := ReadBranchesFile()
	if err != nil {
		return err
	}

	commits, ok := branches[branchName]
	if !ok {
		commits = []models.Commit{}
	}

	commits = append(commits, cmt)
	branches[branchName] = commits

	// Marshal the branches map into JSON
	data, err := json.Marshal(branches)
	if err != nil {
		return err
	}

	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}

	filePath := filepath.Join(currentDir, directoryName, fileName)

	// Write the data to the file
	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func GetLastCommitForBranchFromFile(branchName string) (*models.Commit, error) {
	branches, err := ReadBranchesFile()
	if err != nil {
		return nil, err
	}

	commits, ok := branches[branchName]
	if !ok {
		return nil, fmt.Errorf("branch " + branchName + " does not exist")
	}

	return &commits[len(commits)-1], nil

}

func ReadInput(multi bool) string {
	reader := bufio.NewReader(os.Stdin)

	if !multi {

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		return input
	}

	var lines []string
	for {
		line, _ := reader.ReadString('\n')

		// Remove the newline character from the end
		line = strings.TrimSuffix(line, "\n")

		// Check if the line is empty (only Enter was pressed)
		if line == "" {
			break
		}

		lines = append(lines, line)
	}

	input := strings.Join(lines, "\n")
	input = strings.TrimSpace(input)
	return input
}
