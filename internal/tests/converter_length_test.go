package tests

import (
	"math"
	"testing"

	"github.com/PavYtr/nofreedom/internal/converter"
)

const tolerance = 0.001

func TestConvertLengthNormalToFreedom(t *testing.T) {
	value := 451.0
	from := "m"
	to := "mi"
	expected := 0.280238

	result, err := converter.Convert(value, from, to)
	if err != nil {
		t.Errorf("Should not produce an error")
	}
	if math.Abs(result-expected) > tolerance {
		t.Errorf("Result was incorrect, got %f, want: %f.", result, expected)
	}
}

func TestConvertLengthFreedomToNormal(t *testing.T) {
	value := 13.0
	from := "mi"
	to := "km"
	expected := 20.9215

	result, err := converter.Convert(value, from, to)
	if err != nil {
		t.Errorf("Should not produce an error")
	}
	if math.Abs(result-expected) > tolerance {
		t.Errorf("Result was incorrect, got %f, want: %f.", result, expected)
	}
}

func TestConvertLengthUnknownUnit(t *testing.T) {
	value := 13.0
	from := "mi"
	to := "some_bs_unit"

	result, err := converter.Convert(value, from, to)
	if err == nil {
		t.Errorf("Should produce an error")
	}
	if _, ok := err.(converter.ErrUnknownUnit); !ok {
		t.Errorf("Error should be of type ErrUnknownUnit")
	}
	if result != 0 {
		t.Errorf("Result should be 0, got: %f", result)
	}
}

func TestConvertLengthIncompatibleUnits(t *testing.T) {
	value := 13.0
	from := "mi"
	to := "kg"

	result, err := converter.Convert(value, from, to)
	if err == nil {
		t.Errorf("Should produce an error")
	}
	if _, ok := err.(converter.ErrIncompatibleUnits); !ok {
		t.Errorf("Error should be of type ErrIncompatibleUnits")
	}
	if result != 0 {
		t.Errorf("Result should be 0, got: %f", result)
	}
}

func TestConvertLengthSameKindOfUnitFreedom(t *testing.T) {
	value := 2077.0
	from := "mi"
	to := "ft"
	expected := 10966560.0

	result, err := converter.Convert(value, from, to)
	if err != nil {
		t.Errorf("Should not produce an error")
	}
	if result != expected {
		t.Errorf("Result was incorrect, got %f, want: %f.", result, expected)
	}
}

func TestConvertLengthSameKindOfUnitNormal(t *testing.T) {
	value := 2077.0
	from := "km"
	to := "m"
	expected := 2077000.0

	result, err := converter.Convert(value, from, to)
	if err != nil {
		t.Errorf("Should not produce an error")
	}
	if math.Abs(result-expected) > tolerance {
		t.Errorf("Result was incorrect, got %f, want: %f.", result, expected)
	}
}
