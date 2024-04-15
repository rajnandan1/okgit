package cmd

import (
	"github.com/rajnandan1/okgit/models"
	"github.com/rajnandan1/okgit/utils"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "ad",
	Short: "Stage files for commit. Similar to `git add`",
	Long:  "Stage files for commit by providing the file names as arguments. Similar to `git add`. Example: okgit ad file1 file2 / okgit ad . to stage all files",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			//push .
			args = append(args, ".")
		}
		gitAdd := models.AllCommands["gitAdd"]
		gitAdd.Arguments = append(gitAdd.Arguments, args...)
		_, cmdErr := utils.RunCommand(gitAdd.Name, gitAdd.Arguments, "")
		if cmdErr != nil {
			utils.LogFatal(cmdErr)
		}
		utils.LogOutput("Staged files successfully")

	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
