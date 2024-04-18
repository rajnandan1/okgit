package cmd

import (
	"errors"

	"github.com/rajnandan1/okgit/models"
	"github.com/rajnandan1/okgit/utils"
	"github.com/spf13/cobra"
)

var pullCmd = &cobra.Command{
	Use:   "pl",
	Short: "Pull remote branch changes. Similar to `git pull`",
	Long:  "Pull remote branch changes. Similar to `git pull`. Example usage: okgit pl",
	Run: func(cmd *cobra.Command, args []string) {

		//get current branch
		gitBracnh := models.AllCommands["gitBranch"]
		branch, err := utils.RunCommand(gitBracnh.Name, gitBracnh.Arguments, "")
		if err != nil {
			branch = ""
		}

		//expect the args[0] to be a branch name
		if len(args) > 0 {
			branch = args[0]
		}

		if len(branch) == 0 {
			utils.LogFatal(errors.New("Please provide the branch name to pull changes"))
		}

		//checkout the branch
		gitCheckout := models.AllCommands["gitCheckout"]
		gitCheckout.Arguments = append(gitCheckout.Arguments, string(branch))
		cmdOut, cmdErr := utils.RunCommand(gitCheckout.Name, gitCheckout.Arguments, "")
		if cmdErr != nil {
			utils.LogFatal(cmdErr)
		}
		utils.LogOutput(cmdOut)

		gitPull := models.AllCommands["gitPull"]
		gitPull.Arguments = append(gitPull.Arguments, string(branch))
		utils.RunCommand(gitPull.Name, gitPull.Arguments, "")

	},
}

func init() {
	rootCmd.AddCommand(pullCmd)
}
