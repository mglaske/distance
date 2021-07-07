package distance_test

import (
	"distance"
	"testing"
)

type DistanceTest struct {
	nanometers int64
	golden     parsedDistance
}

type parsedDistance struct {
	nanometers  float64
	micrometers float64
	millimeters float64
	centimeters float64
	meters      float64
	kilometers  float64
	thous       float64
	inches      float64
	feet        float64
	yards       float64
	miles       float64
	fathoms     float64
}

var dtests = []DistanceTest{
	{635000, parsedDistance{635000, 635, 0.635, 0.0635, 0.000635, 0, 25, 0.025, 0.00208333, 0, 0, 0}},
	{1.609e15, parsedDistance{1.609e15, 1.609e12, 1.609e9, 1.609e8, 1.609e6, 1609.34, 6.336e10, 6.336e7, 5.28e6, 1.76e6, 1000, 0}}, // 1000 miles
	{1.609e12, parsedDistance{1.609e12, 1.609e9, 1.609e6, 160934, 1609.34, 1.60934, 6.336e7, 63360, 5280, 1760, 1, 0}},             // 1 mile
	{1e6, parsedDistance{1e6, 1000, 1, 0.001, 0.1, 1e-6, 39.3701, 0.393701, 0.00328084, 0, 0, 0}},                                  // 1000 Micrometers
	{5e8, parsedDistance{5e8, 500000, 500, 50, 0.5, 0.0005, 19685, 19.685, 1.64042, 0.546807, 0.000310686, 0}},                     // 500 Millimeters
	{1.27e10, parsedDistance{1.27e10, 1.27e7, 12700, 1270, 12.7, 0.0127, 500000, 500, 41.6667, 13.8889, 0.00789141, 0}},            // 500 inches
	{3.048e9, parsedDistance{3.048e9, 3.048e6, 3048, 304.8, 3.048, 0.003048, 120000, 120, 10, 3.33333, 00.189394, 0}},              // 10 Feet
}

func TestDistances(t *testing.T) {
	for _, test := range dtests {
		d := distance.Distance{test.nanometers}
		tv := d.Micrometers()
		gv := test.golden.micrometers
		if gv != 0 && tv != gv {
			t.Errorf("Test for nanometers=%d test_value=%f != %f", tv, gv)
		}
		tv = d.Millimeters()
		gv = test.golden.millimeters
		if gv != 0 && tv != gv {
			t.Errorf("Test for nanometers=%d test_value=%f != %f", tv, gv)
		}
		tv = d.Centimeters()
		gv = test.golden.centimeters
		if gv != 0 && tv != gv {
			t.Errorf("Test for nanometers=%d test_value=%f != %f", tv, gv)
		}
		tv = d.Meters()
		gv = test.golden.meters
		if gv != 0 && tv != gv {
			t.Errorf("Test for nanometers=%d test_value=%f != %f", tv, gv)
		}
		tv = d.Kilometers()
		gv = test.golden.kilometers
		if gv != 0 && tv != gv {
			t.Errorf("Test for nanometers=%d test_value=%f != %f", tv, gv)
		}
		tv = d.Thous()
		gv = test.golden.thous
		if gv != 0 && tv != gv {
			t.Errorf("Test for nanometers=%d test_value=%f != %f", tv, gv)
		}
		tv = d.Feet()
		gv = test.golden.feet
		if gv != 0 && tv != gv {
			t.Errorf("Test for nanometers=%d test_value=%f != %f", tv, gv)
		}
		tv = d.Yards()
		gv = test.golden.yards
		if gv != 0 && tv != gv {
			t.Errorf("Test for nanometers=%d test_value=%f != %f", tv, gv)
		}
		tv = d.Miles()
		gv = test.golden.miles
		if gv != 0 && tv != gv {
			t.Errorf("Test for nanometers=%d test_value=%f != %f", tv, gv)
		}
	}
}
