// calc_test.go
package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 1, 2, 3},
		{"negative numbers", -1, -2, -3},
		{"mixed numbers", -1, 2, 1},
		{"zero", 0, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 5, 3, 2},
		{"negative result", 3, 5, -2},
		{"zero result", 4, 4, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Subtract(tt.a, tt.b)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMultiply(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"positive", args{2, 3}, 6},
		{"zero", args{5, 0}, 0},
		{"negative", args{-2, 3}, -6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Multiply(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiply2(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive", 2, 3, 6},
		{"zero", 5, 0, 0},
		{"negative", -2, 3, -6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, Multiply(tt.a, tt.b))
		})
	}
}
func TestDivide(t *testing.T) {
	tests := []struct {
		name        string
		a, b        int
		expected    int
		expectError bool
	}{
		{"normal", 6, 3, 2, false},
		{"division by zero", 5, 0, 0, true},
		{"negative division", -6, 2, -3, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Divide(tt.a, tt.b)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}
