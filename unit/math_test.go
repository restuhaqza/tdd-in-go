package unit_test

import (
	"testing"

	"github.com/restuhaqza/tdd-in-go/unit"
)

func TestAdd(t *testing.T) {
	executor := unit.New()
	t.Run("1 + 2 = 3", func(t *testing.T) {
		result := executor.Add(1, 2)

		if result != 3 {
			t.Errorf("Expected 3, got %d", result)
		}
	})
}

func TestSubtract(t *testing.T) {
	executor := unit.New()
	t.Run("2 - 1 = 1", func(t *testing.T) {
		result := executor.Subtract(2, 1)

		if result != 1 {
			t.Errorf("Expected 1, got %d", result)
		}
	})
}

func TestMultiply(t *testing.T) {
	executor := unit.New()
	t.Run("2 * 2 = 4", func(t *testing.T) {
		result := executor.Multiply(2, 2)

		if result != 4 {
			t.Errorf("Expected 4, got %d", result)
		}
	})
}

func TestDivide(t *testing.T) {
	executor := unit.New()
	t.Run("4 / 2 = 2", func(t *testing.T) {
		result := executor.Divide(4, 2)

		if result != 2 {
			t.Errorf("Expected 2, got %d", result)
		}
	})
}
