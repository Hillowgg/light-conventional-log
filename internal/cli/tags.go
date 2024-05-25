package cli

import (
    "fmt"
    "slices"

    "github.com/spf13/cobra"
    "lightConventionalLog/internal/git"
    "lightConventionalLog/internal/repo"
)

var tagsCmd = &cobra.Command{
    Use:     "tags",
    Short:   "list all tags",
    Args:    cobra.NoArgs,
    PreRunE: checkGit,
    Run: func(cmd *cobra.Command, args []string) {

        cfg := repo.Tags{}
        cfg.Dir, _ = cmd.Flags().GetString("repo")
        tags := git.GetTags(cfg)
        rev, _ := cmd.Flags().GetBool("reverse")
        date, _ := cmd.Flags().GetBool("date")
        if rev {
            slices.Reverse(tags)
        }
        for _, t := range tags {
            if date {
                fmt.Println(t.Tag + " " + t.Date)
            } else {
                fmt.Println(t.Tag)
            }
        }
    },
}

func init() {
    tagsCmd.Flags().BoolP("date", "d", false, "show date")
    tagsCmd.Flags().StringP("repo", "r", "", "repository path")
    tagsCmd.Flags().BoolP("reverse", "v", false, "reverse order")
    rootCmd.AddCommand(tagsCmd)
}
