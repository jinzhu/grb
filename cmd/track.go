package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(&cobra.Command{
		Use:   "track [branch]",
		Short: "track branch `branch`",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			branch := getCurrentBranch()

			if len(args) >= 1 {
				branch = args[0]
			}
			values := map[string]string{"branch": branch}

			sh("{{git}} fetch {{origin}}", values)
			if getCurrentBranch() != branch {
				sh("{{git}} branch --track {{branch}} origin/{{branch}}", values)
			}
			sh("{{git}} checkout {{branch}}", values)
		},
	})
}
