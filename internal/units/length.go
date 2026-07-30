package units

// Normal(metric) units
// Base length unit is meter
var Meter = Unit{
	Name:      "meter",
	Symbol:    "m",
	Category:  Length,
	BaseFactor: 1.0,
}

var Kilometer = Unit{
	Name:      "kilometer",
	Symbol:    "km",
	Category:  Length,
	BaseFactor: 1000.0,
}

var Centimeter = Unit{
	Name:      "centimeter",
	Symbol:    "cm",
	Category:  Length,
	BaseFactor: 0.01,
}

var Millimeter = Unit{
	Name:      "millimeter",
	Symbol:    "mm",
	Category:  Length,
	BaseFactor: 0.001,
}

// Freedom(imperial) units
var Foot = Unit{
	Name:      "foot",
	Symbol:    "ft",
	Category:  Length,
	BaseFactor: 0.3048,
}

var Mile = Unit{
	Name:      "mile",
	Symbol:    "mi",
	Category:  Length,
	BaseFactor: 1609.344,
}

var Inch = Unit{
	Name:      "inch",
	Symbol:    "in",
	Category:  Length,
	BaseFactor: 0.0254,
}

var Yard = Unit{
	Name:      "yard",
	Symbol:    "yd",
	Category:  Length,
	BaseFactor: 0.9144,
}
