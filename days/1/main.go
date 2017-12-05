package main

import (
   "fmt"
   "log"
)

func main() {
    captcha := 0
    var half int

    // If the input file does not have at least two values, or does not has an even number of values stop now 
    // otherwise save the half-length of the input array
    if !(len(input) > 1 && (len(input) % 2) == 0) {
        log.Fatal("Bad input file.")
    } else {
        half = len(input)/2
    }

    // Otherwise, we'll compare each value to the value one-half the length of the ciricular array away
    for i := 0; i < len(input); i++ {
        if input[i] == input[(i + half) % len(input)] {
            if input[i] > 47 && input[i] < 58 {
                captcha += int(input[i] - 48)
            } else {
                log.Fatal("Invalid number character in input.")
            }
        }
    }
    fmt.Println("Captcha: ", captcha)
}
