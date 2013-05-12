package units

import (
	"time"
)

// float64 representation of velocity in mm/sec
type Velocity float64

const (
	NanometerPerSecond           = 1e-9 * MeterPerSecond
	MicrometerPerSecond          = 1e-6 * MeterPerSecond
	MillimeterPerSecond          = 1e-3 * MeterPerSecond
	CentimeterPerSecond          = 1e-2 * MeterPerSecond
	MeterPerSecond      Velocity = 1
	KilometerPerHour             = 1e3 * MeterPerSecond / 60 / 60
)

func (v Velocity) MillimetersPerSecond() float64 {
	return float64(v / MillimeterPerSecond)
}

func (v Velocity) CentimetersPerSecond() float64 {
	return float64(v / CentimeterPerSecond)
}

func (v Velocity) MetersPerSecond() float64 {
	return float64(v / MeterPerSecond)
}

func (v Velocity) KilometersPerSecond() float64 {
	return float64(v / KilometerPerHour)
}

func (v Velocity) MultiplyWithDuration(t time.Duration) Distance {
	return Distance(float64(v) * float64(t) / float64(time.Second))
}
