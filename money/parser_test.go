package money

import "testing"

func TestParseAmount(t *testing.T) {
	tests := []struct {
		input        string
		wantAmount   float64
		wantCurrency string
	}{
		{"1 000,50 грн", 1000.50, "UAH"},
		{"$1,200.99", 1200.99, "USD"},
		{"   99.99 USD", 99.99, "USD"},
		{"€ 9 876,00", 9876.00, "EUR"},
		{"₴100", 100.00, "UAH"},
	}

	for _, tt := range tests {
		gotAmount, gotCurrency, err := ParseAmount(tt.input)
		if err != nil {
			t.Errorf("error parsing %s: %v", tt.input, err)
		}
		if gotAmount != tt.wantAmount || gotCurrency != tt.wantCurrency {
			t.Errorf("ParseAmount(%q) = %f, %s; want %f, %s",
				tt.input, gotAmount, gotCurrency, tt.wantAmount, tt.wantCurrency)
		}
	}
}
