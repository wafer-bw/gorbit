package gravity

import (
	"math"

	"golang.org/x/image/math/f64"
)

const (
	G        float64 = 6.6725985e-11
	Pi       float64 = 3.14159265358979
	Epsilon  float64 = 1e-15 // very small number
	Epsilon5 float64 = 1e-05 // small number
)

// Force vector due to gravity (N) using Netwon's law of universal gravitation.
//
// p1: position of the primary body   (m),
// p2: position of the secondary body (m),
// m1: mass of the primary body       (kg),
// m2: mass of the secondary body     (kg)
//
// https://en.wikipedia.org/wiki/Newton%27s_law_of_universal_gravitation#Vector_form
func Force(p1 f64.Vec3, p2 f64.Vec3, m1 float64, m2 float64) f64.Vec3 {
	r := sub(p1, p2)
	d := magnitude(r)
	rhat := divScalar(r, d)
	return mulScalar(rhat, -(G*m1*m2)/(d*d))
}

// EccentricAnomoly (rad) using Newton's method.
//
// e: eccentricity (0-1),
// m: mean anomaly (rad)
//
// http://www.csun.edu/~hcmth017/master/node16.html
func EccentricAnomoly(e float64, m float64) float64 {
	eca := m + e/2
	if m >= 0.99 {
		eca = Pi
	}
	e1 := float64(0)
	diff := math.MaxFloat64
	for diff > Epsilon5 {
		e1 = eca - (eca-e*math.Sin(eca)-m)/(1-e*math.Cos(eca))
		diff = math.Abs(e1 - eca)
		eca = e1
	}
	return eca
}

// OrbitalElements from state vectors.
//
// accepts
// r:  position relative to primary body (m),
// v:  velocity relative to primary body (m/s),
// m1: mass of the primary body          (kg),
// m2: mass of the secondary body        (kg)
//
// returns
// a:   semi-major axis                  (m),
// e:   eccentricity                     (0-1),
// w:   argument of periapsis            (rad),
// lan: longitude of ascending node      (rad),
// i:   inclination                      (rad),
// m:   mean anomaly                     (rad),
// t:   time since epoch                 (seconds)
//
// If the primary body is on-rails then set m2 to 0. If you don't
// set m2 to 0 then the elements will be predicted based on the
// bodies both orbiting the center of mass between the two.
//
// https://downloads.rene-schwarz.com/download/M002-Cartesian_State_Vectors_to_Keplerian_Orbit_Elements.pdf
func OrbitalElements(r f64.Vec3, v f64.Vec3, m1 float64, m2 float64) (a, e, w, lan, i, m float64) {
	mu := G * (m1 + m2)

	rmag := magnitude(r)
	vmag := magnitude(v)

	h := cross(r, v)
	evec := sub(divScalar(cross(v, h), mu), divScalar(r, rmag))
	e = magnitude(evec)
	n := f64.Vec3{-h[1], h[0], 0}
	nmag := magnitude(n)

	rvdot := dot(r, v)
	ta := math.Acos(dot(evec, r) / (e * rmag))
	if rvdot < 0 {
		ta = 2*Pi - ta
	}

	i = math.Acos(h[2] / magnitude(h))
	eca := 2 * math.Atan(math.Tan(ta/2)/math.Sqrt((1+e)/(1-e)))

	lan = math.Acos(n[0] / nmag)
	if n[1] < 0 {
		lan = 2*Pi - lan
	}
	if math.IsNaN(lan) {
		lan = 0
	}

	w = math.Acos(dot(n, evec) / (nmag * e))
	if evec[2] < 0 {
		w = 2*Pi - w
	}
	if math.IsNaN(w) {
		w = 0
	}

	m = eca - (e * math.Sin(eca))

	a = 1 / ((2 / rmag) - ((vmag * vmag) / mu))

	return
}

