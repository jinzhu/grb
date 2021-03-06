package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(&cobra.Command{
		Use:   "pull [branch]",
		Short: "pull branch `branch`,default current_branch",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			branch := getCurrentBranch()

			if len(args) >= 1 {
				branch = args[0]
			}
			values := map[string]string{"branch": branch}

			sh("{{git}} fetch {{origin}}", values)

			if getCurrentBranch() != branch {
				sh("{{git}} checkout {{branch}}", values)
			}

			if hasLocalBranch(branch) {
				sh("{{git}} config branch.{{branch}}.remote {{origin}}", values)
				sh("{{git}} config branch.{{branch}}.merge refs/heads/{{branch}}", values)
			} else {
				sh("{{git}} branch --track {{branch}} {{origin}}/{{branch}}", values)
			}

			sh("{{git}} rebase {{origin}}/{{branch}}", values)
		},
	})
}
