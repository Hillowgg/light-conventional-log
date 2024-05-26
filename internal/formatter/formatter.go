package formatter

import (
    "bytes"

    "lightConventionalLog/internal/conventional"
    "lightConventionalLog/internal/git"
    "lightConventionalLog/internal/repo"
)

func CreateFullChangeLog(cfg repo.Full) map[git.Tag]string {
    tags := git.GetTags(repo.Tags{Dir: cfg.GetDir()})
    if len(tags) == 0 {
        panic("No tags found")
    }
    res := make(map[git.Tag]string, len(tags))
    from := repo.From{}
    from.From = tags[0].Tag
    from.IncludeScopes = cfg.GetIncludeScopes()
    res[tags[0]] = CreateChangeLogFrom(from)

    for i := len(tags) - 1; i > 0; i-- {
        c := repo.From{}
        c.From = tags[i].Tag
        c.To = tags[i-1].Tag
        c.Dir = cfg.GetDir()
        commit := ParseCommits(git.GetCommitsFromToTags(c))
        res[tags[i]] = commit
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
                ret += "- " + title + "\n"
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

type From interface {
    git.FromToDir
    GetIncludeScopes() bool
}

func CreateChangeLogFrom(cfg From) string {
    var commitsText []byte
    commitsText = git.GetCommitsFromToTags(cfg)
    if cfg.GetIncludeScopes() {
        return ParseCommits(commitsText)
    }
    return ParseCommitsWithoutScopes(commitsText)
}

func LastChangeLog(cfg repo.Update) (string, git.Tag) {
    tags := git.GetTags(repo.Tags{Dir: cfg.GetDir()})
    c := repo.From{}
    c.Dir = cfg.GetDir()
    c.IncludeScopes = cfg.IncludeScopes
    return CreateChangeLogFrom(c), tags[0]
}
