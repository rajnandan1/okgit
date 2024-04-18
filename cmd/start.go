package cmd

import (
	"errors"
	"strings"

	"github.com/rajnandan1/okgit/models"
	"github.com/rajnandan1/okgit/utils"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start working on new or existing branch",
	Long:  "Start working on new or existing branch. Similar to `git checkout -b` or `git checkout`. Example usage: okgit start branchName",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			utils.LogFatal(errors.New("Please provide the branch name to start working on"))
		}

		branch := strings.TrimSpace(args[0])
		gitFetchBranch := models.AllCommands["gitFetchBranch"]
		gitFetchBranch.Arguments = append(gitFetchBranch.Arguments, branch)
		res, err := utils.RunCommand(gitFetchBranch.Name, gitFetchBranch.Arguments, "")
		if err != nil {
			utils.LogFatal(err)
		}

		splitRes := strings.Split(res, "\n")
		branchPresent := false
		for _, line := range splitRes {
			trimmed := strings.TrimSpace(line)
			trimmed = strings.TrimPrefix(trimmed, "* ")
			if trimmed == branch {
				branchPresent = true
				break
			}
		}

		if !branchPresent {
			createBranch := models.AllCommands["createBranch"]
			createBranch.Arguments = append(createBranch.Arguments, branch)
			cmdOut, cmdErr := utils.RunCommand(createBranch.Name, createBranch.Arguments, "")
			if cmdErr != nil {
				utils.LogFatal(cmdErr)
			}
			utils.LogOutput(cmdOut)
		}

		checkoutCmd.Run(cmd, args)

		gitPull := models.AllCommands["gitPull"]
		gitPull.Arguments = append(gitPull.Arguments, branch)

		utils.RunCommand(gitPull.Name, gitPull.Arguments, "")

	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