// StateVectors at t seconds after epoch from orbital elements.
//
// accepts
// a:   semi-major axis                  (m),
// e:   eccentricity                     (0-1),
// w:   argument of periapsis            (rad),
// lan: longitude of ascending node      (rad),
// i:   inclination                      (rad),
// m0:  mean anomaly at epoch            (rad),
// t:   time since epoch                 (seconds),
// m1:  mass of the primary body         (kg),
// m2:  mass of the secondary body       (kg)
//
// returns
// r:  position relative to primary body (m),
// v:  velocity relative to primary body (m/s)
//
// If the primary body is on-rails then set m2 to 0. If you don't
// set m2 to 0 then the elements will be predicted based on the
// bodies both orbiting the center of mass between the two.
//
// https://downloads.rene-schwarz.com/download/M001-Keplerian_Orbit_Elements_to_Cartesian_State_Vectors.pdf
func StateVectors(a, e, w, lan, i, m0, t, m1, m2 float64) (f64.Vec3, f64.Vec3) {
	mu := G * (m1 + m2)

	mT := float64(0)
	if t != 0 {
		mT = m0 + (t * math.Sqrt(mu/(a*a*a)))
	}
	ecaT := EccentricAnomoly(e, mT)
	taT := 2 * (math.Atan2(math.Sqrt(1+e)*math.Sin(ecaT/2), math.Sqrt(1-e)*math.Cos(ecaT/2)))
	rcT := a * (1 - e*math.Cos(ecaT))

	orT := f64.Vec3{
		math.Cos(taT),
		math.Sin(taT),
		0,
	}
	orT = mulScalar(orT, rcT)

	ovT := f64.Vec3{
		-math.Sin(ecaT),
		math.Sqrt(1-(e*e)) * math.Cos(ecaT),
		0,
	}
	ovT = mulScalar(ovT, math.Sqrt(mu*a)/rcT)

	rT := f64.Vec3{
		orT[0]*(math.Cos(w)*math.Cos(lan)-math.Sin(w)*math.Cos(i)*math.Sin(lan)) - orT[1]*(math.Sin(w)*math.Cos(lan)+math.Cos(w)*math.Cos(i)*math.Sin(lan)),
		orT[0]*(math.Cos(w)*math.Sin(lan)+math.Sin(w)*math.Cos(i)*math.Cos(lan)) + orT[1]*(math.Cos(w)*math.Cos(i)*math.Cos(lan)-math.Sin(w)*math.Sin(lan)),
		orT[0]*(math.Sin(w)*math.Sin(i)) + orT[1]*(math.Cos(w)*math.Sin(i)),
	}

	vT := f64.Vec3{
		ovT[0]*(math.Cos(w)*math.Cos(lan)-math.Sin(w)*math.Cos(i)*math.Sin(lan)) - ovT[1]*(math.Sin(w)*math.Cos(lan)+math.Cos(w)*math.Cos(i)*math.Sin(lan)),
		ovT[0]*(math.Cos(w)*math.Sin(lan)+math.Sin(w)*math.Cos(i)*math.Cos(lan)) + ovT[1]*(math.Cos(w)*math.Cos(i)*math.Cos(lan)-math.Sin(w)*math.Sin(lan)),
		ovT[0]*(math.Sin(w)*math.Sin(i)) + ovT[1]*(math.Cos(w)*math.Sin(i)),
	}

	return rT, vT
}

// Periapsis distance (m).
//
// a: semi-major axis (m),
// e: eccentricity    (0-1)
//
// https://en.wikipedia.org/wiki/Apsis
func Periapsis(a, e float64) float64 {
	return a * (1 - e)
}

// Apoapsis distance (m).
//
// a: semi-major axis (m),
// e: eccentricity    (0-1)
//
// https://en.wikipedia.org/wiki/Apsis
func Apoapsis(a, e float64) float64 {
	return a * (1 + e)
}

// Period (s).
//
// a:  semi-major axis   (m),
// mu: G*(m1+m2) or G*m1 (kg)
//
// https://en.wikipedia.org/wiki/Orbital_period
func Period(a, mu float64) float64 {
	return (2 * Pi) * math.Sqrt((a*a*a)/mu)
}

// Degrees from radians.
func Degrees(rad float64) float64 {
	return 180 * rad / Pi
}

// Radians from degrees.
func Radians(deg float64) float64 {
	return Pi * deg / 180
}
