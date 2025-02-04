package dtos

import (
	"CalculateTax/helpers"
	"CalculateTax/models"
)

type SalaryDto struct {
	GrossSalary      float64
	NetSalary        float64
	AnualGrossSalary float64
	AnualNetSalary   float64
	Taxes            []TaxDto
}

func MapSalaryDto(salary *models.Salary) *SalaryDto {
	return &SalaryDto{
		GrossSalary:      helpers.ToTwoPlaces(salary.GrossSalary),
		NetSalary:        helpers.ToTwoPlaces(salary.GetNetSalary()),
		AnualGrossSalary: helpers.ToTwoPlaces(salary.GrossSalary * 12),
		AnualNetSalary:   helpers.ToTwoPlaces(salary.GetNetSalary() * 12),
		Taxes:            MapTaxesDto(salary.Taxes),
	}
}
