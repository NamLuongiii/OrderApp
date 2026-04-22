package common

import (
	"github.com/shopspring/decimal"
)

type Money struct {
	amount decimal.Decimal
}

func NewMoney(amountStr string) (Money, error) {
	d, err := decimal.NewFromString(amountStr)
	if err != nil {
		return Money{}, err
	}
	return Money{amount: d}, nil
}

func (m Money) Add(other Money) Money {
	return Money{amount: m.amount.Add(other.amount)}
}

func (m Money) Subtract(other Money) Money {
	return Money{amount: m.amount.Sub(other.amount)}
}

func (m Money) IsPositive() bool {
	return m.amount.IsPositive()
}

func (m Money) String() string {
	return m.amount.String()
}
