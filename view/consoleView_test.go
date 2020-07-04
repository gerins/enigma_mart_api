package view

import (
	"reflect"
	"testing"

	"github.com/vivaldy22/simple-calc-testing/service"
)

func TestNewConsoleView(t *testing.T) {
	type args struct {
		service service.ICalculatorService
	}
	tests := []struct {
		name string
		args args
		want ICalculatorView
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConsoleView(tt.args.service); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConsoleView() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_consoleView_ShowResult(t *testing.T) {
	tests := []struct {
		name    string
		c       *consoleView
		want    float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.ShowResult()
			if (err != nil) != tt.wantErr {
				t.Errorf("consoleView.ShowResult() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("consoleView.ShowResult() = %v, want %v", got, tt.want)
			}
		})
	}
}
