package dtos

type SalaryDto struct {
	GrossSalary      float64
	NetSalary        float64
	AnualGrossSalary float64
	AnualNetSalary   float64
	Taxes            []TaxDto
}
