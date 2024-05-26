package cli

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
    "lightConventionalLog/internal/config"
    "lightConventionalLog/internal/formatter"
    "lightConventionalLog/internal/repo"
)

var fromCmd = &cobra.Command{
    Use:     "from <tag>",
    Short:   "create changelog from tag",
    Args:    cobra.ExactArgs(1),
    PreRunE: checkGit,
    Run: func(cmd *cobra.Command, args []string) {
        ns, _ := cmd.Flags().GetBool("no-scopes")
        if config.Config.IncludeScopes {
            ns = !ns
        }
        // date, _ := cmd.Flags().GetBool("date")
        // if config.Config.Dates {
        //     date = !date
        // }
        cfg := repo.From{}
        // cfg.TagDate = date
        cfg.From = args[0]
        cfg.IncludeScopes = !ns
        if config.Config.IncludeScopes {
            cfg.IncludeScopes = !cfg.IncludeScopes
        }
        cfg.To, _ = cmd.Flags().GetString("to")
        cfg.Dir, _ = cmd.Flags().GetString("repo")
        log := formatter.CreateChangeLogFrom(cfg)
        fileName, err := cmd.Flags().GetString("file")
        if err != nil {
            panic(err)
        }
        if fileName != "" {
            file, err := os.Create(fileName)
            _, err = file.WriteString(log)
            if err != nil {
                panic(err)
            }
            fmt.Println("Log created in " + fileName)
        } else {
            fmt.Println(log)
        }
    },
}

func init() {
    rootCmd.AddCommand(fromCmd)
    fromCmd.Flags().StringP("file", "f", "", "file to save log")
    // flag -n present

    // fromCmd.Flags().BoolP("date", "d", false, "show date")
    fromCmd.Flags().StringP("repo", "r", "", "repository path")
    fromCmd.Flags().BoolP("no-scopes", "n", false, "exclude scopes")
    fromCmd.Flags().StringP("to", "t", "", "create log from tag to tag")
}
