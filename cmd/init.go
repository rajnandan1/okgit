package cmd

import (
	"github.com/rajnandan1/okgit/utils"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Add .okgit/ to home folder",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := utils.StartOkgit()
		if err != nil {
			utils.LogFatal(err)
		}
		utils.LogOutput(".okgit/ added to home folder")

	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
