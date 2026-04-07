package main

import (
    "fmt"
    "os"
)

func main() {
    err := os.Remove("filedelete.go")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("Succesfully deleted")
}
