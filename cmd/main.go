package main

import (
    "fmt"

    git2 "lightConventionalLog/internal/git"
)

func main() {
    git := git2.New()
    tags := git.GetCommitFromToTags("0.109.0", "0.110.0")
    fmt.Println(string(tags))
}
