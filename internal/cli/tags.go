package cli

import (
    "fmt"

    "github.com/spf13/cobra"
    "lightConventionalLog/internal/git"
)

var tagsCmd = &cobra.Command{
    Use:   "tags",
    Short: "list all tags",
    Run: func(cmd *cobra.Command, args []string) {
        tags := git.GetTags()
        for _, t := range tags {
            fmt.Println(t)
        }
    },
}

func init() {
    rootCmd.AddCommand(tagsCmd)
}