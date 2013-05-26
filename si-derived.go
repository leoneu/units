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

var area = map[string]Area{
	"square_millimeter": SquareMillimeter,
	"square_mm":         SquareMillimeter,
	"square_meter":      SquareMeter,
	"square_m":          SquareMeter,
	"square_kilometer":  SquareKilometer,
	"square_km":         SquareKilometer,
	"square_foot":       SquareFoot,
	"square_ft":         SquareFoot,
}

// Creates a new area.
func NewArea(v float64, s string) Area {
	return Area(v) * area[s]
}

func (a Area) SquareMillimeters() float64 {
	return float64(a / SquareMillimeter)
}
func (a Area) SquareMeters() float64 {
	return float64(a / SquareMeter)
}
func (a Area) SquareKilometers() float64 {
	return float64(a / SquareKilometer)
}

const (
	MeterPerSecondSquared Acceleration = 1
)

var acceleration = map[string]Acceleration{
	"meter_per_second_squared": MeterPerSecondSquared,
	"m_per_s_squared":          MeterPerSecondSquared,
	"m_per_s_sq":               MeterPerSecondSquared,
}

// Creates a new acceleration.
func NewAcceleration(v float64, s string) Acceleration {
	return Acceleration(v) * acceleration[s]
}

const (
	KilogramPerCubicMeter MassDensity = 1
)

var massDensity = map[string]MassDensity{
	"kilogram_per_cubic_meter": KilogramPerCubicMeter,
	"kg_per_cubic_m":           KilogramPerCubicMeter,
}

// Creates a new mass density.
func NewMassDensity(v float64, s string) MassDensity {
	return MassDensity(v) * massDensity[s]
}

const (
	Hertz     Frequency = 1
	Kilohertz           = 1e3 * Hertz
	Megahertz           = 1e6 * Hertz
	Gigahertz           = 1e9 * Hertz
)

var frequency = map[string]Frequency{
	"hertz":     Hertz,
	"Hz":        Hertz,
	"kilohertz": Kilohertz,
	"kHz":       Kilohertz,
	"megahertz": Megahertz,
	"MHz":       Megahertz,
	"gigahertz": Gigahertz,
	"GHz":       Gigahertz,
}

// Creates a new frequency.
func NewFrequency(v float64, s string) Frequency {
	return Frequency(v) * frequency[s]
}

func (f Frequency) Hertz() float64 {
	return float64(f / Hertz)
}
func (f Frequency) Kilohertz() float64 {
	return float64(f / Kilohertz)
}
func (f Frequency) Megahertz() float64 {
	return float64(f / Megahertz)
}
func (f Frequency) Gigahertz() float64 {
	return float64(f / Gigahertz)
}

const (
	Newton Force = 1
)

var force = map[string]Force{
	"newton": Newton,
	"N":      Newton,
}

// Creates a new mass force.
func NewForce(v float64, s string) Force {
	return Force(v) * force[s]
}

/* Operations */

func (f Frequency) Period() time.Duration {
	return time.Duration(1e9/(float64(f))) * time.Nanosecond
}
