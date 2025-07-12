# Monex

🪙 A Go library to parse money strings like `"1 200,50 грн"` or `"$3,999.99"` into float and currency.

## Usage

```go
amount, currency, err := money.ParseAmount("1 250,75 грн")
// amount: 1250.75
// currency: "UAH"