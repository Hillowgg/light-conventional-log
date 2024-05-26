package repo

type Full struct {
    Common
    TagDate bool
}

func (f Full) GetTagDate() bool {
    return f.TagDate
}
