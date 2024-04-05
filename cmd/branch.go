package cmd

import (
	"os/exec"

	"github.com/fatih/color"
	"github.com/rajnandan1/okgit/models"
	"github.com/spf13/cobra"
)

var branchCmd = &cobra.Command{
	Use:   "bn",
	Short: "Get current branch name. Similar to `git branch`",
	Long:  "Get current branch name. Similar to `git branch`. Example usage: okgit bn",
	Run: func(cmd *cobra.Command, args []string) {

		gitBranch := models.AllCommands["gitBranch"]
		branch, err := exec.Command(gitBranch.Name, gitBranch.Arguments...).Output()
		if err != nil {
			color.Red("Is it a git repo? Error getting current branch")
			return
		}
		branch = branch[:len(branch)-1]
		color.Green("Current branch: %s", branch)

		// xmd.Stdout = os.Stdout
		// xmd.Stderr = os.Stderr
		// if xmd.Run() == nil {
		// 	color.Green("✔ Pulled changes successfully")
		// } else {
		// 	color.Red("⨯ Error pulling changes")
		// }
	},
}

func init() {
	rootCmd.AddCommand(branchCmd)
}
