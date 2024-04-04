package cmd

import (
	"os"
	"os/exec"

	"github.com/fatih/color"
	"github.com/rajnandan1/okgit/models"
	"github.com/spf13/cobra"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset changes in the working directory",
	Long:  "Reset changes in the working directory by providing the file names as arguments. Similar to `git reset`.",
	Run: func(cmd *cobra.Command, args []string) {

		gitReset := models.AllCommands["gitReset"]
		xmd := exec.Command(gitReset.Name, gitReset.Arguments...)
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
	rootCmd.AddCommand(resetCmd)
}
