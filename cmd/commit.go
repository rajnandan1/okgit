package cmd

import (
	"errors"
	"strings"

	"github.com/fatih/color"
	"github.com/rajnandan1/okgit/models"
	"github.com/rajnandan1/okgit/utils"
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

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "cm",
	Short: "Create a conventional commit. Similar to `git commit`",
	Long:  `Create a conventional commit message by providing the necessary details. Example usage: okgit cm`,
	Run: func(cmd *cobra.Command, args []string) {

		// Get the current branch
		fgCyan := color.New(color.FgCyan)
		fgGray := color.New(color.FgBlack)

		myCommit := models.Commit{}

		branchFile, err := utils.StartOkgit()
		if err != nil {
			utils.LogFatal(err)
		}

		//Get the last commit if available

		gitBranch := models.AllCommands["gitBranch"]

		branch, cmdErr := utils.RunCommand(gitBranch.Name, gitBranch.Arguments, "")
		if cmdErr != nil {
			utils.LogFatal(cmdErr)
		}
		if storedCommit, err := utils.GetLastCommitForBranchFromFile(branch, branchFile); err == nil {
			myCommit = *storedCommit
		}
		if myCommit.Type == "" {
			myCommit.Type = getCommitTypeFromBranchName(branch)
		}

		// Ask for the commit type
		fgCyan.Print("[Required] Type (feat, fix, docs, build, chore, ci, docs, style, refactor, perf, test, others): ")
		if myCommit.Type != "" {
			fgGray.Print(myCommit.Type + " ")
		}

		commitTypeInput := utils.ReadInput(false)
		if commitTypeInput == "" && myCommit.Type == "" {
			utils.LogFatal(errors.New("Commit type is required."))
		}
		if commitTypeInput != "" {
			myCommit.Type = commitTypeInput
		}

		if !contains(commitTypes, myCommit.Type) {
			utils.LogFatal(errors.New("Invalid commit type. Please provide a valid commit type."))
		}

		// Ask for the commit scope
		fgCyan.Print("[Optional] Scope: ")
		if myCommit.Scope != "" {
			fgGray.Print(myCommit.Scope + " (enter . to remove me)")
		}
		commitScope := utils.ReadInput(false)
		if commitScope == "." {
			myCommit.Scope = ""
		} else if commitScope != "" {
			myCommit.Scope = commitScope
		}

		// Ask for the commit summary
		fgCyan.Println("[Required] Summary: ")
		if myCommit.Summary != "" {
			fgGray.Println(myCommit.Summary)
		}

		commitSummaryInput := utils.ReadInput(false)

		if commitSummaryInput != "" {
			myCommit.Summary = commitSummaryInput
		}

		if myCommit.Summary == "" {
			utils.LogFatal(errors.New("Commit summary is required."))
		}

		// Ask for the commit message
		fgCyan.Println("[Optional] Details: ")
		if myCommit.Details != "" {
			fgGray.Println(myCommit.Details + " (enter . to remove me)")
		}
		commitDetailsInput := utils.ReadInput(false)
		if commitDetailsInput == "." {
			myCommit.Details = ""
		} else if commitDetailsInput != "" {
			myCommit.Details = commitDetailsInput
		}

		// Ask if it is a breaking change
		fgCyan.Print("[Required] Breaking change? (y/n): ")
		if myCommit.BreakingChange {
			fgGray.Print("y ")
		} else {
			fgGray.Print("n ")
		}

		breakingChangeInput := utils.ReadInput(false)
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
				fgGray.Println(myCommit.BreakingMessage + " (enter . to remove me)")
			}
			breakingMessageInput := utils.ReadInput(false)

			if breakingMessageInput == "." {
				myCommit.BreakingMessage = ""
			} else if breakingMessageInput != "" {
				myCommit.BreakingMessage = breakingMessageInput
			}
		}
		// Ask for the footer
		fgCyan.Println("[Optional] Footer: ")
		if myCommit.Footer != "" {
			fgGray.Println(myCommit.Footer)
		}
		footerInput := utils.ReadInput(false)
		if footerInput == "." {
			myCommit.Footer = ""
		} else if footerInput != "" {
			myCommit.Footer = footerInput
		}

		commit := generateCommit(myCommit)

		// utils.CreateDirectoryAndFileIfNotExist()
		errWr := utils.AddCommitToBranchFile(string(branch), myCommit, branchFile)
		if errWr != nil {
			utils.LogFatal(errWr)
		}
		gitCommit := models.AllCommands["gitCommit"]
		cmdOutCommit, cmdErr := utils.RunCommand(gitCommit.Name, gitCommit.Arguments, commit)
		if cmdErr != nil {
			utils.LogFatal(cmdErr)
		}
		utils.LogOutput(cmdOutCommit)

	},
}

func generateCommit(cmt models.Commit) string {

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
		if strings.Contains(branchName, commitType) {
			return commitType
		}
	}

	return ""
}
