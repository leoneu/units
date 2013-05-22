package units

import (
	"math"
	"time"
)

// float64 representation of length in meters.
type Length float64

const (
	Nanometer         = 1e-9 * Meter
	Micrometer        = 1e-6 * Meter
	Millimeter        = 1e-3 * Meter
	Centimeter        = 1e-2 * Meter
	Decimeter         = 1e-1 * Meter
	Meter      Length = 1
	Kilometer         = 1e3 * Meter
	Yard              = 0.9144 * Meter
	Mile              = 1609.34 * Meter
	Foot              = 0.3048 * Meter
	Inch              = 0.0254 * Meter
)

var length = map[string]Length{
	"nanometer":  Nanometer,
	"nm":         Nanometer,
	"micrometer": Micrometer,
	"\u03BCm":    Micrometer,
	"millimeter": Millimeter,
	"mm":         Millimeter,
	"centimeter": Centimeter,
	"cm":         Centimeter,
	"meter":      Meter,
	"m":          Meter,
	"kilometer":  Kilometer,
	"km":         Kilometer,
	"yard":       Yard,
	"yd":         Yard,
	"mile":       Mile,
	"mi":         Mile,
	"foot":       Foot,
	"ft":         Foot,
	"inch":       Inch,
	"in":         Inch,
}

// Creates a new length.
func NewLength(v float64, s string) Length {
	return Length(v) * length[s]
}

// Converts length to Nanometers.
func (d Length) Nanometers() float64 {
	return float64(d / Nanometer)
}

func (d Length) Micrometers() float64 {
	return float64(d / Micrometer)
}

func (d Length) Millimeters() float64 {
	return float64(d / Millimeter)
}

func (d Length) Centimeters() float64 {
	return float64(d / Centimeter)
}

func (d Length) Meters() float64 {
	return float64(d / Meter)
}

func (d Length) Kilometers() float64 {
	return float64(d / Kilometer)
}

func (d Length) Yards() float64 {
	return float64(d / Yard)
}

func (d Length) Miles() float64 {
	return float64(d / Mile)
}

func (d Length) Feet() float64 {
	return float64(d / Foot)
}

func (d Length) Inches() float64 {
	return float64(d / Inch)
}

/* Operations */

func (d Length) Abs() Length {
	if d < 0 {
		return -d
	}
	return d
}

func (d Length) DivideWithDuration(t time.Duration) Velocity {
	return Velocity(float64(d) / float64(t) * float64(time.Second))
}

func Hypot(x, y Length) Length {
	return Length(math.Hypot(float64(x), float64(y)))
}

func (d Length) MultiplyWithLength(dist Length) Area {
	return Area(float64(d) * float64(dist))
}
