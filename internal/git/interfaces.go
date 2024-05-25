package git

type From interface {
    GetFrom() string
}
type To interface {
    GetTo() string
}
type Dir interface {
    GetDir() string
}
type FromToDir interface {
    From
    To
    Dir
}
