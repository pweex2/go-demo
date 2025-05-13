// mathsvc_test.go
package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// 创建 Mock 对象
type MockCalculator struct {
	mock.Mock
}

func (m *MockCalculator) Add(a, b int) int {
	args := m.Called(a, b)
	return args.Int(0)
}

func (m *MockCalculator) Multiply(a, b int) int {
	args := m.Called(a, b)
	return args.Int(0)
}

func (m *MockCalculator) Divide(a, b int) int {
	args := m.Called(a, b)
	return args.Int(0)
}

func TestComputeExpression(t *testing.T) {
	mockCalc := new(MockCalculator)

	// 设置期望：传入 2 和 3，返回 5 和 6
	mockCalc.On("Add", 2, 3).Return(5)
	mockCalc.On("Multiply", 2, 3).Return(6)
	// mockCalc.On("Divide", 2, 3).Return(2)

	service := NewMathService(mockCalc)

	result := service.ComputeExpression(2, 3)

	assert.Equal(t, 11, result)
	mockCalc.AssertExpectations(t) // 验证方法是否真的被调用
}
