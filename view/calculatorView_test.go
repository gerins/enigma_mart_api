package view

import "testing"

type serviceMock struct {
	expectedResult float64
}

func (s *serviceMock) Result() float64 {
	return s.expectedResult
}

func TestNewConsoleViews(t *testing.T) {
	t.Run("It should return Console view", func(t *testing.T) {
		consoleView := NewConsoleView(&serviceMock{})

		if consoleView == nil {
			t.Error("Can not be nil")
		}
	})
}

func TestConsoleView_ShowResult(t *testing.T) {
	tests := []struct {
		serviceMock
	}{
		{serviceMock{10}},
		{serviceMock{-10}},
	}
	for _, test := range tests {
		t.Run("It should show the result", func(t *testing.T) {
			consoleView := NewConsoleView(&test)
			res, err := consoleView.ShowResult()

			if res != test.expectedResult {
				t.Error("Result is wrong")
			}

			if err != nil {
				t.Error("Error is occurred")
			}
		})
	}
}
