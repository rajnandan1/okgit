package utils

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/rajnandan1/okgit/models"
)

const (
	directoryName = ".okgit"
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

func ReadBranchesFile(filePath string) (map[string][]models.Commit, error) {

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

func AddCommitToBranchFile(branchName string, cmt models.Commit, filePath string) error {
	branches, err := ReadBranchesFile(filePath)
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

	// Write the data to the file
	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func GetLastCommitForBranchFromFile(branchName string, filepath string) (*models.Commit, error) {
	branches, err := ReadBranchesFile(filepath)
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
func RunCommand(name string, args []string, stdin string) (string, error) {
	color.HiBlue("✨" + name + " " + strings.Join(args, " "))
	cmd := exec.Command(name, args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if stdin != "" {
		cmd.Stdin = strings.NewReader(stdin)
	}
	err := cmd.Run()
	if err != nil {
		return strings.TrimRight(stderr.String(), "\n"), err
	}
	return strings.TrimRight(stdout.String(), "\n"), nil
}

func LogFatal(err error) {
	if err != nil {
		red := color.New(color.FgRed).SprintFunc()
		log.Fatal(red(err.Error()))
	}
}

func LogOutput(output string) {
	if output != "" {
		green := color.New(color.FgGreen).SprintFunc()
		fmt.Println(green(output))
	}
}

func GetHomeDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return homeDir, nil
}

func CreateDirectory(directoryPath string) error {
	err := os.MkdirAll(directoryPath, 0755)
	if err != nil {
		return err
	}
	return nil
}

func GetRepoNameHash() string {
	remoteOriginUrlCmd := models.AllCommands["remoteOriginUrl"]
	remoteOriginUrl, err := RunCommand(remoteOriginUrlCmd.Name, remoteOriginUrlCmd.Arguments, "")
	if err != nil {
		remoteOriginUrl = "okgit"
	}
	//generate md5 hash of the remote origin url
	return fmt.Sprintf("%x", md5.Sum([]byte(remoteOriginUrl)))
}

func StartOkgit() (string, error) {
	homeDir, err := GetHomeDir()
	if err != nil {
		return "", err
	}

	directoryPath := filepath.Join(homeDir, "okgit")
	//filePath := filepath.Join(directoryPath, "branches.json")

	// Check if the directory already exists
	if _, err := ioutil.ReadDir(directoryPath); err != nil {
		// Directory does not exist, create it
		err := CreateDirectory(directoryPath)
		if err != nil {
			return "", err
		}
	}

	repoFolder := GetRepoNameHash()
	repoFolderPath := filepath.Join(directoryPath, repoFolder)

	// Check if the directory already exists
	if _, err := ioutil.ReadDir(repoFolderPath); err != nil {
		// Directory does not exist, create it
		err := CreateDirectory(repoFolderPath)
		if err != nil {
			return "", err
		}
	}

	fileName := "branches.json"
	filePath := filepath.Join(repoFolderPath, fileName)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		err = ioutil.WriteFile(filePath, []byte("{}"), 0644)
		if err != nil {
			return "", err
		}
	}

	return filePath, nil
}
