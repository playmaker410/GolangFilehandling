package main

import (
    "fmt"
    "os"
)

func main() {
    fileInfo, err := os.Stat("contact.txt")
    if err != nil {
        fmt.Println(err)
    }

  workdir, err := os.Getwd()
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println("File name:", fileInfo.Name())
    fmt.Println("Size:", fileInfo.Size())
    fmt.Println("Permissions:", fileInfo.Mode())
    fmt.Println("Last modified:", fileInfo.ModTime())
   fmt.Println("Current working directory:", workdir)
   
}
