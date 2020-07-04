package service

import (
	"github.com/vivaldy22/simple-calc-testing/model"
)

type subtractionService struct {
	c model.Calculator
}

func NewSubtractionService(c model.Calculator) ICalculatorService {
	return &subtractionService{c}
}

func (a *subtractionService) Result() float64 {
	return a.c.Num1 / a.c.Num2
}
