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

func getAllTags() []byte {
    cmd := exec.Command(
        "git",
        "log",
        "--tags",
        "--simplify-by-decoration",
        "--pretty=\"%D\\\\%ci\"",
    )
    out, err := cmd.Output()
    if err != nil {
        panic(err)
    }
    return out
}

func GetTags() []Tag {
    tagsText := getAllTags()
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

func GetCommitsFromTag(tag string) []byte {
    cmd := exec.Command(
        "git",
        "log",
        "--oneline",
        tag+"..",
    )
    out, err := cmd.Output()
    if err != nil {
        panic(err)
    }
    return out
}

func GetCommitsFromToTags(fromTag Tag, toTag Tag) []byte {
    cmd := exec.Command(
        "git",
        "log",
        "--oneline",
        fromTag.Tag+".."+toTag.Tag,
    )
    out, err := cmd.Output()
    if err != nil {
        panic(err)
    }
    return out
}
