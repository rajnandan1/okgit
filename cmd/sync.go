package cmd

import (
	"errors"

	"github.com/rajnandan1/okgit/models"
	"github.com/rajnandan1/okgit/utils"
	"github.com/spf13/cobra"
)

var syncCmd = &cobra.Command{
	Use:   "sn",
	Short: "Sync local branch with remote from -> to",
	Long:  "Sync local branch with remote from -> to. if to is not given it will sync the current branch. Example usage: okgit sync fromBranchName toBranchName",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			utils.LogFatal(errors.New("Please provide the branch name to sync with"))
		}
		fromBranch := args[0]

		gitBranch := models.AllCommands["gitBranch"]

		toBranch, cmdErr := utils.RunCommand(gitBranch.Name, gitBranch.Arguments, "")
		if cmdErr != nil {
			utils.LogFatal(cmdErr)
		}
		if len(args) >= 2 {
			args = append(args, args[0])
			toBranch = args[1]
		}

		//check if unstaged or uncommit files are there in toBranch
		gitStatus := models.AllCommands["gitStatus"]
		cmdOut, cmdErr := utils.RunCommand(gitStatus.Name, gitStatus.Arguments, "")
		if cmdErr != nil {
			utils.LogFatal(cmdErr)
		}
		if len(cmdOut) > 0 {
			utils.LogFatal(errors.New("Please commit or stash the changes in the branch " + toBranch))
		}
		utils.LogOutput(cmdOut)
		//now checkout fromBranch
		gitCheckout := models.AllCommands["gitCheckout"]
		gitCheckout.Arguments = append(gitCheckout.Arguments, fromBranch)
		cmdOut, cmdErr = utils.RunCommand(gitCheckout.Name, gitCheckout.Arguments, "")
		if cmdErr != nil {
			utils.LogFatal(cmdErr)
		}
		utils.LogOutput(cmdOut)
		//now pull changes from remote
		gitPull := models.AllCommands["gitPull"]
		gitPull.Arguments = append(gitPull.Arguments, fromBranch)
		cmdOut, cmdErr = utils.RunCommand(gitPull.Name, gitPull.Arguments, "")
		if cmdErr != nil {
			utils.LogFatal(cmdErr)
		}
		utils.LogOutput(cmdOut)
		//now checkout toBranch
		gitCheckout.Arguments = []string{"checkout"}
		gitCheckout.Arguments = append(gitCheckout.Arguments, toBranch)
		cmdOut, cmdErr = utils.RunCommand(gitCheckout.Name, gitCheckout.Arguments, "")
		if cmdErr != nil {
			utils.LogFatal(cmdErr)
		}
		utils.LogOutput(cmdOut)

		//now pull changes from remote
		gitPull.Arguments = []string{"pull", "origin", toBranch}
		cmdOut, cmdErr = utils.RunCommand(gitPull.Name, gitPull.Arguments, "")
		if cmdErr != nil {
			utils.LogFatal(cmdErr)
		}
		utils.LogOutput(cmdOut)

		//now merge toBranch with fromBranch
		gitMerge := models.AllCommands["gitMerge"]
		gitMerge.Arguments = append(gitMerge.Arguments, fromBranch)
		cmdOut, cmdErr = utils.RunCommand(gitMerge.Name, gitMerge.Arguments, "")
		if cmdErr != nil {
			utils.LogFatal(cmdErr)
		}
		utils.LogOutput(cmdOut)
		utils.LogOutput("Synced branch " + fromBranch + " with " + toBranch)

	},
}

func init() {
	rootCmd.AddCommand(syncCmd)
}
