package cli

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
    "lightConventionalLog/internal/config"
    "lightConventionalLog/internal/formatter"
    "lightConventionalLog/internal/repo"
)

var fullCmd = &cobra.Command{
    Use:     "full",
    Short:   "create full log",
    PreRunE: checkGit,
    Run: func(cmd *cobra.Command, args []string) {
        ns, _ := cmd.Flags().GetBool("no-scopes")
        cfg := repo.Full{}
        cfg.IncludeScopes = !ns
        cfg.Dir, _ = cmd.Flags().GetString("repo")
        logs := formatter.CreateFullChangeLog(cfg)
        one, _ := cmd.Flags().GetString("one-file")
        if one != "" {
            file, err := os.Create(one)
            if err != nil {
                panic(err)
            }
            for tag, log := range logs {
                _, err = file.WriteString("# " + tag.Tag + "\n")
                _, err = file.WriteString(log)
                if err != nil {
                    panic(err)
                }
            }
            fmt.Println("Log created in " + one)
            return
        }

        for tag, log := range logs {
            file, err := os.Create(tag.Tag + ".md")
            if err != nil {
                panic(err)
            }
            _, err = file.WriteString(log)
            if err != nil {
                panic(err)
            }
            fmt.Println("Log created in " + tag.Tag + ".md")
        }
    },
}

func init() {
    rootCmd.AddCommand(fullCmd)
    fullCmd.Flags().StringP("one-file", "o", "", "create one file with all logs")
    fromCmd.Flags().StringP("repo", "r", "", "repository path")
    fullCmd.Flags().BoolP("no-scopes", "n", config.Config.IncludeScopes, "exclude scopes")
}
