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

		gitPull := models.AllCommands["gitPull"]
		xmd := exec.Command(gitPull.Name, gitPull.Arguments...)
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
