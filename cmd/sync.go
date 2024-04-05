package cmd

import (
	"os"
	"os/exec"

	"github.com/fatih/color"
	"github.com/rajnandan1/okgit/models"
	"github.com/spf13/cobra"
)

var syncCmd = &cobra.Command{
	Use:   "sn",
	Short: "Sync local branch with remote from -> to",
	Long:  "Sync local branch with remote from -> to. if to is not given it will sync the current branch. Example usage: okgit sync fromBranchName toBranchName",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			color.Red("Please provide the branch name to sync with")
			return
		}
		fromBranch := args[0]

		gitBranch := models.AllCommands["gitBranch"]
		data, err := exec.Command(gitBranch.Name, gitBranch.Arguments...).Output()
		if err != nil {
			color.Red("Is it a git repo?")
			return
		}
		toBranch := string(data)
		if len(args) >= 2 {
			args = append(args, args[0])
			toBranch = args[1]
		}

		//check if unstaged or uncommit files are there in toBranch
		gitStatus := models.AllCommands["gitStatus"]
		data, err = exec.Command(gitStatus.Name, gitStatus.Arguments...).Output()
		if err != nil {
			color.Red("Is it a git repo?")
			return
		}
		if len(data) > 0 {
			color.Red("Please commit or stash the changes in the branch %s", toBranch)
			return
		}

		//now checkout fromBranch
		color.Yellow("Checking out the branch %s", fromBranch)
		gitCheckout := models.AllCommands["gitCheckout"]
		gitCheckout.Arguments = append(gitCheckout.Arguments, fromBranch)
		xmd := exec.Command(gitCheckout.Name, gitCheckout.Arguments...)
		xmd.Stdout = os.Stdout
		xmd.Stderr = os.Stderr
		if xmd.Run() != nil {
			color.Red("Error in checking out the branch %s", fromBranch)
			return
		}

		//now pull changes from remote
		color.Yellow("Pulling the branch %s", fromBranch)
		gitPull := models.AllCommands["gitPull"]
		xmd = exec.Command(gitPull.Name, gitPull.Arguments...)
		xmd.Stdout = os.Stdout
		xmd.Stderr = os.Stderr
		if xmd.Run() != nil {
			color.Red("Error in pulling the branch %s", fromBranch)
			return
		}

		//now checkout toBranch
		color.Yellow("Checking out the branch %s", toBranch)
		gitCheckout.Arguments = []string{"checkout"}
		gitCheckout.Arguments = append(gitCheckout.Arguments, toBranch)
		xmd = exec.Command(gitCheckout.Name, gitCheckout.Arguments...)
		xmd.Stdout = os.Stdout
		xmd.Stderr = os.Stderr
		if xmd.Run() != nil {
			color.Red("Error in checking out the branch %s", toBranch)
			return
		}

		//now pull changes from remote
		color.Yellow("Pulling the branch %s", toBranch)
		xmd = exec.Command(gitPull.Name, gitPull.Arguments...)
		xmd.Stdout = os.Stdout
		xmd.Stderr = os.Stderr
		if xmd.Run() != nil {
			color.Red("Error in pulling the branch %s", toBranch)
			return
		}

		//now merge toBranch with fromBranch
		color.Yellow("Merging the branch %s with %s", fromBranch, toBranch)
		gitMerge := models.AllCommands["gitMerge"]
		gitMerge.Arguments = append(gitMerge.Arguments, fromBranch)
		xmd = exec.Command(gitMerge.Name, gitMerge.Arguments...)
		xmd.Stdout = os.Stdout
		xmd.Stderr = os.Stderr
		if xmd.Run() != nil {
			color.Red("Error in merging the branch %s with %s", fromBranch, toBranch)
			return
		} else {
			color.Green("✔ Synced the branch %s with %s successfully", fromBranch, toBranch)
		}

	},
}

func init() {
	rootCmd.AddCommand(syncCmd)
}
