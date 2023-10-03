package main

import (
    "testing"
)

func TestNumberToCurrencyConverter(t *testing.T) {
    tests := map[float64]string {
        0.01                : "One Cent",
        0.99                : "Ninety-Nine Cents",
        1.00                : "One Dollar",
        1.15                : "One Dollar and Fifteen Cents",
        20.00               : "Twenty Dollars",
        20.99               : "Twenty Dollars and Ninety-Nine Cents",
        77.00               : "Seventy-Seven Dollars",
        77.01               : "Seventy-Seven Dollars and One Cent",
        77.35               : "Seventy-Seven Dollars and Thirty-Five Cents",
        99.99               : "Ninety-Nine Dollars and Ninety-Nine Cents",
        100.00              : "One Hundred Dollars",
        100.01              : "One Hundred Dollars and One Cent",
        999.99              : "Nine Hundred Ninety-Nine Dollars and Ninety-Nine Cents",
        1_000.00            : "One Thousand Dollars",
        1_000.01            : "One Thousand Dollars and One Cent",
        1_001.00            : "One Thousand One Dollars",
        1_010.00            : "One Thousand Ten Dollars",
        1_099.99            : "One Thousand Ninety-Nine Dollars and Ninety-Nine Cents",
        1_100.00            : "One Thousand One Hundred Dollars",
        9_999.99            : "Nine Thousand Nine Hundred Ninety-Nine Dollars and Ninety-Nine Cents",
        10_000.00           : "Ten Thousand Dollars",
        10_001.00           : "Ten Thousand One Dollars",
        100_000.00          : "One Hundred Thousand Dollars",
        1_000_000.00        : "One Million Dollars",
        1_000_000.01        : "One Million Dollars and One Cent",
        1_000_001.00        : "One Million One Dollars",
        1_000_010.00        : "One Million Ten Dollars",
        1_000_100.00        : "One Million One Hundred Dollars",
        1_001_000.00        : "One Million One Thousand Dollars",
        1_010_000.00        : "One Million Ten Thousand Dollars",
        1_100_000.00        : "One Million One Hundred Thousand Dollars",
        10_000_000.00       : "Ten Million Dollars",
        100_000_000.00      : "One Hundred Million Dollars",
        999_999_999.99      : "Nine Hundred Ninety-Nine Million Nine Hundred Ninety-Nine Thousand Nine Hundred Ninety-Nine Dollars and Ninety-Nine Cents",
        1_000_000_000.00    : "One Billion Dollars",
        999_999_999_999.99  : "Nine Hundred Ninety-Nine Billion Nine Hundred Ninety-Nine Million Nine Hundred Ninety-Nine Thousand Nine Hundred Ninety-Nine Dollars and Ninety-Nine Cents",
     }

    for k, v := range tests {
        if NumberToCurrencyConverter(k) != v {
            t.Fatalf("Received \x1b[;31m%s\x1b[;0m for %v, it should be \x1b[;32m%s\x1b[;0m\n", NumberToCurrencyConverter(k), k, v)
        }
        // Negatives
        if NumberToCurrencyConverter(-k) != "Minus " + v {
            t.Fatalf("Received \x1b[;31m%s\x1b[;0m for %v, it should be \x1b[;32m%s\x1b[;0m\n", "Minus " + NumberToCurrencyConverter(k), k, v)
        }
    }
}
