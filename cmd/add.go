package cmd

import (
	"os"
	"os/exec"

	"github.com/fatih/color"
	"github.com/rajnandan1/okgit/models"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Stage files for commit",
	Long:  "Stage files for commit by providing the file names as arguments. Similar to `git add`.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			//push .
			args = append(args, ".")
		}
		gitAdd := models.AllCommands["gitAdd"]
		gitAdd.Arguments = append(gitAdd.Arguments, args...)
		xmd := exec.Command(gitAdd.Name, gitAdd.Arguments...)
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
	rootCmd.AddCommand(addCmd)
}