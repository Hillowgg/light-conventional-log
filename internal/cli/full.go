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
        ns, _ := cmd.Flags().GetBool("no-scopes")

        logs := formatter.CreateFullChangeLog(!ns)
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
    fullCmd.Flags().BoolP("no-scopes", "n", false, "exclude scopes")
}
