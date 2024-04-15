package cli

import (
    "fmt"
    "os"

    "lightConventionalLog/internal/git"
)

// lcl full
// lcl full file.txt
// lcl from tag
// lcl from tag file.txt
func ParseArguments(args ...string) {
    if len(args) == 0 {
        fmt.Errorf("zero arguments")
        return
    }
    if args[0] == "full" {
        if len(args) > 1 {
            makeFullLogToFile(args[1])
        } else {
            makeFullLog()
        }
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
func makeFullLogToFile(fileName string) {
    file, err := os.Create(fileName)
    if err != nil {
        panic(err)
    }
    log := git.CreateFullChangeLog()
    _, err = file.WriteString(log)
    if err != nil {
        panic(err)
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
