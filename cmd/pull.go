package cmd

import (
	"os"
	"os/exec"

	"github.com/fatih/color"
	"github.com/rajnandan1/okgit/models"
	"github.com/spf13/cobra"
)

var pullCmd = &cobra.Command{
	Use:   "pl",
	Short: "Pull remote branch changes. Similar to `git pull`",
	Long:  "Pull remote branch changes. Similar to `git pull`. Example usage: okgit pl",
	Run: func(cmd *cobra.Command, args []string) {

		//get current branch
		gitBracnh := models.AllCommands["gitBranch"]
		branch, err := exec.Command(gitBracnh.Name, gitBracnh.Arguments...).Output()
		if err != nil {
			branch = []byte("")
		} else {
			branch = branch[:len(branch)-1]
		}

		//expect the args[0] to be a branch name
		if len(args) > 0 {
			branch = []byte(args[0])
		}

		if len(branch) == 0 {
			color.Red("Error getting branch name")
			return
		}

		//checkout the branch
		gitCheckout := models.AllCommands["gitCheckout"]
		gitCheckout.Arguments = append(gitCheckout.Arguments, string(branch))
		xmd := exec.Command(gitCheckout.Name, gitCheckout.Arguments...)
		xmd.Stdout = os.Stdout
		xmd.Stderr = os.Stderr
		if xmd.Run() == nil {
			color.Green("✔ Checked out branch successfully")
		} else {
			color.Red("⨯ Error checking out branch")
			return
		}

		gitPull := models.AllCommands["gitPull"]
		gitPull.Arguments = append(gitPull.Arguments, string(branch))
		xmd = exec.Command(gitPull.Name, gitPull.Arguments...)
		xmd.Stdout = os.Stdout
		xmd.Stderr = os.Stderr
		if xmd.Run() == nil {
			color.Green("✔ Pulled changes successfully")
		} else {
			color.Red("⨯ Error pulling changes")
		}
	},
}

func init() {
	rootCmd.AddCommand(pullCmd)
}
