package main

import (
    "fmt"
    "os"
    "os/user"
    "monkey/repl"
)


func main() {
    user, err := user.Current()
    if err != nil {
        panic(err)
    }

    fmt.Printf("Hello %s! Welcome to Monkey!\n", user.Username)
    fmt.Printf("Type In Commands Now:\n")
    repl.Start(os.Stdin, os.Stdout)
}
