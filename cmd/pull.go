package cmd

import (
	"os"
	"os/exec"

	"github.com/fatih/color"
	"github.com/rajnandan1/okgit/models"
	"github.com/spf13/cobra"
)

var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull remote branch changes",
	Long:  "Pull remote branch changes. Similar to `git pull`",
	Run: func(cmd *cobra.Command, args []string) {

		gitPull := models.AllCommands["gitPull"]
		xmd := exec.Command(gitPull.Name, gitPull.Arguments...)
		xmd.Stdout = os.Stdout
		xmd.Stderr = os.Stderr
		if xmd.Run() == nil {
			color.Green(models.LINE)
		} else {
			color.Red(models.LINE)
		}
	},
}

func init() {
	rootCmd.AddCommand(pullCmd)
}
