// go test ./calculator

package calculator

import (
	"testing"
)

func TestDivide(t *testing.T) {
	// Arrange
	expected := 2.0

	// Act
	r, e := Divide(4, 2)
	if e != nil {
		t.Fatal(e)
	}

	// Assert
	if r != expected {
		t.Errorf("Expected %v, got %v", expected, r)
	}
}

func TestDivideByZero(t *testing.T) {
	_, e := Divide(4, 0)
	if e == nil {
		t.Error("Divide by Zero Excepted")
	}
}

func TestDivideNegative(t *testing.T) {
	// Arrange
	expected := -3.0

	// Act
	r, e := Divide(-6, 2)
	if e != nil {
		t.Fatal(e)
	}

	// Assert
	if r != expected {
		t.Errorf("Expected %v, got %v", expected, r)
	}
}
