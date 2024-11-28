package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

// getRepoDiff runs the git diff command and returns the output as a string
func GetRepoDiff(repoPath string) (string, error) {
	cmd := exec.Command("git", "-C", repoPath, "diff")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), nil
}
func newfunc() {
	fmt.Println("ceshixin")
	//新增一个函数都不会处罚
}
