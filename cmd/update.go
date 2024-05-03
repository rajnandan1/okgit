package cmd

import (
	"strings"

	"github.com/rajnandan1/okgit/models"
	"github.com/rajnandan1/okgit/utils"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update okgit to the specific/latest version",
	Run: func(cmd *cobra.Command, args []string) {
		version := "latest"
		if len(args) > 0 {
			version = args[0]
		}
		updateOkgit := models.AllCommands["updateOkgit"]
		//get last index of the slice
		lastIndex := len(updateOkgit.Arguments) - 1
		if version != "latest" {
			updateOkgit.Arguments[lastIndex] = strings.Replace(updateOkgit.Arguments[lastIndex], "latest", version, 1)
		}
		cmdOut, cmdErr := utils.RunCommand(updateOkgit.Name, updateOkgit.Arguments, "")
		if cmdErr != nil {
			utils.LogFatal(cmdErr)
		}
		utils.LogOutput(cmdOut + "\nokgit updated successfully")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
