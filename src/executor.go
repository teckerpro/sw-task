package main

import (
    "fmt"
    "math/rand"
    "strings"
    "time"
)

func main() {
    rand.Seed(time.Now().Unix())
    var output strings.Builder
    charSet := "abcdedfghijklmnopqrstABCDEFGHIJKLMNOP!@#$^&*()_+-"
    length := 20
    for i := 0; i < length; i++ {
        random := rand.Intn(len(charSet))
        randomChar := charSet[random]
        output.WriteString(string(randomChar))
    }
    fmt.Println(output.String())
}