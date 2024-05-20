package cli

import (
    "errors"
    "os/exec"

    "github.com/spf13/cobra"
)

func checkGit(cmd *cobra.Command, args []string) error {
    // check if git is installed
    isGit := exec.Command("git", "--version")
    if err := isGit.Run(); err != nil {
        return errors.New("git is not installed")
    }
    isRepo := exec.Command("git", "rev-parse", "--is-inside-work-tree")
    if out, err := isRepo.Output(); err != nil || string(out) == "false\n" {
        return errors.New("not a git repository")
    }
    return nil
}
