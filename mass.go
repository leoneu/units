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
	Pound          = 0.453592 * Kilogram
	Ounce          = 0.0283495 * Kilogram
)

func (m Mass) Microgram() float64 {
	return float64(m / Microgram)
}

func (m Mass) Milligram() float64 {
	return float64(m / Milligram)
}

func (m Mass) Gram() float64 {
	return float64(m / Gram)
}
func (m Mass) Kilogram() float64 {
	return float64(m / Kilogram)
}

func (m Mass) Megagram() float64 {
	return float64(m / Megagram)
}

func (m Mass) Tonne() float64 {
	return float64(m / Tonne)
}

func (m Mass) Pound() float64 {
	return float64(m / Pound)
}

func (m Mass) Ounce() float64 {
	return float64(m / Ounce)
}

/* Operations */
