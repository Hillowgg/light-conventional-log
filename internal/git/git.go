package git

import (
    "bytes"
    "os/exec"
    "regexp"
)

var tagRegExp = regexp.MustCompile("tag: \\S+\\\\")
var dateRegExp = regexp.MustCompile("\\\\.+")

type Tag struct {
    Tag  string
    Date string
}

func getAllTags(dir string) []byte {
    cmd := exec.Command(
        "git",
        "log",
        "--tags",
        "--simplify-by-decoration",
        "--pretty=\"%D\\\\%ci\"",
    )
    cmd.Dir = dir
    out, err := cmd.Output()
    if err != nil {
        panic(err)
    }
    return out
}

type Tags interface {
    GetDir() string
}

func GetTags(cfg Tags) []Tag {
    tagsText := getAllTags(cfg.GetDir())
    tags := bytes.Split(tagsText, []byte("\n"))
    res := make([]Tag, 0, len(tags))
    for _, t := range tags {

        tag := tagRegExp.Find(t)
        date := dateRegExp.Find(t)
        if len(tag) == 0 {
            continue
        }
        res = append(res, Tag{string(tag[5 : len(tag)-2]), string(date[2 : len(date)-1])})
    }
    return res
}

func GetCommitsFromTag(tag From) []byte {
    cmd := exec.Command(
        "git",
        "log",
        "--oneline",
        tag.GetFrom()+"..",
    )

    out, err := cmd.Output()
    if err != nil {
        panic(err)
    }
    return out
}

func GetCommitsFromToTags(cfg FromToDir) []byte {
    cmd := exec.Command(
        "git",
        "log",
        "--oneline",
        cfg.GetFrom()+".."+cfg.GetTo(),
    )
    if d := cfg.GetDir(); d != "" {
        cmd.Dir = d
    }

    out, err := cmd.Output()
    if err != nil {
        panic(err)
    }
    return out
}
