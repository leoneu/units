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
)

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

/* Operations */
