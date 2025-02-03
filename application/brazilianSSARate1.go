package application

import "CalculateTax/models"

type BrazilianSSARate struct {
	next TaxCalculate
}

func (bsr *BrazilianSSARate) Execute(m *models.Salary) {

	var taxRate, discount = GetTaxRateAndDiscount(m)
	var amount = GetAmount(m)

	if taxRate > 0 && amount > 0 {
		var tax = models.NewTax((amount*taxRate)-discount, taxRate, "INSS")
		m.AddTax(*tax)
	}

	if bsr.next != nil {
		bsr.next.Execute(m)
	}
}

func GetAmount(m *models.Salary) float64 {
	var amount float64

	if m.GrossSalary > 8157 {
		amount = 8157
	} else {
		amount = m.GrossSalary
	}
	return amount
}

func GetTaxRateAndDiscount(m *models.Salary) (float64, float64) {
	var taxRate float64
	var discount float64

	if m.GrossSalary > 1518 && m.GrossSalary < 2794 {
		taxRate = 0.09
		discount = 22.77
	} else if m.GrossSalary > 2794 && m.GrossSalary < 4191 {
		taxRate = 0.12
		discount = 106.59
	} else if m.GrossSalary > 4191 {
		taxRate = 0.14
		discount = 190.4
	}
	return taxRate, discount
}

func (bsr *BrazilianSSARate) setNext(next TaxCalculate) {
	bsr.next = next
}
