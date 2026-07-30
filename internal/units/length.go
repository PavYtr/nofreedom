package units

// Base length unit is meter
var Meter = Unit{
	Name:      "meter",
	Symbol:    "m",
	Category:  Length,
	coeffBase: 1.0,
}

var Kilometer = Unit{
	Name:      "kilometer",
	Symbol:    "km",
	Category:  Length,
	coeffBase: 1000.0,
}

var Centimeter = Unit{
	Name:      "centimeter",
	Symbol:    "cm",
	Category:  Length,
	coeffBase: 0.01,
}

var Millimeter = Unit{
	Name:      "millimeter",
	Symbol:    "mm",
	Category:  Length,
	coeffBase: 0.001,
}

// Freedom units
var Foot = Unit{
	Name:      "foot",
	Symbol:    "ft",
	Category:  Length,
	coeffBase: 0.3048,
}

var Mile = Unit{
	Name:      "mile",
	Symbol:    "mi",
	Category:  Length,
	coeffBase: 1609.34,
}

var Inch = Unit{
	Name:      "inch",
	Symbol:    "in",
	Category:  Length,
	coeffBase: 0.0254,
}

var Yard = Unit{
	Name:      "yard",
	Symbol:    "yd",
	Category:  Length,
	coeffBase: 0.9144,
}
