package service

import (
	"reflect"
	"testing"

	"github.com/vivaldy22/simple-calc-testing/model"
)

func TestNewAdditionService(t *testing.T) {
	type args struct {
		c model.Calculator
	}
	tests := []struct {
		name string
		args args
		want ICalculatorService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAdditionService(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAdditionService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_additionService_Result(t *testing.T) {
	tests := []struct {
		name string
		a    *additionService
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Result(); got != tt.want {
				t.Errorf("additionService.Result() = %v, want %v", got, tt.want)
			}
		})
	}
}
