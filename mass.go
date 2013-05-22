package units

// float64 representation of mass in kilograms.
type Mass float64

const (
	Microgram      = 1e-9 * Kilogram
	Milligram      = 1e-6 * Kilogram
	Gram           = 1e-3 * Kilogram
	Kilogram  Mass = 1
	Megagram       = 1e3 * Kilogram
	Tonne          = 1e3 * Kilogram
	Pound          = 0.45359237 * Kilogram
	Ounce          = 0.02834952312 * Kilogram
	Stone          = 6.35029 * Kilogram
)

var mass = map[string]Mass{
	"microgram": Microgram,
	"\u03BCg":   Microgram,
	"milligram": Milligram,
	"mg":        Milligram,
	"gram":      Gram,
	"g":         Gram,
	"kilogram":  Kilogram,
	"kg":        Kilogram,
	"megagram":  Megagram,
	"tonne":     Tonne,
	"t":         Tonne,
	"pound":     Pound,
	"lb":        Pound,
	"ounce":     Ounce,
	"oz":        Ounce,
	"stone":     Stone,
}

// Creates a new mass.
func NewMass(v float64, s string) Mass {
	return Mass(v) * mass[s]
}

func (m Mass) Micrograms() float64 {
	return float64(m / Microgram)
}

func (m Mass) Milligrams() float64 {
	return float64(m / Milligram)
}

func (m Mass) Grams() float64 {
	return float64(m / Gram)
}
func (m Mass) Kilograms() float64 {
	return float64(m / Kilogram)
}

func (m Mass) Megagrams() float64 {
	return float64(m / Megagram)
}

func (m Mass) Tonnes() float64 {
	return float64(m / Tonne)
}

func (m Mass) Pounds() float64 {
	return float64(m / Pound)
}

func (m Mass) Ounces() float64 {
	return float64(m / Ounce)
}

func (m Mass) Stones() float64 {
	return float64(m / Stone)
}

/* Operations */

func (m Mass) MultiplyWithAcceleration(a Acceleration) Force {
	return Force(float64(m) * float64(a))
}
