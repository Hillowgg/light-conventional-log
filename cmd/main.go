package main

import (
    "os"

    "lightConventionalLog/internal/cli"
)

func main() {
    cli.ParseArguments(os.Args[1:]...)
}
