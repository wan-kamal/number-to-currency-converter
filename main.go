package main

import (
    "math"
    "strings"
    "log"
    "fmt"
    "strconv"
)

// Does it need to support negatives?
// what is the min max constraint
func main() {
    var input string
    fmt.Println("Please provide numeric argument:")

    _, err := fmt.Scan(&input)
    if err != nil {
        log.Fatalf("\x1b[;31;1mError while reading stdin: %v\n", err)
    }

    x, err := strconv.ParseFloat(input, 64)
    if err != nil {
        log.Fatalf("\x1b[;31;1mPlease provide valid numeric argument\n")
    }

    fmt.Println(NumberToCurrencyConverter(x))
}

func NumberToCurrencyConverter(input float64) string {
    base := map[int64]string {
        1   : "One",
        2   : "Two",
        3   : "Three",
        4   : "Four",
        5   : "Five",
        6   : "Six",
        7   : "Seven",
        8   : "Eight",
        9   : "Nine",
        10  : "Ten",
        11  : "Eleven",
        12  : "Twelve",
        13  : "Thirteen",
        14  : "Fourteen",
        15  : "Fifteen",
        16  : "Sixteen",
        17  : "Seventeen",
        18  : "Eighteen",
        19  : "Nineteen",
        20  : "Twenty",
        30  : "Thirty",
        40  : "Fourty",
        50  : "Fifty",
        60  : "Sixty",
        70  : "Seventy",
        80  : "Eighty",
        90  : "Ninety",
    }

    // https://en.wikipedia.org/wiki/Names_of_large_numbers
    ext := map[int]string {
        1   : "Thousand",
        2   : "Million",
        3   : "Billion",
    }

    // maximum input constraint
    m := 999_999_999_999.99

    // closure
    f := func(x float64) string{
        // 0 < x < m
        if x > m || x <= 0 {
            log.Fatalf("\x1b[;31;1mUnsupported amount: %v", x)
        }

        // responding variables
        dollars, cents := "", ""

        // manipulating variables
        i, d := 0, int64(x)
        c := int64(math.Round((x - float64(d)) * 100))

        // dollars handler
        for d > 0 {
            r := d % 1000

            db, dr := r / 100, r % 100
            s := base[db]

            // edge case
            if db > 0 {
                s += " Hundred"
            }

            if dr > 20 {
                fr := dr % 10
                dr -= fr
                // apparently correct format for amount > 20 is like this "thirty-two"
                s += " " + base[dr] + "-" + base[fr]
            } else {
                s += " " + base[dr]
            }

            // edge case
            if r == 0 {
                dollars = strings.TrimSpace(s) + " " + dollars
            } else {
            // ex: billion .. million .. thousand ..
                dollars = strings.TrimSpace(s) + " " + ext[i] + " " + strings.TrimSpace(dollars)
            }

            // amount / 1000 for next iteration
            d /= 1000
            i++
        }

        // cents handler
        if c > 0 {
            if c > 20 {
                fr := c % 10
                c -= fr
                cents = base[c] + "-" + base[fr]
            } else {
                cents = base[c]
            }
        }

        return formatter(dollars, cents)
    }

    return f(input)
}

// formatting helper
func formatter(d string, c string) string {
    s := ""
    
    if strings.TrimSpace(d) == "One" {
        // singular noun
        s = strings.TrimSpace(d) + " Dollar"
    } else if d == "" {
        s = ""
    } else {
        s = strings.TrimSpace(d) + " Dollars"
    }

    // append cents
    if c != "" {
        if c == "One" {
            // singular noun
            c = c + " Cent"
        } else {
            c = c + " Cents"
        }
        if s == "" {
            s += c
        } else {
            s += " and " + c
        }
    }

    return s
}
