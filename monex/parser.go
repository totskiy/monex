package monex

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var currencyMap = map[string]string{
	"грн": "UAH", "uah": "UAH", "₴": "UAH",
	"$": "USD", "usd": "USD",
	"€": "EUR", "eur": "EUR",
}

func ParseAmount(input string) (float64, string, error) {
	s := strings.ToLower(strings.TrimSpace(input))
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, " ", "")

	currency := "UAH"
	for symbol, code := range currencyMap {
		if strings.Contains(s, symbol) {
			currency = code
			s = strings.ReplaceAll(s, symbol, "")
		}
	}

	if strings.Contains(s, ",") && !strings.Contains(s, ".") {
		s = strings.ReplaceAll(s, ",", ".")
	}

	re := regexp.MustCompile(`[^\d\.]`)
	s = re.ReplaceAllString(s, "")

	if s == "" {
		return 0, "", errors.New("invalid amount string")
	}

	amount, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, "", err
	}

	return amount, currency, nil
}
