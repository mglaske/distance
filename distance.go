package distance

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
	Thou       Distance = 25400
	Mil                 = Thou
	Barleycorn          = 846670
	Inch                = Mil * 1000
	Hand                = Inch * 4
	Foot                = Inch * 12
	Yard                = Foot * 3
	Chain               = Yard * 22
	Furlong             = Chain * 10
	Mile                = Furlong * 8
	League              = Mile * 3
)

func (d Distance) String() string {
	var buf [32]byte
	w := len(buf)

	u := uint64(d)
	neg := d < 0
	if neg {
		u = -u
	}

	if Imperial {
		if u < uint64(Inch) {
			// Special case: if distance is smaller than a inch
			// use smaller units, like 1.2Mil
			var prec int
			w--
			switch {
			case u == 0:
				return "0in"
			default:
				// print mil
				prec = 6
				buf[w] = 'l'
				w--
				buf[w] = 'i'
				w--
				buf[w] = 'm'
				w--
			}
			w, u = fmtFrac(buf[:w], u, 10, prec)
			w = fmtInt(buf[:w], u)
		} else {
			w--
			buf[w] = 'n'
			w--
			buf[w] = 'i'

			w, u = fmtFrac(buf[:w], u, 12, 9)

			// u is now integer feet
			w = fmtInt(buf[:w], u%5280)
			u /= 5280

			// u is now integer miles
			if u > 0 {
				w--
				buf[w] = 'i'
				w--
				buf[w] = 'm'
				w = fmtInt(buf[:w], u)
			}
		}

	} else {
		if u < uint64(Centimeter) {
			// Special case: if distance is smaller than a centimeter,
			// use smaller units, like 1.2ns
			var prec int
			w--
			buf[w] = 'm' // suffix
			w--
			switch {
			case u == 0:
				return "0cm"
			case u < uint64(Micrometer):
				// print nanometers
				prec = 0
				buf[w] = 'n'
			case u < uint64(Millimeter):
				// print micrometer
				prec = 3
				// U+00B5 'µ' micro sign == 0xC2 0xB5
				w-- // Need room for two bytes.
				copy(buf[w:], "µ")
			default:
				// print millimeters
				prec = 6
				buf[w] = 'm'
			}
			w, u = fmtFrac(buf[:w], u, 10, prec)
			w = fmtInt(buf[:w], u)
		} else {
			w--
			buf[w] = 'm'

			w, u = fmtFrac(buf[:w], u, 10, 9)

			// u is now integer centimeters
			w = fmtInt(buf[:w], u%100)
			u /= 100

			// u is now integer meters
			if u > 0 {
				w--
				buf[w] = 'M'
				w = fmtInt(buf[:w], u%1000)
				u /= 1000

				// u is now integer kilometer
				// Stop at kilometers
				if u > 0 {
					w--
					buf[w] = 'k'
					w = fmtInt(buf[:w], u)
				}
			}
		}
	}

	if neg {
		w--
		buf[w] = '-'
	}

	return string(buf[w:])
}

func fmtFrac(buf []byte, v uint64, base, prec int) (nw int, nv uint64) {
	// Omit trailing zeros up to and including decimal point.
	w := len(buf)
	print := false
	for i := 0; i < prec; i++ {
		digit := v % base
		print = print || digit != 0
		if print {
			w--
			buf[w] = byte(digit) + '0'
		}
		v /= base
	}
	if print {
		w--
		buf[w] = '.'
	}
	return w, v
}

// fmtInt formats v into the tail of buf.
// It returns the index where the output begins.
func fmtInt(buf []byte, v uint64) int {
	w := len(buf)
	if v == 0 {
		w--
		buf[w] = '0'
	} else {
		for v > 0 {
			w--
			buf[w] = byte(v%10) + '0'
			v /= 10
		}
	}
	return w
}

func (d Distance) Nanometers() int64 { return int64(d) }

func (d Distance) Micrometers() int64 { return int64(d) / 1e3 }

func (d Distance) Millimeters() int64 { return int64(d) / 1e6 }

func (d Distance) Centimeters() float64 {
	cm := d / Centimeter
	mm := d % Centimeter
	return float64(cm) + float64(mm)/10
}

func (d Distance) Decimeters() float64 {
	dm := d / Decimeter
	mm := d % Decimeter
	return float64(dm) + float64(mm)/100
}

func (d Distance) Meters() float64 {
	M := d / Meter
	mm := d % Meter
	return float64(M) + float64(mm)/1000
}

func (d Distance) Dekameters() float64 {
	Dm := d / Dekameter
	M := d % Dekameter
	return float64(Dm) + float64(M)/10
}

func (d Distance) Hectometers() float64 {
	Hm := d / Hectometer
	M := d % Hectometer
	return float64(Hm) + float64(M)/100
}

func (d Distance) Kilometers() float64 {
	km := d / Kilometer
	M := d % Kilometer
	return float64(km) + float64(M)/1000
}

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