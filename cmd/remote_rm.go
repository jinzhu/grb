package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(&cobra.Command{
		Use:   "remote_rm [name]",
		Short: "remove a remote repo",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			values := map[string]string{"name": args[0]}

			sh("{{git}} remote rm {{name}}", values)
		},
	})
}
