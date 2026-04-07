package main

import (
    "fmt"
    "io"
    "os"
)

func main() {
    srcFile, err := os.Open("contact.txt")
    if err != nil {
        fmt.Println(err)
    }
    defer srcFile.Close()

    destFile, err := os.Create("./log.txt")
    if err != nil {
        fmt.Println(err)
    }
    defer destFile.Close()

    _, err = io.Copy(destFile, srcFile)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("Copy done!")
}
