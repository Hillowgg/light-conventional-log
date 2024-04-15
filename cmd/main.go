package main

import (
    "fmt"

    git2 "lightConventionalLog/internal/git"
)

func main() {
    git := git2.New()
    tags := git.CreateChangeLogFrom("18.0.0-next.3")
    fmt.Println(tags)
    // for _, t := range tags {
    //     fmt.Println(t)
    // }
    // fmt.Println(string(tags))
}
