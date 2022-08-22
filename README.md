Distance Module
===============

Written on the same idea as the golang time module.  Set a distance of any
measurement type, and retrieve it as any other type, including conversions
between imperial and metric values.  Also output a nice string representation
of your distance for printing.

NOTE: take this for what it's worth, it works for my purposes.  And yes, 
I know this is somewhat inaccurate due to A) conversions aren't 100%, and
B) Floats are not very accurate either.

# Usage
Also, see the example directory or the `distance_test.go`
```
x := distance.Mile * 5
x = x + (500 * distance.Feet)

distance.Imperial = true
fmt.Printf("Distance: %s\n", x)
5.09mi

fmt.Printf("Distance in Feet: %f\n", x.Feet())
26900.000000

distance.Imperial = false
fmt.Printf("Distance: %s\n", x)
8.20km
```

# Future
Possibly todo, would be to store the values as they are, then allow direct
conversions from them, instead of changing everything to nanometers.
