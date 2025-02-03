package api

import (
	"CalculateTax/application"
	"CalculateTax/dtos"
	"CalculateTax/models"
	"errors"
	"math"
)

type Resolver struct{}

func (r *Resolver) CalculateTax(country string, salaryPerMonth float64) (*dtos.SalaryDto, error) {
	if country == "" || salaryPerMonth < 0 {
		return nil, errors.New("name and email cannot be empty")
	}

	var salary = models.NewSalary(salaryPerMonth)

	var instanceCalculateTax = application.GetInstance("brazil")

	instanceCalculateTax.Execute(salary)

	return MapSalaryDto(salary), nil
}

func MapSalaryDto(salary *models.Salary) *dtos.SalaryDto {
	return &dtos.SalaryDto{
		GrossSalary:      ToTwoPlaces(salary.GrossSalary),
		NetSalary:        ToTwoPlaces(salary.GetNetSalary()),
		AnualGrossSalary: ToTwoPlaces(salary.GrossSalary * 12),
		AnualNetSalary:   ToTwoPlaces(salary.GetNetSalary() * 12),
		Taxes:            MapTaxesDto(salary.Taxes),
	}
}

func MapTaxesDto(taxes []models.Tax) []dtos.TaxDto {
	var taxesDto []dtos.TaxDto
	for _, tax := range taxes {
		taxesDto = append(taxesDto, MapTaxDto(tax))
	}
	return taxesDto
}

func MapTaxDto(tax models.Tax) dtos.TaxDto {
	var taxDto = &dtos.TaxDto{
		Amount:  ToTwoPlaces(tax.Amount),
		Type:    tax.Type,
		TaxRate: ToTwoPlaces(tax.TaxRate),
	}

	return *taxDto
}

func ToTwoPlaces(decimal float64) float64 {
	return math.Round(decimal*100) / 100
}
