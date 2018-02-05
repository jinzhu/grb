package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(&cobra.Command{
		Use:   "mv [branch1] [branch2]",
		Short: "rename `branch1` to `branch2`, rename current branch to `branch1` if `branch2` not specified",
		Args:  cobra.RangeArgs(1, 2),
		Run: func(cmd *cobra.Command, args []string) {
			var branch1, branch2 string

			if len(args) == 1 {
				branch1, branch2 = getCurrentBranch(), args[0]
			} else {
				branch1, branch2 = args[0], args[1]
			}

			values := map[string]string{"branch1": branch1, "branch2": branch2}

			sh("{{git}} push {{origin}} {{branch1}}:refs/heads/{{branch2}}", values)
			sh("{{git}} fetch {{origin}}", values)
			sh("{{git}} branch --track {{branch2}} {{origin}}/{{branch2}}", values)
			sh("{{git}} checkout {{branch2}}", values)
			sh("{{git}} branch -d {{branch1}}", values)
			sh("{{git}} push {{origin}} :refs/heads/{{branch1}}", values)
		},
	})
}
