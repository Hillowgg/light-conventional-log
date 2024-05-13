package cli

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
    "lightConventionalLog/internal/formatter"
)

var fromCmd = &cobra.Command{
    Use:   "from",
    Short: "create changelog from tag",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        ns, err := cmd.Flags().GetBool("no-scopes")
        if err != nil {
            panic(err)
        }
        log := formatter.CreateChangeLogFrom(args[0], ns)
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
    fromCmd.Flags().BoolP("no-scopes", "n", false, "exclude scopes")
}
