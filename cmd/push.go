package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(&cobra.Command{
		Use:   "push [branch]",
		Short: "push branch `branch`, default current_branch",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			branch := getCurrentBranch()
			if len(args) >= 1 {
				branch = args[0]
			}

			sh("%v push %v %v:refs/heads/%v", GitCmd, origin, branch, branch)
			sh("%v fetch %v", GitCmd, origin)
			sh("%v config branch.%v.remote %v", GitCmd, branch, origin)
			sh("%v config branch.%v.merge refs/heads/%v", GitCmd, branch, branch)
			sh("%v checkout %v", GitCmd, branch)
		},
	})
}
