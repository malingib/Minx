package money

import (
	"errors"
	"math"
	"testing"
)

func TestAmountAdd(t *testing.T) {
	a, _ := New(125, "KES")
	b, _ := New(75, "KES")
	got, err := a.Add(b)
	if err != nil || got.MinorUnits != 200 {
		t.Fatalf("Add() = %+v, %v", got, err)
	}
}

func TestAmountAddRejectsCurrencyMismatch(t *testing.T) {
	a, _ := New(100, "KES")
	b, _ := New(100, "USD")
	if _, err := a.Add(b); !errors.Is(err, ErrCurrencyMismatch) {
		t.Fatalf("expected currency mismatch, got %v", err)
	}
}

func TestAmountAddRejectsOverflow(t *testing.T) {
	a, _ := New(math.MaxInt64, "KES")
	b, _ := New(1, "KES")
	if _, err := a.Add(b); !errors.Is(err, ErrAmountOverflow) {
		t.Fatalf("expected overflow, got %v", err)
	}
}

func TestValidateJournal(t *testing.T) {
	debit, _ := New(1000, "KES")
	credit, _ := New(700, "KES")
	credit2, _ := New(300, "KES")
	if err := ValidateJournal([]Line{
		{Amount: debit, Direction: Debit},
		{Amount: credit, Direction: Credit},
		{Amount: credit2, Direction: Credit},
	}); err != nil {
		t.Fatalf("valid journal rejected: %v", err)
	}
}

func TestValidateJournalRejectsUnbalancedCurrency(t *testing.T) {
	debit, _ := New(1000, "KES")
	credit, _ := New(1000, "USD")
	if err := ValidateJournal([]Line{
		{Amount: debit, Direction: Debit},
		{Amount: credit, Direction: Credit},
	}); err == nil {
		t.Fatal("expected currency-level imbalance")
	}
}
