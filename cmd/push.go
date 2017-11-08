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
			values := map[string]string{"branch": getCurrentBranch()}
			if len(args) >= 1 {
				values["branch"] = args[0]
			}

			sh("{{git}} push {{origin}} {{origin}}:refs/heads/{{branch}}", values)
			sh("{{git}} fetch {{origin}}", values)
			sh("{{git}} config branch.{{branch}}.remote {{branch}}", values)
			sh("{{git}} config branch.{{branch}}.merge refs/heads/{{branch}}", values)
			sh("{{git}} checkout {{branch}}", values)
		},
	})
}
