package main

import (
    "fmt"
    "log"
    "os"
    "os/user"
    "github.com/ewanvalentine/pine/repl"
)

func main() {
    user, err := user.Current()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Hello %s! Welcome to Pine.", user.Username)
    repl.Start(os.Stdin, os.Stdout)
}
