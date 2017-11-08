package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(&cobra.Command{
		Use:   "prune",
		Short: "cleaning up old remote git branches",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			sh("{{git}} remote prune {{origin}}", map[string]string{})
		},
	})
}
