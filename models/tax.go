package models

type Tax struct {
	Amount  float64
	Type    string
	TaxRate float64
}

func NewTax(amount float64, taxRate float64, tp string) *Tax {
	return &Tax{
		Amount:  amount,
		TaxRate: taxRate,
		Type:    tp,
	}
}
