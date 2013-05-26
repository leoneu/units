package units

import (
	"math"
	"time"
)

type AngularVelocity float64

const (
	MilliradianPerSecond AngularVelocity = 1
	RadianPerSecond                      = MilliradianPerSecond * 1e3
	DegreePerSecond                      = RadianPerSecond * 360.0 / (2.0 * math.Pi)
	GradianPerSecond                     = RadianPerSecond * 400.0 / (2.0 * math.Pi)
)

var angularVelocity = map[string]AngularVelocity{
	"milliradian_per_second": MilliradianPerSecond,
	"mrad_per_second":        MilliradianPerSecond,
	"radian_per_second":      RadianPerSecond,
	"rad_per_s":              RadianPerSecond,
	"degree_per_second":      DegreePerSecond,
	"gradian_per_second":     GradianPerSecond,
	"grad_per_second":        GradianPerSecond,
}

// Creates a new angular velocity.
func NewAngularVelocity(v float64, s string) AngularVelocity {
	return AngularVelocity(v) * angularVelocity[s]
}

func (v AngularVelocity) MilliradiansPerSecond() float64 {
	return float64(v) / float64(Milliradian)
}

func (v AngularVelocity) RadiansPerSecond() float64 {
	return float64(v) / float64(Radian)
}

func (v AngularVelocity) DegreesPerSecond() float64 {
	return float64(v) / float64(Degree)
}

func (v AngularVelocity) GradiansPerSecond() float64 {
	return float64(v) / float64(Gradian)
}

func (v AngularVelocity) MultiplyWithDuration(t time.Duration) Angle {
	return Angle(float64(v) * float64(t) / float64(time.Second))
}
