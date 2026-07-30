package tests

import (
	"math"
	"testing"

	"github.com/PavYtr/nofreedom/internal/converter"
)

func TestConvertMassNormalToNormal(t *testing.T) {
	value := 1000.0
	from := "kg"
	to1 := "g"
	to2 := "mg"
	expected1 := 1000000.0
	expected2 := 1000000000.0

	result1, err := converter.Convert(value, from, to1)
	if err != nil {
		t.Errorf("Should not produce an error")
	}
	if math.Abs(result1-expected1) > tolerance {
		t.Errorf("Result was incorrect, got %f, want: %f.", result1, expected1)
	}

	result2, err := converter.Convert(value, from, to2)
	if err != nil {
		t.Errorf("Should not produce an error")
	}
	if math.Abs(result2-expected2) > tolerance {
		t.Errorf("Result was incorrect, got %f, want: %f.", result2, expected2)
	}
}

func TestConvertMassNormalToFreedom(t *testing.T) {
	value := 1.0
	from := "kg"
	to1 := "lb"
	to2 := "oz"
	expected1 := 2.20462
	expected2 := 35.274

	result1, err := converter.Convert(value, from, to1)
	if err != nil {
		t.Errorf("Should not produce an error")
	}
	if math.Abs(result1-expected1) > tolerance {
		t.Errorf("Result was incorrect, got %f, want: %f.", result1, expected1)
	}

	result2, err := converter.Convert(value, from, to2)
	if err != nil {
		t.Errorf("Should not produce an error")
	}
	if math.Abs(result2-expected2) > tolerance {
		t.Errorf("Result was incorrect, got %f, want: %f.", result2, expected2)
	}
}

func TestConvertMassFreedomToNormal(t *testing.T) {
	value := 10.0
	from := "lb"
	to1 := "kg"
	to2 := "g"
	expected1 := 4.53592
	expected2 := 4535.9237

	result1, err := converter.Convert(value, from, to1)
	if err != nil {
		t.Errorf("Should not produce an error")
	}
	if math.Abs(result1-expected1) > tolerance {
		t.Errorf("Result was incorrect, got %f, want: %f.", result1, expected1)
	}

	result2, err := converter.Convert(value, from, to2)
	if err != nil {
		t.Errorf("Should not produce an error")
	}
	if math.Abs(result2-expected2) > tolerance {
		t.Errorf("Result was incorrect, got %f, want: %f.", result2, expected2)
	}
}
