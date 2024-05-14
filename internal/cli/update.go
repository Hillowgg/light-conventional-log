package cli

import (
    "fmt"
    "io"
    "os"
    "os/exec"

    "github.com/spf13/cobra"
    "lightConventionalLog/internal/formatter"
)

var updateCmd = &cobra.Command{
    Use:   "update",
    Short: "update latest tag",
    Run: func(cmd *cobra.Command, args []string) {
        fileName, _ := cmd.Flags().GetString("file")
        ns, _ := cmd.Flags().GetBool("no-scopes")
        interactive, _ := cmd.Flags().GetBool("interactive")

        log, tag := formatter.LastChangeLog(!ns)
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
        }

        if fileName != "" {
            file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
            _, err = file.WriteString("\n# " + tag + "\n")
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
    updateCmd.Flags().StringP("file", "f", "", "file to save log")
    updateCmd.Flags().BoolP("no-scopes", "n", false, "exclude scopes")
    updateCmd.Flags().BoolP("interactive", "i", false, "interactive mode")

}
