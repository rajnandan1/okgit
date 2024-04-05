package cmd

import (
	"os"
	"os/exec"

	"github.com/fatih/color"
	"github.com/rajnandan1/okgit/models"
	"github.com/spf13/cobra"
)

var checkoutCmd = &cobra.Command{
	Use:   "ch",
	Short: "Switch branches or restore working tree files. Similar to `git checkout`",
	Long:  "Switch branches or restore working tree files. Similar to `git checkout`. Example usage: okgit checkout branchName / okgit checkout file1 file2 / okgit checkout . to checkout all files",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			//checkout .
			args = append(args, ".")
		}

		gitCheckout := models.AllCommands["gitCheckout"]
		gitCheckout.Arguments = append(gitCheckout.Arguments, args...)
		xmd := exec.Command(gitCheckout.Name, gitCheckout.Arguments...)
		xmd.Stdout = os.Stdout
		xmd.Stderr = os.Stderr
		if xmd.Run() == nil {
			color.Green("✔ Switched branches or restored files successfully")
		} else {
			color.Red("⨯ Error switching branches or restoring files")
		}

	},
}

func init() {
	rootCmd.AddCommand(checkoutCmd)
}
