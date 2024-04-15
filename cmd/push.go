package cmd

import (
	"github.com/rajnandan1/okgit/models"
	"github.com/rajnandan1/okgit/utils"
	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:   "ps",
	Short: "Push local branch changes to remote. Similar to `git push`",
	Long:  "Push local branch changes to remote. Similar to `git push`. Example usage: okgit ps",
	Run: func(cmd *cobra.Command, args []string) {

		gitPush := models.AllCommands["gitPush"]
		cmdOut, cmdErr := utils.RunCommand(gitPush.Name, gitPush.Arguments, "")
		if cmdErr != nil {
			utils.LogFatal(cmdErr)
		}
		utils.LogOutput(cmdOut)

	},
}

func init() {
	rootCmd.AddCommand(pushCmd)
}
