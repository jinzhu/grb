package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var (
	GitCmd = "git"
	origin = "origin"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "grb",
	Short: "A tool to simplify working with remote git branches",
	Long:  `A tool to simplify working with remote git branches`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func sh(cmd string, args ...interface{}) {
	cmd = fmt.Sprintf(cmd, args...)

	fmt.Println("command is ", cmd)
	// splitting head => g++ parts => rest of the command
	parts := strings.Fields(cmd)
	head := parts[0]
	parts = parts[1:len(parts)]

	out, err := exec.Command(head, parts...).Output()
	if err != nil {
		fmt.Printf("%s: ", err)
	}
	fmt.Printf("%s\n", out)
}

func getCurrentBranch() string {
	branch, _ := exec.Command(GitCmd, "rev-parse", "--abbrev-ref", "HEAD").Output()
	return strings.TrimSpace(string(branch))
}
