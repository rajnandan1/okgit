package cmd

import (
	"github.com/rajnandan1/okgit/models"
	"github.com/rajnandan1/okgit/utils"
	"github.com/spf13/cobra"
)

var checkoutCmd = &cobra.Command{
	Use:   "ch",
	Short: "Switch branches or restore working tree files. Similar to `git checkout`",
	Long:  "Switch branches or restore working tree files. Similar to `git checkout`. Example usage: okgit ch branchName / okgit ch file1 file2 / okgit ch . to ch all files",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			//checkout .
			args = append(args, ".")
		}

		gitCheckout := models.AllCommands["gitCheckout"]
		gitCheckout.Arguments = append(gitCheckout.Arguments, args...)
		cmdOut, cmdErr := utils.RunCommand(gitCheckout.Name, gitCheckout.Arguments, "")
		if cmdErr != nil {
			utils.LogFatal(cmdErr)
		}
		utils.LogOutput(cmdOut)

	},
}

func init() {
	rootCmd.AddCommand(checkoutCmd)
}
