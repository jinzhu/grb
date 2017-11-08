package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(&cobra.Command{
		Use:   "remote_add [branch] [repo]",
		Short: "add a remote repo to branch",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			values := map[string]string{"branch": args[0], "repo": args[1]}

			sh("{{git}} remote add {{branch}} {{repo}}", values)
			sh("{{git}} fetch {{branch}}", values)
		},
	})
}
