package main

import (
   "fmt"
   "log"
)

func main() {
    captcha := 0
    var prev byte

    // If the input file has at least two values, put the last value into "prev" placeholder
    // to handle the wrap around scenario first
    if len(input) > 1 {
        prev = input[len(input) - 1]
    } else {
        log.Fatal("Bad input file.")
    }

    // First pass will compare last value (above) with first value, and continue from there
    for i := 0; i < len(input); i++ {
        if prev == input[i] {
            if prev > 47 && prev < 58 {
                captcha += int(prev - 48)
            } else {
                log.Fatal("Invalid number character in input.")
            }
        }
        prev = input[i]
    }
    fmt.Println("Captcha: ", captcha)
}
