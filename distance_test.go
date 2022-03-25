package distance_test

import (
	//	"fmt"
	"gitlab.glaske.net/mglaske/distance"
	"math"
	"testing"
)

type DistanceTest struct {
	nanometers int64
	golden     parsedDistance
}

type parsedDistance struct {
	Nanometers     int64
	Micrometers    float64
	Millimeters    float64
	Centimeters    float64
	Meters         float64
	Kilometers     float64
	Thous          float64
	Inches         float64
	Feet           float64
	Yards          float64
	Miles          float64
	Fathoms        float64
	ImperialString string
	MetricString   string
}

var oneMile int64 = 5280 * 12 * 1000 * 25400 // Feet * inches * thousanths * nanometers

var dtests = []DistanceTest{
	{635000, parsedDistance{635000, 635, 0.635, 0.0635, 0.000635, 0, 25, 0.025, 0.00208333, 0, 0, 0, "25mil", "635μm"}},
	{oneMile * 1000, parsedDistance{oneMile * 1000, 1.609344e12, 1.609344e9, 1.609344e8, 1.609344e6, 1609.344, 6.336e10, 6.336e7, 5.28e6, 1.76e6, 1000, 0, "1000mi", "1609.34km"}}, // 1000 miles
	{oneMile, parsedDistance{oneMile, 1.609344e9, 1.609344e6, 160934.4, 1609.344, 1.609344, 6.336e7, 63360, 5280, 1760, 1, 0, "1mi", "1.61km"}},                                    // 1 mile
	{1e6, parsedDistance{1e6, 1000, 1, 0.1, 0.001, 1e-6, 39.370079, 0.39370079, 0.00328084, 0, 0, 0, "39.37mil", "1mm"}},                                                           // 1000 Micrometers
	{5e8, parsedDistance{5e8, 500000, 500, 50, 0.5, 0.0005, 19685.0393701, 19.685, 1.64042, 0.546807, 0.000310686, 0, "1ft7.69in", "50cm"}},                                        // 500 Millimeters
	{1.27e10, parsedDistance{1.27e10, 1.27e7, 12700, 1270, 12.7, 0.0127, 500000, 500, 41.666667, 13.888889, 0.00789141, 0, "41ft8in", "12m70cm"}},                                  // 500 inches
	{3.048e9, parsedDistance{3.048e9, 3.048e6, 3048, 304.8, 3.048, 0.003048, 120000, 120, 10, 3.3333333, 0.00189394, 0, "10ft0in", "3m4cm8mm"}},                                    // 10 Feet
}

const testTolerance = 0.000001

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= testTolerance
}

func TestDistances(t *testing.T) {
	var tv, gv float64
	var tvs, gvs string
	for _, test := range dtests {
		d := distance.Nanometer * distance.Distance(test.nanometers)
		tv = float64(d.Micrometers())
		gv = test.golden.Micrometers
		if gv != 0 && !almostEqual(tv, gv) {
			t.Errorf("Micrometers Test for nanometers=%d lib_value=%f != %f", test.nanometers, tv, gv)
		}
		tv = float64(d.Millimeters())
		gv = test.golden.Millimeters
		if gv != 0 && !almostEqual(tv, gv) {
			t.Errorf("Millimeters Test for nanometers=%d lib_value=%f != %f", test.nanometers, tv, gv)
		}
		tv = float64(d.Centimeters())
		gv = test.golden.Centimeters
		if gv != 0 && !almostEqual(tv, gv) {
			t.Errorf("Centimeters Test for nanometers=%d lib_value=%f != %f", test.nanometers, tv, gv)
		}
		tv = float64(d.Meters())
		gv = test.golden.Meters
		if gv != 0 && !almostEqual(tv, gv) {
			t.Errorf("Meters Test for nanometers=%d lib_value=%f != %f", test.nanometers, tv, gv)
		}
		tv = float64(d.Kilometers())
		gv = test.golden.Kilometers
		if gv != 0 && !almostEqual(tv, gv) {
			t.Errorf("Kilometers Test for nanometers=%d lib_value=%f != %f", test.nanometers, tv, gv)
		}
		tv = d.Thous()
		gv = test.golden.Thous
		if gv != 0 && !almostEqual(tv, gv) {
			t.Errorf("Thous Test for nanometers=%d lib_value=%f != %f", test.nanometers, tv, gv)
		}
		tv = d.Feet()
		gv = test.golden.Feet
		if gv != 0 && !almostEqual(tv, gv) {
			t.Errorf("Feet Test for nanometers=%d lib_value=%f != %f", test.nanometers, tv, gv)
		}
		tv = d.Yards()
		gv = test.golden.Yards
		if gv != 0 && !almostEqual(tv, gv) {
			t.Errorf("Yards Test for nanometers=%d lib_value=%f != %f", test.nanometers, tv, gv)
		}
		tv = d.Miles()
		gv = test.golden.Miles
		if gv != 0 && !almostEqual(tv, gv) {
			t.Errorf("Miles Test for nanometers=%d lib_value=%f != %f", test.nanometers, tv, gv)
		}
		distance.Imperial = true
		tvs = d.String()
		gvs = test.golden.ImperialString
		if tvs != gvs {
			t.Errorf("String test for nanometers=%d Imperial lib_value=%s != %s", test.nanometers, tvs, gvs)
		}
		distance.Imperial = false
		tvs = d.String()
		gvs = test.golden.MetricString
		if tvs != gvs {
			t.Errorf("String test for nanometers=%d Metric lib_value=%s != %s", test.nanometers, tvs, gvs)
			//t.Errorf("String test for nanometers=%d Metric lib_value=%s(%+v) != %s(%+v)", test.nanometers, tvs, []byte(tvs), gvs, []byte(gvs))
		}
	}
}
