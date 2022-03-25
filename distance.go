package distance

import (
	"encoding/json"
	"fmt"
	"math"
)

type Distance int64

const (
	minDistance Distance = -1 << 63
	maxDistance Distance = 1<<63 - 1
)

const (
	Nanometer  Distance = 1
	Micrometer          = 1000 * Nanometer
	Micron              = Micrometer
	Millimeter          = 1000 * Micrometer
	Centimeter          = 10 * Millimeter
	Decimeter           = 100 * Millimeter
	Meter               = 1000 * Millimeter
	Dekameter           = 10 * Meter
	Hectometer          = 100 * Meter
	Kilometer           = 1000 * Meter
	Megameter           = 1000 * Kilometer
	Gigameter           = 1000 * Megameter

//	Terameter           = 1000 * Gigameter
)

// Change output of String
var Imperial bool = false

const (
	Thou         Distance = 25400
	Mil                   = Thou
	Barleycorn            = 846670
	Inch                  = Mil * 1000
	Hand                  = Inch * 4
	Foot                  = Inch * 12
	Yard                  = Foot * 3
	Chain                 = Foot * 66
	Furlong               = Foot * 660
	Mile                  = Foot * 5280
	League                = Foot * 15840
	Fathom                = Thou * 72000
	Cable                 = Fathom * 100
	NauticalMile          = Cable * 10
	Link                  = Thou * 7920
	Rod                   = Thou * 198000
)

func (d Distance) String() string {
	var val float64
	var out string
	//neg := d < 0

	if Imperial {
		switch {
		case d < Inch:
			val = float64(d) / float64(Mil)
			return fmt.Sprintf("%smil", fmtWholeOrFrac(val))
		case d >= Inch && d < Mile:
			ft := d / Foot
			in := float64(d%Foot) / float64(Inch)
			return fmt.Sprintf("%dft%sin", ft, fmtWholeOrFrac(in))
		case d >= Mile:
			val = float64(d) / float64(Mile)
			return fmt.Sprintf("%smi", fmtWholeOrFrac(val))
		}
	} else {
		out = ""
		switch {
		case d >= Kilometer:
			val = float64(d) / float64(Kilometer)
			return fmt.Sprintf("%skm", fmtWholeOrFrac(val))
		case d < Kilometer && d >= Meter:
			// Show meters
			M := d / Meter
			if M > 0 {
				out += fmt.Sprintf("%dm", M)
			}
			d -= M * Meter
			fallthrough
		case d < Meter && d >= Centimeter:
			cm := d / Centimeter
			if cm > 0 {
				out += fmt.Sprintf("%dcm", cm)
			}
			d -= cm * Centimeter
			fallthrough
		case d < Centimeter && d >= Millimeter:
			mm := d / Millimeter
			if mm > 0 {
				out += fmt.Sprintf("%dmm", mm)
			}
			d -= mm * Millimeter
			fallthrough
		case d < Millimeter && d >= Micrometer:
			um := d / Micrometer
			if um > 0 {
				out += fmt.Sprintf("%dμm", um)
			}
		}
	}
	return out
}

func fmtWholeOrFrac(v float64) string {
	w, f := math.Modf(v)
	if f == 0 {
		// Whole Number
		return fmt.Sprintf("%.0f", w)
	}
	return fmt.Sprintf("%.2f", v)
}

func (d Distance) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

func (d Distance) Nanometers() int64 { return int64(d) }

func (d Distance) Micrometers() float64 { return float64(d) / float64(Micrometer) }

func (d Distance) Millimeters() float64 { return float64(d) / float64(Millimeter) }

func (d Distance) Centimeters() float64 { return float64(d) / float64(Centimeter) }

func (d Distance) Decimeters() float64 { return float64(d) / float64(Decimeter) }

func (d Distance) Meters() float64 { return float64(d) / float64(Meter) }

func (d Distance) Dekameters() float64 { return float64(d) / float64(Dekameter) }

func (d Distance) Hectometers() float64 { return float64(d) / float64(Hectometer) }

func (d Distance) Kilometers() float64 { return float64(d) / float64(Kilometer) }

func (d Distance) Thous() float64 { return float64(d) / float64(Thou) }

func (d Distance) Mils() float64 { return d.Thous() }

func (d Distance) Barleycorns() float64 { return float64(d) / float64(Barleycorn) }

func (d Distance) Inches() float64 { return float64(d) / float64(Inch) }

func (d Distance) Feet() float64 { return float64(d) / float64(Foot) }

func (d Distance) Yards() float64 { return float64(d) / float64(Yard) }

func (d Distance) Furlongs() float64 { return float64(d) / float64(Furlong) }

func (d Distance) Miles() float64 { return float64(d) / float64(Mile) }

func (d Distance) Fathoms() float64 { return float64(d) / float64(Fathom) }

func (d Distance) Cables() float64 { return float64(d) / float64(Cable) }

func (d Distance) NauticalMiles() float64 { return float64(d) / float64(NauticalMile) }

func (d Distance) Links() float64 { return float64(d) / float64(Link) }

func (d Distance) Rods() float64 { return float64(d) / float64(Rod) }

// Truncate returns the result of rounding d toward zero to a multiple of m.
// If m <= 0, Truncate returns d unchanged.
func (d Distance) Truncate(m Distance) Distance {
	if m <= 0 {
		return d
	}
	return d - d%m
}

// lessThanHalf reports whether x+x < y but avoids overflow,
// assuming x and y are both positive (Distance is signed).
func lessThanHalf(x, y Distance) bool {
	return uint64(x)+uint64(x) < uint64(y)
}

// Round returns the result of rounding d to the nearest multiple of m.
// The rounding behavior for halfway values is to round away from zero.
// If the result exceeds the maximum (or minimum)
// value that can be stored in a Distance,
// Round returns the maximum (or minimum) duration.
// If m <= 0, Round returns d unchanged.
func (d Distance) Round(m Distance) Distance {
	if m <= 0 {
		return d
	}
	r := d % m
	if d < 0 {
		r = -r
		if lessThanHalf(r, m) {
			return d + r
		}
		if d1 := d - m + r; d1 < d {
			return d1
		}
		return minDistance // overflow
	}
	if lessThanHalf(r, m) {
		return d - r
	}
	if d1 := d + m - r; d1 > d {
		return d1
	}
	return maxDistance // overflow
}
