package converter

type ErrUnknownUnit struct {
	Unit string
}

type ErrIncompatubleUnits struct {
	ToUnit   string
	FromUnit string
}

func (e ErrUnknownUnit) Error() string {
	return "unknown unit: " + e.Unit
}

func (e ErrIncompatubleUnits) Error() string {
	return "incompatible units: " + e.FromUnit + " and " + e.ToUnit
}
