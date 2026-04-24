package class

import (
	"errors"

	"github.com/shopspring/decimal"
)

type Money interface {
	Add(Money) Money
	Subtract(Money) Money
	IsPositive() bool
	String() string
	getAmount() decimal.Decimal
}

type moneyImpl struct {
	amount decimal.Decimal
}

func NewMoney(amountStr string) (Money, error) {
	d, err := decimal.NewFromString(amountStr)
	if err != nil {
		return moneyImpl{}, err
	}
	return moneyImpl{amount: d}, nil
}

func NewPositiveMoney(amountStr string) (Money, error) {
	d, err := decimal.NewFromString(amountStr)
	if err != nil {
		return moneyImpl{}, err
	}
	if !d.IsPositive() {
		return moneyImpl{}, errors.New("amount must be positive")
	}
	return moneyImpl{amount: d.Abs()}, nil
}

func (m moneyImpl) Add(other Money) Money {
	return moneyImpl{amount: m.amount.Add(other.getAmount())}
}

func (m moneyImpl) Subtract(other Money) Money {
	return moneyImpl{amount: m.amount.Sub(other.getAmount())}
}

func (m moneyImpl) IsPositive() bool {
	return m.amount.IsPositive()
}

func (m moneyImpl) String() string {
	return m.amount.String()
}

func (m moneyImpl) getAmount() decimal.Decimal {
	return m.amount
}
