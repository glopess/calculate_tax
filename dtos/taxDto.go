package dtos

import (
	"CalculateTax/helpers"
	"CalculateTax/models"
)

type TaxDto struct {
	Amount  float64
	Type    string
	TaxRate float64
}

func MapTaxesDto(taxes []models.Tax) []TaxDto {
	var taxesDto []TaxDto
	for _, tax := range taxes {
		taxesDto = append(taxesDto, MapTaxDto(tax))
	}
	return taxesDto
}

func MapTaxDto(tax models.Tax) TaxDto {
	var taxDto = &TaxDto{
		Amount:  helpers.ToTwoPlaces(tax.Amount),
		Type:    tax.Type,
		TaxRate: helpers.ToTwoPlaces(tax.TaxRate),
	}

	return *taxDto
}
