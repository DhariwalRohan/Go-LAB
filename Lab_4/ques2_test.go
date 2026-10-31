package main

import "testing"

func TestDivide(t *testing.T) {
	result, err := divide(10, 2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != 5 {
		t.Errorf("Expected 5, got %v", result)
	}
}

func TestSquareRoot(t *testing.T) {
	result, err := squareRoot(25)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != 5 {
		t.Errorf("Expected 5, got %v", result)
	}
}

func TestValidateInput(t *testing.T) {
	err := validateInput("")
	if err == nil {
		t.Errorf("Expected error for empty input")
	}
}
