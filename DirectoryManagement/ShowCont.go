package main

import (
    "fmt"
    "os"
)

func main() {
    showdir, err := os.ReadDir("users")
    if err != nil {
        fmt.Println(err)
    }

    for _, files := range showdir {
        fmt.Println(files.Name())
    }
}
