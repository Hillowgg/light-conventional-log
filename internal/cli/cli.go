package cli

import (
    "fmt"
    "os"

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
            makeLogFromTag(git.Tag(args[1]))
        } else if len(args) == 3 {
            makeLogFromTagToFile(git.Tag(args[1]), args[2])
        }
    } else if args[0] == "tags" {
        printTags()
    } else {
        panic("wrong arguments")
    }

}
func makeFullLog() {
    log := git.CreateFullChangeLog()
    fmt.Println(log)
}
func makeFullLogToFiles() {
    logs := git.CreateFullChangeLog()
    for tag, log := range logs {
        file, err := os.Create(string(tag) + ".md")
        if err != nil {
            panic(err)
        }
        _, err = file.WriteString(log)
        if err != nil {
            panic(err)
        }
    }
}

func makeLogFromTag(tag git.Tag) {
    log := git.CreateChangeLogFrom(tag)
    fmt.Println(log)
}

func makeLogFromTagToFile(tag git.Tag, fileName string) {
    file, err := os.Create(fileName)
    if err != nil {
        panic(err)
    }
    log := git.CreateChangeLogFrom(tag)
    _, err = file.WriteString(log)
    if err != nil {
        panic(err)
    }
}
func printTags() {
    tags := git.PrettyTags()
    for _, t := range tags {
        fmt.Println(t)
    }
}
