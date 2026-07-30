package units

// Base mass unit is kilogram
var Kilogram = Unit{
	Name:       "kilogram",
	Symbol:     "kg",
	Category:   Mass,
	BaseFactor: 1.0,
}

var Gram = Unit{
	Name:       "gram",
	Symbol:     "g",
	Category:   Mass,
	BaseFactor: 0.001,
}

var Milligram = Unit{
	Name:       "milligram",
	Symbol:     "mg",
	Category:   Mass,
	BaseFactor: 0.000001,
}

// Freedom units
var Pound = Unit{
	Name:       "pound",
	Symbol:     "lb",
	Category:   Mass,
	BaseFactor: 0.45359237,
}

var Ounce = Unit{
	Name:       "ounce",
	Symbol:     "oz",
	Category:   Mass,
	BaseFactor: 0.028349523125,
}
