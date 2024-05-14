package conventional

import (
    "bytes"
    "regexp"
)

var commitTypeRegExp = regexp.MustCompile("(feat|fix|docs|style|refactor|perf|test|chore|build|ci)(\\(.+\\))?:")

type ConventionalCommit struct {
    Type  string
    Scope string // if not, then general
    Title string
}

func ParseConventional(commit []byte) ConventionalCommit {
    type_ := commitTypeRegExp.Find(commit)
    scope := regexp.MustCompile("\\(.+\\)").Find(type_)
    conv := ConventionalCommit{}
    if len(scope) == 0 {
        conv.Type = string(type_)
        if len(conv.Type) > 1 {
            conv.Type = conv.Type[:len(conv.Type)-1]
        }

        conv.Scope = "general"
    } else {
        suff := string(scope) + ":"
        t, _ := bytes.CutSuffix(type_, []byte(suff))
        conv.Type = string(t)
        conv.Scope = string(scope[1 : len(scope)-1])
    }
    pref := regexp.MustCompile(".+(feat|fix|docs|style|refactor|perf|test|chore|build|ci)(\\(.+\\))?:")
    title, _ := bytes.CutPrefix(commit, pref.Find(commit))
    conv.Title = string(title)
    return conv
}
