package main

import (
    "fmt"
    "math/rand"
    "strconv"
    "time"
)

const (
    letters         = "QWERTYUIOPASDFGHJKLZXCVBNM"
    maxOrderLength  = 8
    useOfHashtag    = true // Использовать # в начале?
    
    TYPE_DEV    = "dev"
    TYPE_DESIGN = "design"
)

var ordersType = map[string]string{
    TYPE_DEV:    "DEV",
    TYPE_DESIGN: "DESIGN",
}

// randBool возвращает случайное булево значение
func randBool() bool {
    return rand.Intn(2) == 1
}

// randRange возвращает случайное число в заданном диапазоне [min, max)
func randRange(min, max int) int {
    return rand.Intn(max-min) + min
}

// GenerateCode генерирует случайный код
func GenerateCode() string {
    result := ""

    if useOfHashtag {
        result += "#"
    }

    result += strconv.Itoa(randRange(0, 10))
    for i := 0; i < maxOrderLength; i++ {
        if randBool() {
            randomLetter := letters[randRange(0, len(letters))]
            result += string(randomLetter)
        } else {
            result += strconv.Itoa(randRange(0, 10))
        }
    }
    return result
}

// SelectOrderType возвращает тип заказа по его строковому представлению
func SelectOrderType(typeOrder string) string {
    valueType, ok := ordersType[typeOrder]
    if !ok {
        panic("Не найден тип: " + typeOrder)
    }
    return valueType
}

func main() {
    rand.Seed(time.Now().UnixNano()) // Инициализация генератора случайных чисел

    for a := 0; a < 50; a++ {
        randomTypes := []string{TYPE_DEV, TYPE_DESIGN}
        randomType := randomTypes[randRange(0, len(randomTypes))] // Исправлено на randomTypes

        fmt.Printf("[%d] %s\n", a, fmt.Sprintf("%s-%s", SelectOrderType(randomType), GenerateCode()))
    }
}