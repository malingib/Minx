package money

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrInvalidCurrency  = errors.New("currency must be a three-letter uppercase code")
	ErrCurrencyMismatch = errors.New("currency mismatch")
	ErrAmountOverflow   = errors.New("amount overflow")
)

// Amount stores money in integer minor units; it never uses floating point.
type Amount struct {
	MinorUnits int64
	Currency   string
}

func New(minorUnits int64, currency string) (Amount, error) {
	if len(currency) != 3 || currency != strings.ToUpper(currency) {
		return Amount{}, ErrInvalidCurrency
	}
	for _, r := range currency {
		if r < 'A' || r > 'Z' {
			return Amount{}, ErrInvalidCurrency
		}
	}
	return Amount{MinorUnits: minorUnits, Currency: currency}, nil
}

func (a Amount) Add(b Amount) (Amount, error) {
	if a.Currency != b.Currency {
		return Amount{}, ErrCurrencyMismatch
	}
	if (b.MinorUnits > 0 && a.MinorUnits > int64(^uint64(0)>>1)-b.MinorUnits) ||
		(b.MinorUnits < 0 && a.MinorUnits < -int64(^uint64(0)>>1)-1-b.MinorUnits) {
		return Amount{}, ErrAmountOverflow
	}
	return Amount{MinorUnits: a.MinorUnits + b.MinorUnits, Currency: a.Currency}, nil
}

type Direction string

const (
	Debit  Direction = "DEBIT"
	Credit Direction = "CREDIT"
)

type Line struct {
	Amount    Amount
	Direction Direction
}

// ValidateJournal enforces positive lines and balanced totals per currency.
func ValidateJournal(lines []Line) error {
	if len(lines) < 2 {
		return errors.New("journal requires at least two lines")
	}
	totals := make(map[string]int64)
	for _, line := range lines {
		if line.Amount.MinorUnits <= 0 {
			return errors.New("journal line amount must be positive")
		}
		if line.Direction != Debit && line.Direction != Credit {
			return fmt.Errorf("invalid journal direction %q", line.Direction)
		}
		current := totals[line.Amount.Currency]
		if line.Direction == Debit {
			if current > int64(^uint64(0)>>1)-line.Amount.MinorUnits {
				return ErrAmountOverflow
			}
			totals[line.Amount.Currency] = current + line.Amount.MinorUnits
		} else {
			if current < -int64(^uint64(0)>>1)-1+line.Amount.MinorUnits {
				return ErrAmountOverflow
			}
			totals[line.Amount.Currency] = current - line.Amount.MinorUnits
		}
	}
	for currency, total := range totals {
		if total != 0 {
			return fmt.Errorf("unbalanced journal for %s", currency)
		}
	}
	return nil
}
