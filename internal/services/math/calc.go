// calc.go
package math

import "errors"

type Calculator interface {
	Add(a, b int) int
	Multiply(a, b int) int
	Divide(a, b int) int
}

type MathService struct {
	calculator Calculator
}

func NewMathService(c Calculator) *MathService {
	return &MathService{calculator: c}
}

func (m *MathService) ComputeExpression(a, b int) int {
	return m.calculator.Add(a, b) + m.calculator.Multiply(a, b)
}

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}
