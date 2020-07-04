package view

import (
	"log"

	"github.com/vivaldy22/simple-calc-testing/service"
)

type consoleView struct {
	calculatorService service.ICalculatorService
}

func NewConsoleView(service service.ICalculatorService) ICalculatorView {
	return &consoleView{service}
}

func (c *consoleView) ShowResult() (float64, error) {
	res := c.calculatorService.Result()
	log.Print(res)
	return res, nil
}
