package cmd

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
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

func sh(str string, args map[string]string) {
	var (
		result  = bytes.NewBufferString("")
		tmpl    = template.New("")
		funcMap = template.FuncMap{}
	)

	if _, ok := args["git"]; !ok {
		args["git"] = GitCmd
	}

	if _, ok := args["origin"]; !ok {
		args["origin"] = origin
	}

	if _, ok := args["current_branch"]; !ok {
		args["current_branch"] = getCurrentBranch()
	}

	for key, value := range args {
		funcValue := value
		funcMap[key] = func() string {
			return funcValue
		}
	}

	tmpl, err := tmpl.Funcs(funcMap).Parse(str)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	tmpl.Execute(result, args)
	cmd := result.String()

	color.Red(cmd)

	parts := strings.Fields(cmd)
	head := parts[0]
	parts = parts[1:len(parts)]

	output, err := exec.Command(head, parts...).CombinedOutput()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	if o := strings.TrimSpace(string(output)); o != "" {
		fmt.Println(o)
	}
}

func getCurrentBranch() string {
	branch, _ := exec.Command(GitCmd, "rev-parse", "--abbrev-ref", "HEAD").Output()
	return strings.TrimSpace(string(branch))
}
