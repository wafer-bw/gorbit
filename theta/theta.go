package theta

const Pi float64 = 3.14159265358979

// Degrees from radians.
func Degrees(rad float64) float64 {
	return 180 * rad / Pi
}

// Radians from degrees.
func Radians(deg float64) float64 {
	return Pi * deg / 180
}
