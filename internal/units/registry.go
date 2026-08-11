package units

var Registry = map[string]Unit{
	"m":  Meter,
	"km": Kilometer,
	"mm": Millimeter,
	"cm": Centimeter,
	"mi": Mile,
	"ft": Foot,
	"yd": Yard,
	"in": Inch,

	"kg": Kilogram,
	"g":  Gram,
	"mg": Milligram,
	"lb": Pound,
	"oz": Ounce,
}

var LengthUnits = []Unit{
	Meter,
	Kilometer,
	Centimeter,
	Millimeter,
	Foot,
	Mile,
	Inch,
	Yard,
}

var MassUnits = []Unit{
	Kilogram,
	Gram,
	Milligram,
	Pound,
	Ounce,
}

var RegistryByCategory = [][]Unit{
	LengthUnits,
	MassUnits,
}
