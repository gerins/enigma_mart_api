package service

import (
	"github.com/vivaldy22/simple-calc-testing/model"
)

type additionService struct {
	c model.Calculator
}

func NewAdditionService(c model.Calculator) ICalculatorService {
	return &additionService{c}
}

func (a *additionService) Result() float64 {
	return a.c.Num1 + a.c.Num2
}
