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
    var decimalValue := { 1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1 };
    //var romanNumeral = [ "M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I" ];
    //var romanized = "";
    
    if n == 2 {
        return "II"
    }
    return "I"
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
