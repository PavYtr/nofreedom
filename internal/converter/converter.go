package converter

import "github.com/PavYtr/nofreedom/internal/units"

func Convert(value float64, from string, to string) (float64, error) {
	fromUnit, ok := units.Registry[from]
	if !ok {
		return 0, ErrUnknownUnit{Unit: from}
	}
	toUnit, ok := units.Registry[to]
	if !ok {
		return 0, ErrUnknownUnit{Unit: to}
	}
	if fromUnit.Category != toUnit.Category {
		return 0, ErrIncompatibleUnits{ToUnit: to, FromUnit: from}
	}

	result := value * fromUnit.BaseFactor / toUnit.BaseFactor
	return result, nil
}
