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

var velocity = map[string]Velocity{
	"nanometer_per_second":  NanometerPerSecond,
	"nm_per_s":              NanometerPerSecond,
	"micrometer_per_second": MicrometerPerSecond,
	"\u03BCm_per_s":         MicrometerPerSecond,
	"millimeter_per_second": MillimeterPerSecond,
	"mm_per_s":              MillimeterPerSecond,
	"centimeter_per_second": CentimeterPerSecond,
	"cm_per_s":              CentimeterPerSecond,
	"meter_per_second":      MeterPerSecond,
	"m_per_s":               MeterPerSecond,
	"kilometer_per_hour":    KilometerPerHour,
	"km_per_h":              KilometerPerHour,
}

// Creates a new velocity.
func NewVelocity(v float64, s string) Velocity {
	return Velocity(v) * velocity[s]
}

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

func (v Velocity) MultiplyWithDuration(t time.Duration) Length {
	return Length(float64(v) * float64(t) / float64(time.Second))
}
