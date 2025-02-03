package application

func GetInstance(country string) TaxCalculate {
	switch country {
	default:
		return GetBrazilianTaxes()
	}
}

func GetBrazilianTaxes() TaxCalculate {
	btr := &BrazilianTaxRate{}
	bsr := &BrazilianSSARate{}

	btr.setNext(bsr)

	return btr
}
