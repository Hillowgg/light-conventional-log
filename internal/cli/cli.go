package cli

import (
    "fmt"
    "os"

    "lightConventionalLog/internal/formatter"
    "lightConventionalLog/internal/git"
)

// lcl full
// lcl from tag
// lcl from tag file.txt
func ParseArguments(args ...string) {
    if len(args) == 0 {
        fmt.Errorf("zero arguments")
        return
    }
    if args[0] == "full" {
        makeFullLogToFiles()

    } else if args[0] == "from" {
        if len(args) == 2 {
            makeLogFromTag(args[1])
        } else if len(args) == 3 {
            makeLogFromTagToFile(args[1], args[2])
        }
    } else if args[0] == "tags" {
        printTags()
    } else {
        panic("wrong arguments")
    }

}
func makeFullLog() {
    log := formatter.CreateFullChangeLog()
    fmt.Println(log)
}
func makeFullLogToFiles() {
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
}

func makeLogFromTag(tag string) {
    log := formatter.CreateChangeLogFrom(tag)
    fmt.Println(log)
}

func makeLogFromTagToFile(tag string, fileName string) {
    file, err := os.Create(fileName)
    if err != nil {
        panic(err)
    }
    log := formatter.CreateChangeLogFrom(tag)
    _, err = file.WriteString(log)
    if err != nil {
        panic(err)
    }
    fmt.Println("Log created in " + fileName)
}
func printTags() {
    tags := git.GetTags()
    for _, t := range tags {
        fmt.Println(t)
    }
}
