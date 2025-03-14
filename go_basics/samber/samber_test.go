package samber_test

// https://github.com/samber/lo?tab=readme-ov-file

import (
	"testing"

	"github.com/samber/lo"
)

func TestSamberCount(t *testing.T) {
	// Arrange
	expected := 2

	// Act (Ammount of 1's in the array)
	r := lo.Count([]int{1, 2, 1}, 1)

	// Assert
	if r != expected {
		t.Errorf("Expected %v, got %v", expected, r)
	}
}
