package units

// float64 representation of volume in cubic meters.
type Volume float64

const (
	Milliliter        = 1e-6 * CubicMeter
	Deciliter         = 1e-4 * CubicMeter
	Liter             = 1e-3 * CubicMeter
	CubicMeter Volume = 1
	CubicInch         = 1.63871e-5 * CubicMeter
	CubicFoot         = 0.0283168 * CubicMeter
	CubicYard         = 0.764554858 * CubicMeter
	FluidOunce        = 2.95735e-5 * CubicMeter
	Quart             = 0.000946353 * CubicMeter
	Gallon            = 0.00378541 * CubicMeter
)

func (v Volume) Milliliters() float64 {
	return float64(v / Milliliter)
}
func (v Volume) Deciliters() float64 {
	return float64(v / Deciliter)
}
func (v Volume) Liters() float64 {
	return float64(v / Liter)
}
func (v Volume) CubicMeters() float64 {
	return float64(v / CubicMeter)
}
func (v Volume) CubicInches() float64 {
	return float64(v / CubicInch)
}
func (v Volume) CubicFeet() float64 {
	return float64(v / CubicFoot)
}
func (v Volume) CubicYards() float64 {
	return float64(v / CubicYard)
}
func (v Volume) FluidOunces() float64 {
	return float64(v / FluidOunce)
}
func (v Volume) Quarts() float64 {
	return float64(v / Quart)
}
func (v Volume) Gallons() float64 {
	return float64(v / Gallon)
}

/* Operations */
