package units

type Category string

const (
	Length Category = "length"
	Mass   Category = "mass"
)

type Unit struct {
	Name       string
	Symbol     string
	Category   Category
	Style      string
	BaseFactor float64
}
