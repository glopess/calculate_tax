package application

import "CalculateTax/models"

type BrazilianTaxRate struct {
	next TaxCalculate
}

func (btr *BrazilianTaxRate) Execute(m *models.Salary) {

	var taxRate, discount = GetIRTaxRateAndDiscount(m)

	if taxRate > 0 {
		var tax = models.NewTax((m.GrossSalary*taxRate)-discount, taxRate, "IR")
		m.AddTax(*tax)
	}

	if btr.next != nil {
		btr.next.Execute(m)
	}
}

func GetIRTaxRateAndDiscount(m *models.Salary) (float64, float64) {
	var taxRate float64
	var discount float64

	if m.GrossSalary > 2259 && m.GrossSalary < 2827 {
		taxRate = 0.075
		discount = 169.44
	} else if m.GrossSalary > 2827 && m.GrossSalary < 3752 {
		taxRate = 0.15
		discount = 381.44
	} else if m.GrossSalary > 3752 && m.GrossSalary < 4665 {
		taxRate = 0.225
		discount = 662.77
	} else if m.GrossSalary > 4665 {
		taxRate = 0.275
		discount = 896
	}
	return taxRate, discount
}

func (btr *BrazilianTaxRate) setNext(next TaxCalculate) {
	btr.next = next
}
