package cmd

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/rajnandan1/okgit/utils"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Add .okgit/ to the .gitignore file",
	Run: func(cmd *cobra.Command, args []string) {
		gitignorePath := filepath.Join(".", ".gitignore")
		okgitPath := ".okgit/"

		// Read the contents of the .gitignore file
		data, err := ioutil.ReadFile(gitignorePath)
		if err != nil {
			utils.LogFatal(err)
		}

		// Append .okgit/ to the .gitignore file if it's not already present
		contents := string(data)
		if !strings.Contains(contents, okgitPath) {
			contents += okgitPath + "\n"

			// Write the updated contents back to the .gitignore file
			err = ioutil.WriteFile(gitignorePath, []byte(contents), 0644)
			if err != nil {
				utils.LogFatal(err)
			}

			utils.LogOutput(".okgit/ added to .gitignore file")
		} else {
			utils.LogOutput(".okgit/ already present in .gitignore file")
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
