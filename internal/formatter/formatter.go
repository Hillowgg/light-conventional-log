package formatter

import (
    "bytes"

    "lightConventionalLog/internal/conventional"
    "lightConventionalLog/internal/git"
)

func CreateFullChangeLog() map[git.Tag]string {
    tags := git.GetTags()
    if len(tags) == 0 {
        panic("No tags found")
    }
    res := make(map[git.Tag]string, len(tags))
    res[tags[0]] = CreateChangeLogFrom(tags[0].Tag, true)
    for i := 1; i < len(tags); i++ {
        commit := ParseCommits(git.GetCommitsFromToTags(tags[i], tags[i-1]))
        res[tags[i]] = string(commit)
    }
    return res
}

func ParseCommits(commitsText []byte) string {
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

func ParseCommitsWithoutScopes(commitsText []byte) string {
    commits := bytes.Split(commitsText, []byte("\n"))
    res := make(map[string][]string, 0)
    for _, comm := range commits {
        conv := conventional.ParseConventional(comm)
        if _, ok := res[conv.Type]; !ok {
            res[conv.Type] = []string{}
        }
        res[conv.Type] = append(res[conv.Type], conv.Title)
    }
    ret := ""
    for typeName, t := range res {
        ret += "# " + typeName + "\n"
        for _, title := range t {
            ret += "- " + title + "\n"
        }
    }
    return ret
}

func CreateChangeLogFrom(tag string, scopes bool) string {
    commitsText := git.GetCommitsFromTag(tag)
    if scopes {
        return ParseCommits(commitsText)
    }
    return ParseCommitsWithoutScopes(commitsText)
}
