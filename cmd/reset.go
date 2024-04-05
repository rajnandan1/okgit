package cmd

import (
	"os"
	"os/exec"

	"github.com/fatih/color"
	"github.com/rajnandan1/okgit/models"
	"github.com/spf13/cobra"
)

var resetCmd = &cobra.Command{
	Use:   "rs",
	Short: "Reset changes in the working directory. Similar to `git reset`",
	Long:  "Reset changes in the working directory by providing the file names as arguments. Similar to `git reset`. Example usage: okgit rs ",
	Run: func(cmd *cobra.Command, args []string) {

		gitReset := models.AllCommands["gitReset"]
		xmd := exec.Command(gitReset.Name, gitReset.Arguments...)
		xmd.Stdout = os.Stdout
		xmd.Stderr = os.Stderr
		if xmd.Run() == nil {
			color.Green("✔ Reset changes successfully")
		} else {
			color.Red("⨯ Error resetting changes")
		}
	},
}

func init() {
	rootCmd.AddCommand(resetCmd)
}
