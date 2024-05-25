package repo

type Common struct {
    Dir           string
    IncludeScopes bool
}

func (c Common) GetDir() string {
    return c.Dir
}

func (c Common) GetIncludeScopes() bool {
    return c.IncludeScopes
}
