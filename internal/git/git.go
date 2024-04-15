package git

import (
    "bytes"
    "os/exec"
    "regexp"

    "lightConventionalLog/internal/conventional"
)

var tagRegExp = regexp.MustCompile("tag: \\S+\\)")

type Tag string

func getAllTags() []byte {
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

func PrettyTags() []Tag {
    tagsText := getAllTags()
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

func GetCommitsFromTag(tag Tag) []byte {
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

func GetCommitsFromToTags(fromTag Tag, toTag Tag) []byte {
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

func CreateFullChangeLog() string {
    panic("Not implemented")
}

func CreateChangeLogFrom(tag Tag) string {
    commitsText := GetCommitsFromTag(tag)
    commits := bytes.Split(commitsText, []byte("\n"))
    res := make(map[string]map[string][]string, 0)
    for _, comm := range commits {
        conv := conventional.ParseConventional(comm)
        if _, ok := res[conv.Scope]; !ok {
            res[conv.Scope] = map[string][]string{}
        }
        res[conv.Scope][conv.Type] = append(res[conv.Scope][conv.Type], conv.Title)
    }
    ret := ""
    for scopeName, scope := range res {
        ret += "# " + scopeName + "\n"
        for typeName, t := range scope {
            ret += "- ## " + typeName + "\n"
            for _, title := range t {
                ret += "- - " + title + "\n"
            }
        }
    }
    return ret
}
