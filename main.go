package main

import (
	"fmt"
	"os/exec"

	"github.com/jinzhu/grb/cmd"
)

func main() {
	getCurrentBranch()
	cmd.Execute()
}

func getCurrentBranch() string {
	out, err := exec.Command("git", "branch").Output()
	fmt.Println("kkk")
	return "kk"
}
