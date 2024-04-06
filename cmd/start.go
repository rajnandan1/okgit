package cmd

import (
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
	"github.com/rajnandan1/okgit/models"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start working on new or existing branch",
	Long:  "Start working on new or existing branch. Similar to `git checkout -b` or `git checkout`. Example usage: okgit start branchName",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			color.Red("Please provide a branch name")
			return
		}
		branch := strings.TrimSpace(args[0])
		gitFetchBranch := models.AllCommands["gitFetchBranch"]
		gitFetchBranch.Arguments = append(gitFetchBranch.Arguments, branch)
		res, err := exec.Command(gitFetchBranch.Name, gitFetchBranch.Arguments...).Output()
		if err != nil {
			color.Red("Error fetching branches")
			return
		}
		if len(res) == 0 {
			// checkoutCmd.Run(cmd, args)
			createBranch := models.AllCommands["createBranch"]
			createBranch.Arguments = append(createBranch.Arguments, branch)
			xmd := exec.Command(createBranch.Name, createBranch.Arguments...)
			xmd.Stdout = os.Stdout
			xmd.Stderr = os.Stderr
			if xmd.Run() == nil {
				color.Green("✔ Created new branch %s", branch)
			} else {
				color.Red("⨯ Error creating  new branch %s", branch)
			}
		}

		checkoutCmd.Run(cmd, args)

		gitPull := models.AllCommands["gitPull"]
		gitPull.Arguments = append(gitPull.Arguments, branch)
		_, err1 := exec.Command(gitPull.Name, gitPull.Arguments...).Output()
		if err1 == nil {
			color.Green("✔ Pulled changes successfully")
		}

	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
