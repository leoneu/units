package units

import (
	"fmt"
	"math"
	"testing"
	"time"
)

// Used to check weather two floats are 'different'
const test_k = 0.0001

func cmpf64(f1, f2 float64) bool {
	err := f2 - f1
	if err < 0 {
		err = -err
	}
	if err < test_k {
		return true
	}
	fmt.Println(f1, f2)
	return false
}

func TestDistance(t *testing.T) {
	// Test some selected constants:
	if !cmpf64(float64(1e9*Nanometer), float64(1*Meter)) {
		t.Error("1e9 nm != 1 m")
	}
	if !cmpf64(float64(1000*Millimeter), float64(1*Meter)) {
		t.Error("1000 mm != 1 m")
	}
	if !cmpf64(float64(100*Centimeter), float64(1*Meter)) {
		t.Error("100 cm != 1 m")
	}
	if !cmpf64(float64(1000*Meter), float64(1*Kilometer)) {
		t.Error("10000 m != 1 km")
	}

	// Test f64 converter functions:
	var d Distance = 1 * Meter

	if !cmpf64(d.Kilometers(), 0.001) {
		t.Error("(1m).Kilometers() != 0.001")
	}
	if !cmpf64(d.Meters(), 1) {
		t.Error("(1m).Meters() != 1")
	}
	if !cmpf64(d.Centimeters(), 100) {
		t.Error("(1m).Centimeters() != 100")
	}
	if !cmpf64(d.Millimeters(), 1000) {
		t.Error("(1m).Millimeters() != 1000")
	}
	if !cmpf64(d.Micrometers(), 1e6) {
		t.Error("(1m).Micrometers() != 1e6")
	}
	if !cmpf64(d.Inches(), 39.3701) {
		t.Error("(1m).Inches() != 39.3701")
	}
}

func TestMass(t *testing.T) {
	if !cmpf64(float64(1000*Gram), float64(1*Kilogram)) {
		t.Error("1000 * g != 1 kg")
	}
	if !cmpf64(float64(1*Tonne), float64(1e6*Gram)) {
		t.Error("1 t != 1,000,000 g")
	}
	if !cmpf64(float64(1*Pound), float64(0.453592*Kilogram)) {
		t.Error("1 lb != 0.453592 kg")
	}
	if !cmpf64(float64(1*Pound), float64(16*Ounce)) {
		t.Error("1 lb != 16 oz")
	}

	// Test f64 converter functions:
	var m Mass = 1 * Kilogram

	if !cmpf64(m.Grams(), 1000) {
		t.Error("(1kg).Grams() != 1000")
	}
	if !cmpf64(m.Pounds(), 2.20462) {
		t.Error("(1kg).Pounds() != 2.20462")
	}
	if !cmpf64(m.Tonnes(), 0.001) {
		t.Error("(1kg).Tonnes() != 0.001")
	}
}

func TestVelocity(t *testing.T) {
	// TODO: Implement some tests
}

func TestAngles(t *testing.T) {
	a1 := 2 * math.Pi * Radian

	// Test all constants:
	a2 := 360 * Degree
	if !cmpf64(float64(a1), float64(a2)) {
		t.Error("2 Pi [rad] != 360 [deg]")
	}
	a2 = 400 * Gradian
	if !cmpf64(float64(a1), float64(a2)) {
		t.Error("2 Pi [rad] != 400 [grad]")
	}
	a2 = 2000 * math.Pi * Milliradian
	if !cmpf64(float64(a1), float64(a2)) {
		t.Error("2 Pi [rad] != 2000 Pi [mrad]")
	}

	// Test f64 functions:
	if !cmpf64(a1.Radians(), 2*math.Pi) {
		t.Error("(2 Pi [rad]).Radians() != 2 Pi")
	}
	if !cmpf64(a1.Degrees(), 360) {
		t.Error("(2 Pi [rad]).Degrees() != 360")
	}
	if !cmpf64(a1.Gradians(), 400) {
		t.Error("(2 Pi [rad]).Gradians() != 400")
	}
}

func TestAngularVelocity(t *testing.T) {
	// TODO: Implement some tests
}

func TestUnitOperations(t *testing.T) {
	// velocity = distance / time
	v1 := Meter.DivideWithDuration(time.Second)
	v2 := MeterPerSecond
	if !cmpf64(float64(v1), float64(v2)) {
		t.Error("velocity != distance / time")
	}
	// distance = velocity * time
	d1 := MeterPerSecond.MultiplyWithDuration(time.Second)
	d2 := Meter
	if !cmpf64(float64(d1), float64(d2)) {
		t.Error("distance != velocity * time")
	}

	// area = d1 * d2
	s1 := Meter.MultiplyWithDistance(Meter)
	s2 := SquareMeter
	if !cmpf64(float64(s1), float64(s2)) {
		t.Error("area != distance * distance")
	}
	d1 = Foot
	s1 = d1.MultiplyWithDistance(d1)
	s2 = SquareFoot
	if !cmpf64(float64(s1), float64(s2)) {
		t.Error("area != distance * distance")
	}

	// angular velocity = angle / time
	r1 := Radian.DivideWithDuration(time.Second)
	r2 := RadianPerSecond
	if !cmpf64(float64(r1), float64(r2)) {
		t.Error("angular velocity != angle / time")
	}

	// angle = angular velocity * time
	a1 := RadianPerSecond.MultiplyWithDuration(time.Second)
	a2 := Radian
	if !cmpf64(float64(a1), float64(a2)) {
		t.Error("angle != angular velocity * time")
	}

	f := 500 * Hertz
	p := f.Period()
	if !cmpf64(float64(f), 1.0/float64(p.Seconds())) {
		t.Error("period != 1 / frequency")
	}
}
