package cli

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
    "lightConventionalLog/internal/config"
)

var rootCmd = &cobra.Command{
    Use:   "lcl",
    Short: "Light Conventional Log",
    Long:  `Light and fast changelog generator based on conventional commits.`,
}

func Execute() {
    config.LoadConfig()
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}
