package factory

import (
	"errors"
	"fmt"
)

type PayMethod interface {
	Pay(amount float32) string
}

const (
	Cash = iota
	DebitCard
)

func GetPaymentMethod(m int) (PayMethod, error) {
	switch m {
	case Cash:
		return new(CashPM), nil
	case DebitCard:
		return new(NewDebitCardPM), nil
	default:
		return nil, errors.New(fmt.Sprintf("Payment method %d not recognized", m))
	}
}

type CashPM struct{}

func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f payed using cash\n", amount)
}

type DebitPM struct{}

func (d *DebitPM) Pay(amount float32) string {
	return fmt.Sprintf("%#0.2f payed using debit card\n", amount)
}

type NewDebitCardPM struct{}

func (d *NewDebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%#0.2f payed using debit card (new)\n", amount)
}
