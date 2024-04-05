package cmd

import (
	"os"
	"os/exec"

	"github.com/fatih/color"
	"github.com/rajnandan1/okgit/models"
	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push local branch changes to remote",
	Long:  "Push local branch changes to remote. Similar to `git push`",
	Run: func(cmd *cobra.Command, args []string) {

		gitPush := models.AllCommands["gitPush"]
		xmd := exec.Command(gitPush.Name, gitPush.Arguments...)
		xmd.Stdout = os.Stdout
		xmd.Stderr = os.Stderr
		if xmd.Run() == nil {
			color.Green("✔ Pushed changes successfully")
		} else {
			color.Red("⨯ Error pushing changes")
		}
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)
}
