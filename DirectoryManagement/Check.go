package main

import (
    "fmt"
    "os"
)

func main() {
    if _, err := os.Stat("users/Contact.txt"); os.IsNotExist(err) {
        fmt.Println("Directory not seen")
    } else {
        fmt.Println("Found")
    }

}
