package application

import "CalculateTax/models"

type TaxCalculate interface {
	Execute(*models.Salary)
	setNext(TaxCalculate)
}
