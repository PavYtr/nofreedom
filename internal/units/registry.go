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
