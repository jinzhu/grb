package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(&cobra.Command{
		Use:   "rm [branch]",
		Short: "delete branch `branch`",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			values := map[string]string{"branch": args[0]}

			sh("{{git}} push {{origin}} :refs/heads/{{branch}}", values)
			if getCurrentBranch() == args[0] {
				sh("{{git}} checkout master", values)
			}
			sh("{{git}} branch -d {{branch}}", values)
		},
	})
}
