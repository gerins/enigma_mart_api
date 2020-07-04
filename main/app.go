package main

import (
	"github.com/vivaldy22/simple-calc-testing/model"
	"github.com/vivaldy22/simple-calc-testing/service"
	"github.com/vivaldy22/simple-calc-testing/view"
)

func main() {
	c := model.Calculator{
		Num1: 1,
		Num2: 2,
	}
	additionService := service.NewAdditionService(c)
	view.NewConsoleView(additionService).ShowResult()
}
