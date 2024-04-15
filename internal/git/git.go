package git

import (
    "bytes"
    "os/exec"
    "regexp"
    "sync"
)

var tagRegExp = regexp.MustCompile("tag: \\S+\\)")
var commitTypeRegExp = regexp.MustCompile("(feat|fix|docs|style|refactor|perf|test|chore|build|ci)(\\(.+\\))?:")

type Tag string

type GIT struct {
    tags   []Tag
    logs   map[Tag]map[int]int
    logsMu sync.RWMutex
}

func New() *GIT {
    return &GIT{}
}

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

func (g *GIT) GetCommitsFromToTags(fromTag Tag, toTag Tag) []byte {
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

type ConventionalCommit struct {
    Type  string
    Scope string // if not, then general
    Title string
}

func parseConventional(commit []byte) ConventionalCommit {
    type_ := commitTypeRegExp.Find(commit)
    scope := regexp.MustCompile("\\(.+\\)").Find(type_)
    conv := ConventionalCommit{}
    if len(scope) == 0 {
        conv.Type = string(type_)
        conv.Scope = "general"
    } else {
        t, _ := bytes.CutSuffix(type_, scope)
        conv.Type = string(t)
        conv.Scope = string(scope[1 : len(scope)-1])
    }
    pref := regexp.MustCompile(".+(feat|fix|docs|style|refactor|perf|test|chore|build|ci)(\\(.+\\))?:")
    title, _ := bytes.CutPrefix(commit, pref.Find(commit))
    conv.Title = string(title)
    return conv
}

func (g *GIT) CreateChangeLogFrom(tag Tag) string {
    commitsText := g.GetCommitsFromTag(tag)
    commits := bytes.Split(commitsText, []byte("\n"))
    res := make(map[string]map[string][]string, 0)
    for _, comm := range commits {
        conv := parseConventional(comm)
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
