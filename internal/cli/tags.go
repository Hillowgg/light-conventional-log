package cli

import (
    "fmt"

    "github.com/spf13/cobra"
    "lightConventionalLog/internal/git"
)

var tagsCmd = &cobra.Command{
    Use:     "tags",
    Short:   "list all tags",
    PreRunE: checkGit,
    Run: func(cmd *cobra.Command, args []string) {
        tags := git.GetTags()
        for _, t := range tags {
            fmt.Println(t.Tag)
        }
    },
}

func init() {
    rootCmd.AddCommand(tagsCmd)
}
