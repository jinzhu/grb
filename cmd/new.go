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
			values := map[string]string{"branch": args[0]}

			sh("{{git}} push {{origin}} {{current_branch}}:refs/heads/{{branch}}", values)
			sh("{{git}} fetch {{origin}}", values)
			sh("{{git}} branch --track {{branch}} {{origin}}:{{branch}}", values)
			sh("{{git}} checkout {{branch}}", values)
		},
	})
}
