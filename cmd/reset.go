package cmd

import (
	"github.com/rajnandan1/okgit/models"
	"github.com/rajnandan1/okgit/utils"
	"github.com/spf13/cobra"
)

var resetCmd = &cobra.Command{
	Use:   "rs",
	Short: "Reset changes in the working directory. Similar to `git reset`",
	Long:  "Reset changes in the working directory by providing the file names as arguments. Similar to `git reset`. Example usage: okgit rs ",
	Run: func(cmd *cobra.Command, args []string) {

		gitReset := models.AllCommands["gitReset"]
		cmdOut, cmdErr := utils.RunCommand(gitReset.Name, gitReset.Arguments, "")
		if cmdErr != nil {
			utils.LogFatal(cmdErr)
		}
		utils.LogOutput(cmdOut)

	},
}

func init() {
	rootCmd.AddCommand(resetCmd)
}
