package main

import (
    "fmt"
    "net/http"
    "strconv"
)

func hello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello world!")
}

func to_roman(n int) string {
    decimalValue := [13]int{ 1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1 }
    fmt.Println("decimalValue:", decimalValue)
    romanNumeral := [13]string{ "M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I" }
    fmt.Println("romanNumeral:", romanNumeral)
    var romanized = ""
    
    for index := 0; index < len(decimalValue); index++ {
        for decimalValue[index] <= n {
            romanized += romanNumeral[index];
            n -= decimalValue[index];
        }
    }

    return romanized;

}

type romanGenerator int

func (n romanGenerator) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    number := r.URL.Query().Get("number")
    if len(number) == 0 {
        fmt.Fprintf(w, "Please pass the number as parameter in the URL")
    }
    i, err := strconv.Atoi(number)

    if err == nil {
        fmt.Fprintf(w, "Here's your number: %s\n", to_roman(i))
    }
}

func main() {
    h := http.NewServeMux()

    h.Handle("/roman/", romanGenerator(1))
    h.HandleFunc("/", hello)

    err := http.ListenAndServe(":8000", h)
    panic(err)
}
