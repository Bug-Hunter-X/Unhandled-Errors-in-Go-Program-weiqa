package main

import (

    "fmt"

    "os"

)

func main() {

    file, err := os.Open("myFile.txt")

    if err != nil {

        fmt.Printf("Error opening file: %v\n", err)

        os.Exit(1)

    }

    defer file.Close()

    // Process the file

    fmt.Println("File opened successfully!")

}