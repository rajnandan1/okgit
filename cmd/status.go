package cmd

import (
	"github.com/rajnandan1/okgit/models"
	"github.com/rajnandan1/okgit/utils"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "st",
	Short: "Check the status of the repository. Similar to `git status`",
	Long:  "Check the status of the repository. Similar to `git status`. Example usage: okgit st",
	Run: func(cmd *cobra.Command, args []string) {

		gitStatus := models.AllCommands["gitStatus"]
		cmdOut, cmdErr := utils.RunCommand(gitStatus.Name, gitStatus.Arguments, "")
		if cmdErr != nil {
			utils.LogFatal(cmdErr)
		}
		utils.LogOutput(cmdOut)

	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
