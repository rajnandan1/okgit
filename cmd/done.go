package cmd

import (
	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Do add commit and push at one go",
	Long:  "Do add commit and push at one go. Example usage: okgit done",
	Run: func(cmd *cobra.Command, args []string) {

		addCmd.Run(cmd, args)
		commitCmd.Run(cmd, args)
		pushCmd.Run(cmd, args)

	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
