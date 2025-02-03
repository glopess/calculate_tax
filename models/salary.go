package models

type Salary struct {
	GrossSalary float64
	Taxes       []Tax
}

func NewSalary(grossSalary float64) *Salary {
	return &Salary{
		GrossSalary: grossSalary,
	}
}

func (s *Salary) AddTax(tax Tax) {
	s.Taxes = append(s.Taxes, tax)
}

func (s *Salary) GetNetSalary() float64 {
	return s.GrossSalary - s.GetDiscounts()
}

func (s *Salary) GetDiscounts() float64 {
	var sum float64

	for _, tax := range s.Taxes {
		sum += tax.Amount
	}

	return sum
}
