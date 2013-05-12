package units

import (
	"math"
	"time"
)

// float64 representation of distance in meters.
type Distance float64

const (
	Nanometer           = 1e-9 * Meter
	Micrometer          = 1e-6 * Meter
	Millimeter          = 1e-3 * Meter
	Centimeter          = 1e-2 * Meter
	Decimeter           = 1e-1 * Meter
	Meter      Distance = 1
	Kilometer           = 1e3 * Meter
	Yard                = 0.9144 * Meter
)

func (d Distance) Nanometers() float64 {
	return float64(d / Nanometer)
}

func (d Distance) Micrometers() float64 {
	return float64(d / Micrometer)
}

func (d Distance) Millimeters() float64 {
	return float64(d / Millimeter)
}

func (d Distance) Centimeters() float64 {
	return float64(d / Centimeter)
}

func (d Distance) Meters() float64 {
	return float64(d / Meter)
}

func (d Distance) Kilometers() float64 {
	return float64(d / Kilometer)
}

func (d Distance) Yards() float64 {
	return float64(d / Yard)
}

/* Operations */

func (d Distance) Abs() Distance {
	if d < 0 {
		return -d
	}
	return d
}

func (d Distance) DivideWithDuration(t time.Duration) Velocity {
	return Velocity(float64(d) / float64(t) * float64(time.Second))
}

func Hypot(x, y Distance) Distance {
	return Distance(math.Hypot(float64(x), float64(y)))
}
