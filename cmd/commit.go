package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var commitTypes = []string{
	"feat",
	"fix",
	"docs",
	"build",
	"chore",
	"ci",
	"docs",
	"style",
	"refactor",
	"perf",
	"test",
	"others",
}

type ShellCommands struct {
	Name      string
	Arguments []string
}

var shellCommands = map[string]ShellCommands{
	"currentBranch": {
		Name:      "git",
		Arguments: []string{"branch", "--show-current"},
	},
	"gitCommit": {
		Name:      "git",
		Arguments: []string{"commit", "-F", "-"},
	},
}

type Commit struct {
	Type            string
	Scope           string
	Summary         string
	Details         string
	BreakingChange  bool
	BreakingMessage string
	Footer          string
}

const (
	directoryName = ".conventionalcommits"
	fileName      = "branches.json"
)

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Create a conventional commit",
	Long:  `Create a conventional commit message by providing the necessary details.`,
	Run: func(cmd *cobra.Command, args []string) {

		// Get the current branch
		fgCyan := color.New(color.FgCyan)
		fgGray := color.New(color.FgBlack)

		myCommit := Commit{}

		//Get the last commit if available

		currentBranch := shellCommands["currentBranch"]

		branch, err := exec.Command(currentBranch.Name, currentBranch.Arguments...).Output()
		if err != nil {
			color.Red("Is it a git repo? Error getting current branch")
			return
		}
		branch = branch[:len(branch)-1]
		if storedCommit, err := getLastCommitForBranchFromFile(string(branch)); err == nil {
			myCommit = *storedCommit
		}
		if myCommit.Type == "" {
			myCommit.Type = getCommitTypeFromBranchName(string(branch))
		}

		// Ask for the commit type
		fgCyan.Print("[Required] Type (feat, fix, docs, build, chore, ci, docs, style, refactor, perf, test, others): ")
		if myCommit.Type != "" {
			fgGray.Print(myCommit.Type + " ")
		}

		commitTypeInput := readInput(false)
		if commitTypeInput == "" && myCommit.Type == "" {
			color.Red("Commit type is required.")
			return
		}
		if commitTypeInput != "" {
			myCommit.Type = commitTypeInput
		}

		if !contains(commitTypes, myCommit.Type) {
			color.Red("Invalid commit type. Please provide a valid commit type.")
			return
		}

		// Ask for the commit scope
		fgCyan.Print("[Optional] Scope: ")
		if myCommit.Scope != "" {
			fgGray.Print(myCommit.Scope + " ")
		}
		commitScope := readInput(false)
		if commitScope != "" {
			myCommit.Scope = commitScope
		}

		// Ask for the commit summary
		fgCyan.Println("[Required] Summary: ")
		if myCommit.Summary != "" {
			fgGray.Println(myCommit.Summary)
		}

		commitSummaryInput := readInput(false)
		if commitSummaryInput == "" && myCommit.Summary == "" {
			color.Red("Commit summary is required.")
			return
		}
		if commitSummaryInput != "" {
			myCommit.Summary = commitSummaryInput
		}

		// Ask for the commit message
		fgCyan.Println("[Optional] Details: ")
		if myCommit.Details != "" {
			fgGray.Println(myCommit.Details)
		}
		commitDetailsInput := readInput(false)
		if commitDetailsInput != "" {
			myCommit.Details = commitDetailsInput
		}

		// Ask if it is a breaking change
		fgCyan.Print("[Required] Breaking change? (y/n): ")
		if myCommit.BreakingChange {
			fgGray.Print("y ")
		} else {
			fgGray.Print("n ")
		}

		breakingChangeInput := readInput(false)
		if breakingChangeInput != "" {
			if strings.ToLower(breakingChangeInput) == "y" {
				myCommit.BreakingChange = true
			} else {
				myCommit.BreakingChange = false
			}
		}

		if myCommit.BreakingChange {
			fgCyan.Println("[Optional] What is breaking?")
			if myCommit.BreakingMessage != "" {
				fgGray.Println(myCommit.BreakingMessage)
			}
			breakingMessageInput := readInput(false)

			if breakingMessageInput != "" {
				myCommit.BreakingMessage = breakingMessageInput
			}
		}
		// Ask for the footer
		fgCyan.Println("[Optional] Footer: ")
		if myCommit.Footer != "" {
			fgGray.Println(myCommit.Footer)
		}
		footerInput := readInput(false)
		if footerInput != "" {
			myCommit.Footer = footerInput
		}

		commit := generateCommit(myCommit)

		fmt.Println("Generated commit message:")
		fmt.Println(commit)
		createDirectoryAndFileIfNotExist()
		err = addCommitToBranchFile(string(branch), myCommit)
		if err != nil {
			color.Red("Error adding commit to branch file:", err)
			return
		}

		gitCommit := shellCommands["gitCommit"]
		xmd := exec.Command(gitCommit.Name, gitCommit.Arguments...)
		xmd.Stdout = os.Stdout
		xmd.Stderr = os.Stderr
		xmd.Stdin = strings.NewReader(commit)
		xmderr := xmd.Run()
		if xmderr != nil {
			color.Red("Error committing the changes:", xmderr)
			return
		}

	},
}

func generateCommit(cmt Commit) string {

	commitType := cmt.Type
	commitScope := cmt.Scope
	commitSummary := cmt.Summary
	commitDetails := cmt.Details
	isBreakingChange := cmt.BreakingChange
	breakingMessage := cmt.BreakingMessage
	footer := cmt.Footer

	commit := commitType
	if commitScope != "" {
		commit += "(" + commitScope + ")"
		if isBreakingChange {
			commit += "!"
		}
	}
	commit += ": " + commitSummary

	if commitDetails != "" {
		commit += "\n\nCHANGE: " + commitDetails
	}

	if isBreakingChange && breakingMessage != "" {
		commit += "\n\nBREAKING CHANGE: " + breakingMessage
	}

	if footer != "" {
		commit += "\n\n" + footer
	}

	return commit
}

func init() {
	rootCmd.AddCommand(commitCmd)
}

// contains checks if a string is in a slice of strings
func contains(slice []string, item string) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}

func getCommitTypeFromBranchName(branchName string) string {
	// Get the commit type from the branch name
	// The branch name should be in the format `type/description`
	if branchName == "" {
		return ""
	}

	for _, commitType := range commitTypes {
		if strings.HasPrefix(branchName, commitType) {
			return commitType
		}
	}

	return ""
}

func createDirectoryAndFileIfNotExist() error {

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

func readBranchesFile() (map[string][]Commit, error) {

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
	var branches map[string][]Commit

	err = json.Unmarshal(data, &branches)
	if err != nil {
		return nil, err
	}

	// Create the map of commits

	return branches, nil
}

func addCommitToBranchFile(branchName string, cmt Commit) error {
	branches, err := readBranchesFile()
	if err != nil {
		return err
	}

	commits, ok := branches[branchName]
	if !ok {
		commits = []Commit{}
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

func getLastCommitForBranchFromFile(branchName string) (*Commit, error) {
	branches, err := readBranchesFile()
	if err != nil {
		return nil, err
	}

	commits, ok := branches[branchName]
	if !ok {
		return nil, fmt.Errorf("branch " + branchName + " does not exist")
	}

	return &commits[len(commits)-1], nil

}

func readInput(multi bool) string {
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
