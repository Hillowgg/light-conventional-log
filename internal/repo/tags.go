package repo

type Tags struct {
    Dir string
}

func (t Tags) GetDir() string {
    return t.Dir
}
