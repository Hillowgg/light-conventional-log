package cli

import (
    "fmt"
    "io"
    "os"
    "os/exec"

    "github.com/spf13/cobra"
    "lightConventionalLog/internal/config"
    "lightConventionalLog/internal/formatter"
    "lightConventionalLog/internal/repo"
)

var updateCmd = &cobra.Command{
    Use:     "update",
    Short:   "update latest tag",
    PreRunE: checkGit,
    Run: func(cmd *cobra.Command, args []string) {
        fileName, _ := cmd.Flags().GetString("file")
        ns, _ := cmd.Flags().GetBool("no-scopes")
        if config.Config.IncludeScopes {
            ns = !ns
        }
        interactive, _ := cmd.Flags().GetBool("interactive")
        if config.Config.Interactive {
            interactive = !interactive
        }
        date, _ := cmd.Flags().GetBool("date")
        if config.Config.Dates {
            date = !date
        }
        cfg := repo.Update{}
        cfg.IncludeScopes = !ns
        cfg.Dir, _ = cmd.Flags().GetString("repo")
        log, tag := formatter.LastChangeLog(cfg)
        if interactive {
            file, _ := os.Create(".tmp-lcl.md")
            _, _ = file.WriteString(log)
            file.Close()
            vim := exec.Command("vim", ".tmp-lcl.md")
            vim.Stdin = os.Stdin
            vim.Stdout = os.Stdout
            vim.Run()
            file, _ = os.Open(".tmp-lcl.md")
            log_, _ := io.ReadAll(file)
            log = string(log_)
            err := os.Remove(".tmp-lcl.md")
            if err != nil {
                panic(err)
            }
        }

        if fileName != "" {
            file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
            if date {
                _, err = file.WriteString("\n# " + tag.Tag + " " + tag.Date + "\n")
            } else {

                _, err = file.WriteString("\n# " + tag.Tag + "\n")
            }
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
    rootCmd.AddCommand(updateCmd)

    updateCmd.Flags().BoolP("date", "d", false, "show date")
    updateCmd.Flags().StringP("repo", "r", "", "repository path")
    updateCmd.Flags().StringP("file", "f", "", "file to save log")
    updateCmd.Flags().BoolP("no-scopes", "n", false, "exclude scopes")
    updateCmd.Flags().BoolP("interactive", "i", false, "interactive mode")

}
