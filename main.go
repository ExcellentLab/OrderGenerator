package main

import (
    "fmt"
    "math/rand"
    "strconv"
)

func randRange(min, max int) int {
    return rand.Intn(max-min) + min
}

const letters string = "QWERTYUIOPASDFGHJKLZXCVBNM";
const maxOrderLenght int = 8;

func GenerateCode() string {
    result := "#"
    result += strconv.Itoa(randRange(0, 9));
    for i := 0; i < maxOrderLenght; i++ {
        randomLetter := letters[randRange(0, len(letters))]
        result += string(randomLetter)
    }
    return result
}

func main() {
    /*fmt.Println(letters)
    fmt.Println(maxOrderLenght)*/
    for a := 0; a < 2; a++ {
        fmt.Printf("Generated: %s\n", GenerateCode())
    }
}
