package api

import (
	"CalculateTax/application"
	"CalculateTax/dtos"
	"CalculateTax/models"
	"errors"
)

type Resolver struct{}

func (r *Resolver) CalculateTax(country string, salaryPerMonth float64) (*dtos.SalaryDto, error) {
	if country == "" || salaryPerMonth < 0 {
		return nil, errors.New("name and email cannot be empty")
	}

	var salary = models.NewSalary(salaryPerMonth)

	var instanceCalculateTax = application.GetInstance("brazil")

	instanceCalculateTax.Execute(salary)

	return dtos.MapSalaryDto(salary), nil
}
