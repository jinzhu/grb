package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(&cobra.Command{
		Use:   "new [branch]",
		Short: "create new branch",
		Long:  `create a new branch with branch name`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			branch := args[0]
			sh("%v push %v %v:refs/heads/%v", GitCmd, origin, getCurrentBranch(), branch)
			sh("%v fetch %v", GitCmd, origin)
			sh("%v branch --track %v %v:refs/heads/%v", GitCmd, branch, origin, branch)
			sh("%v checkout %v", GitCmd, branch)
		},
	})
}
