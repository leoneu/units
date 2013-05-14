package units

import (
	"time"
)

// float64 representation of SI derived units.
type Area float64
type Acceleration float64
type MassDensity float64
type Frequency float64
type Force float64

const (
	SquareMillimeter      = 1e-6 * SquareMeter
	SquareMeter      Area = 1
	SquareKilometer       = 1e6 * SquareMeter
	SquareFoot            = 0.092903 * SquareMeter
)

const (
	MeterPerSecondSquared Acceleration = 1
)

const (
	KilogramPerCubicMeter MassDensity = 1
)

const (
	Hertz Frequency = 1
)

const (
	Newton Force = 1
)

func (a Area) SquareMillimeters() float64 {
	return float64(a / SquareMillimeter)
}
func (a Area) SquareMeters() float64 {
	return float64(a / SquareMeter)
}
func (a Area) SquareKilometers() float64 {
	return float64(a / SquareKilometer)
}

/* Operations */

func (f Frequency) Period() time.Duration {
	return time.Duration(1e9/(float64(f))) * time.Nanosecond
}
