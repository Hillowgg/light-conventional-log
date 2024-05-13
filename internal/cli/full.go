package cli

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
    "lightConventionalLog/internal/formatter"
)

var fullCmd = &cobra.Command{
    Use:   "full",
    Short: "create full log",
    Run: func(cmd *cobra.Command, args []string) {
        logs := formatter.CreateFullChangeLog()
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
}
