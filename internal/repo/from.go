package repo

type From struct {
    Common
    From string
    To   string
    Dir  string
}

func (c From) GetFrom() string {
    return c.From
}

func (f From) GetTo() string {
    return f.To
}
