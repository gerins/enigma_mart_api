package service

import (
	"testing"

	"github.com/vivaldy22/simple-calc-testing/model"
)

type modelCalculatorMock struct {
	model.Calculator
	expectedResult float64
}

func TestNewAdditionService(t *testing.T) {
	// Test Table
	tests := []modelCalculatorMock{
		{
			Calculator: model.Calculator{
				Num1: 1,
				Num2: 2,
			},
			expectedResult: 3,
		},
		{
			Calculator: model.Calculator{
				Num1: -1,
				Num2: 2,
			},
			expectedResult: 1,
		},
		{
			Calculator: model.Calculator{
				Num1: 1,
				Num2: -2,
			},
			expectedResult: -1,
		},
		{
			Calculator: model.Calculator{
				Num1: -1,
				Num2: -2,
			},
			expectedResult: -3,
		},
	}
	for _, test := range tests {
		t.Run("It should return Calculator service", func(t *testing.T) {
			calculatorService := NewAdditionService(test.Calculator)
			if calculatorService == nil {
				t.Error("Cannot be null")
			}
		})
	}
}

func TestAdditionService_Result(t *testing.T) {
	tests := []modelCalculatorMock{
		{
			Calculator: model.Calculator{
				Num1: 1,
				Num2: 2,
			},
			expectedResult: 3,
		},
		{
			Calculator: model.Calculator{
				Num1: -1,
				Num2: 2,
			},
			expectedResult: 1,
		},
		{
			Calculator: model.Calculator{
				Num1: 1,
				Num2: -2,
			},
			expectedResult: -1,
		},
		{
			Calculator: model.Calculator{
				Num1: -1,
				Num2: -2,
			},
			expectedResult: -3,
		},
	}
	for _, test := range tests {
		t.Run("It should return Calculator service", func(t *testing.T) {
			calculatorService := NewAdditionService(test.Calculator)

			expected := test.expectedResult
			actual := calculatorService.Result()

			if actual != expected {
				t.Error("Calculation is failed")
			}
		})
	}
}

func TestNewSubtractionService(t *testing.T) {
	// Test Table
	test := modelCalculatorMock{}

	t.Run("It should return Calculator service", func(t *testing.T) {
		calculatorService := NewSubtractionService(test.Calculator)
		if calculatorService == nil {
			t.Error("Cannot be null")
		}
	})
}
