package main

import (
    "fmt"
    "os"
)
//to create a single directory youse "os.Mkdir" but for nested dir use the one below
func main() {
    err := os.MkdirAll("users/Cotact.txt", 0755)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("Directory Created Successfully")
}
