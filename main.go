package main

import (
    "fmt"
    "math/rand"
    "strconv"
    "time"
)

func randBool() bool {
    return rand.Intn(2) == 1
}

func randRange(min, max int) int {
    return rand.Intn(max-min) + min
}

const letters string = "QWERTYUIOPASDFGHJKLZXCVBNM"
const maxOrderLength int = 8
const useOfHashtag bool = true // Is use # in starts?

var ordersType map[string]string = map[string]string{
    "dev" : "DEV",
    "design" : "DESIGN",
};

func GenerateCode() string {
    result := ""
    
    if useOfHashtag {
        result += "#"
    }
    
    result += strconv.Itoa(randRange(0, 10));
    for i := 0; i < maxOrderLength; i++ { //
        isUseHereNumber := randBool()
        if isUseHereNumber {
            randomLetter := letters[randRange(0, len(letters))]
            result += string(randomLetter)
        } else {
            result += strconv.Itoa(randRange(0, 10))
        }
    }
    return result
}

func AddingOrderTypeInResult(typeOrder string) string {
    valueType, ok := ordersType[typeOrder];
    if !ok  {
        panic("Not founded: " + typeOrder);
    }
    
    return fmt.Sprintf("%s-%s", valueType, GenerateCode())
    
}

func main() {
    rand.Seed(time.Now().UnixNano()) // Seed the random number generator once
    for a := 0; a < 50; a++ {
        fmt.Printf("[%d] %s\n", a, AddingOrderTypeInResult("dev"))
    }
}
