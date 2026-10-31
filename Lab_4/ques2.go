package main

import (
	"fmt"
	"math"
	"strings"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func squareRoot(n float64) (float64, error) {
	if n < 0 {
		return 0, fmt.Errorf("negative number not allowed")
	}
	return math.Sqrt(n), nil
}

func validateInput(s string) error {
	if strings.TrimSpace(s) == "" {
		return fmt.Errorf("input cannot be empty")
	}
	return nil
}
