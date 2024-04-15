package cmd

import (
	"github.com/rajnandan1/okgit/models"
	"github.com/rajnandan1/okgit/utils"
	"github.com/spf13/cobra"
)

var branchCmd = &cobra.Command{
	Use:   "bn",
	Short: "Get current branch name. Similar to `git branch`",
	Long:  "Get current branch name. Similar to `git branch`. Example usage: okgit bn",
	Run: func(cmd *cobra.Command, args []string) {

		gitBranch := models.AllCommands["gitBranch"]

		cmdOut, cmdErr := utils.RunCommand(gitBranch.Name, gitBranch.Arguments, "")
		if cmdErr != nil {
			utils.LogFatal(cmdErr)
		}

		output := cmdOut

		lastCommitData := models.AllCommands["lastCommitData"]
		cmdOut, cmdErr = utils.RunCommand(lastCommitData.Name, lastCommitData.Arguments, "")
		if cmdErr == nil {
			output += "\n" + cmdOut
		}

		lastCommitAuthor := models.AllCommands["lastCommitAuthor"]
		cmdOut, cmdErr = utils.RunCommand(lastCommitAuthor.Name, lastCommitAuthor.Arguments, "")
		if cmdErr == nil {
			output += "\n" + cmdOut
		}

		utils.LogOutput(output)

	},
}

func init() {
	rootCmd.AddCommand(branchCmd)
}
