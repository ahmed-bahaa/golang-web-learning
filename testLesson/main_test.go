package main

import "testing"

var test = []struct {
	name     string
	divident float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{name: "valida-data", divident: 10.0, divisor: 10.0, expected: 1, isErr: false},
	{name: "invalida-data", divident: 10.0, divisor: 0.0, expected: 0.0, isErr: true},
	{name: "Expected-5", divident: 10.0, divisor: 2.0, expected: 5.0, isErr: false},
}

func TestDivison(t *testing.T) {
	for _, tt := range test {
		got, err := divide(tt.divident, tt.divisor)
		if tt.isErr {
			if err == nil {
				t.Error("EExpected Error Bud did not get one")
			}
		} else {
			if err != nil {
				t.Error("Did not Expected Error But get one", err.Error())
			}
		}

		if tt.expected != got {
			t.Errorf("Expected %v but got %v", tt.expected, got)
		}

	}
}

func TestDivide(t *testing.T) {
	_, err := divide(10.0, 10.0)
	if err != nil {
		t.Error("Got an error when we shoud not have")
	}
}

func TestBadDivide(t *testing.T) {
	_, err := divide(10.0, 0)
	if err == nil {
		t.Error("Did not Got an error when we shoud")
	}
}
