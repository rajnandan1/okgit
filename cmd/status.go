package cmd

import (
	"os"
	"os/exec"

	"github.com/fatih/color"
	"github.com/rajnandan1/okgit/models"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "st",
	Short: "Check the status of the repository. Similar to `git status`",
	Long:  "Check the status of the repository. Similar to `git status`. Example usage: okgit st",
	Run: func(cmd *cobra.Command, args []string) {

		gitStatus := models.AllCommands["gitStatus"]
		xmd := exec.Command(gitStatus.Name, gitStatus.Arguments...)
		xmd.Stdout = os.Stdout
		xmd.Stderr = os.Stderr
		if xmd.Run() == nil {
			color.Green("✔ status shown successfully")
		} else {
			color.Red("⨯ Error showing status")
		}
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
