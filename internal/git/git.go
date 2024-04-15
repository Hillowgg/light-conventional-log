package git

import (
    "bytes"
    "os/exec"
    "regexp"
    "sync"
)

type Tag string

type GIT struct {
    tags   []Tag
    logs   map[Tag]map[int]int
    logsMu sync.RWMutex
}

func New() *GIT {
    return &GIT{}
}

var tagRegExp = regexp.MustCompile("tag: \\S+\\)")

func (g *GIT) getAllTags() []byte {
    cmd := exec.Command(
        "git",
        "log",
        "--tags",
        "--simplify-by-decoration",
        "--pretty=\"%d\"",
    )
    out, err := cmd.Output()
    if err != nil {
        panic(err)
    }
    return out
}

func (g *GIT) PrettyTags() []Tag {
    tagsText := g.getAllTags()
    tags := bytes.Split(tagsText, []byte("\n"))
    res := make([]Tag, 0, len(tags))
    for _, t := range tags {
        tag := tagRegExp.Find(t)
        if len(tag) == 0 {
            continue
        }
        res = append(res, Tag(tag[5:len(tag)-1]))

    }
    return res
}

func (g *GIT) GetCommitsFromTag(tag Tag) []byte {
    cmd := exec.Command(
        "git",
        "log",
        "--oneline",
        string(tag)+"..",
    )
    out, err := cmd.Output()
    if err != nil {
        panic(err)
    }
    return out
}

func (g *GIT) GetCommitFromToTags(fromTag Tag, toTag Tag) []byte {
    cmd := exec.Command(
        "git",
        "log",
        "--oneline",
        string(fromTag)+".."+string(toTag),
    )
    out, err := cmd.Output()
    if err != nil {
        panic(err)
    }
    return out
}

func (g *GIT) CreateFullChangeLog() {
    panic("Not implemented")
}
