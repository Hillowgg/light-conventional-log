package git

import (
    "bytes"
    "os/exec"
    "regexp"

    "lightConventionalLog/internal/conventional"
)

var tagRegExp = regexp.MustCompile("tag: \\S+\\\\")
var dateRegExp = regexp.MustCompile("\\\\.+")

type Tag struct {
    tag  string
    date string
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

func PrettyTags() []Tag {
    tagsText := getAllTags()
    tags := bytes.Split(tagsText, []byte("\n"))
    res := make([]Tag, 0, len(tags))
    for _, t := range tags {
        tag := tagRegExp.Find(t)[:len(t)-1]
        date := dateRegExp.Find(t)[1:]
        if len(tag) == 0 {
            continue
        }
        res = append(res, Tag{string(tag), string(date)})
    }
    return res
}

func GetCommitsFromTag(tag Tag) []byte {
    cmd := exec.Command(
        "git",
        "log",
        "--oneline",
        tag.tag+"..",
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
        fromTag.tag+".."+toTag.tag,
    )
    out, err := cmd.Output()
    if err != nil {
        panic(err)
    }
    return out
}

func CreateFullChangeLog() map[Tag]string {
    tags := PrettyTags()
    if len(tags) == 0 {
        panic("No tags found")
    }
    res := make(map[Tag]string, len(tags))
    res[tags[0]] = CreateChangeLogFrom(tags[0])
    for i := 1; i < len(tags); i++ {
        commit := parseCommits(GetCommitsFromToTags(tags[i], tags[i-1]))
        res[tags[i]] = string(commit)
    }
    return res
}

func parseCommits(commitsText []byte) string {
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

func CreateChangeLogFrom(tag Tag) string {
    commitsText := GetCommitsFromTag(tag)
    return parseCommits(commitsText)
}
