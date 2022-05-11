package gravity

import (
	"math"

	"github.com/wafer-bw/gorbit/vec3"
	"golang.org/x/image/math/f64"
)

const (
	G        float64 = 6.6725985e-11
	Pi       float64 = 3.14159265358979
	Epsilon  float64 = 1e-15 // very small number
	Epsilon6 float64 = 1e-06 // small number
)

// Force vector due to gravity (N) using Netwon's law of universal gravitation.
//
// p1: position of the primary body   (m),
// p2: position of the secondary body (m),
// m1: mass of the primary body       (kg),
// m2: mass of the secondary body     (kg).
//
// apply {-x, -y, -z} to b1,
// apply {x, y, z} to b2.
//
// https://en.wikipedia.org/wiki/Newton%27s_law_of_universal_gravitation#Vector_form
func Force(p1 f64.Vec3, p2 f64.Vec3, m1 float64, m2 float64) f64.Vec3 {
	r := vec3.Sub(p2, p1)
	d := vec3.Magnitude(r)
	rhat := vec3.DivScalar(r, d)
	return vec3.MulScalar(rhat, -(G*m1*m2)/(d*d))
}

// EccentricAnomaly (rad) using Newton's method.
//
// e: eccentricity (0-1),
// m: mean anomaly (rad).
//
// http://www.csun.edu/~hcmth017/master/node16.html
func EccentricAnomaly(e float64, m float64) float64 {
	eca := m + e/2
	if e >= 0.60 {
		eca = Pi
	}
	e1 := float64(0)
	diff := math.MaxFloat64
	for diff > Epsilon6 {
		e1 = eca - ((eca - e*math.Sin(eca) - m) / (1 - e*math.Cos(eca)))
		diff = math.Abs(e1 - eca)
		eca = e1
	}
	return eca
}

// OrbitalElements from Cartesian State Vectors.
//
// accepts:
// r:  position relative to primary body (m),
// v:  velocity relative to primary body (m/s),
// m1: mass of the primary body          (kg),
// m2: mass of the secondary body        (kg).
//
// returns:
// a:   semi-major axis                  (m),
// e:   eccentricity                     (0-1),
// w:   argument of periapsis            (rad),
// lan: longitude of ascending node      (rad),
// i:   inclination                      (rad),
// m:   mean anomaly                     (rad).
//
// The various singularities are removed by applying
// the following transformations where epsilon (eps)
// is equal to 1e-15. While the transformations are
// clearly not ideal it is the best I can do for now
// with my limited math knowledge.
//
// rZ=0  -> rZ=eps,
// vZ=0  -> vZ=eps,
// e=0   -> e=eps,
// i=0   -> i=eps,
// i=Pi  -> i=Pi-eps [180 (deg) = Pi (rad)].
//
// If the primary body is on-rails then set m2 to 0. If you don't
// set m2 to 0 then the elements will be predicted based on the
// bodies both orbiting their combined center of mass which is
// incorrect if the primary body is on-rails.
//
// https://downloads.rene-schwarz.com/download/M002-Cartesian_State_Vectors_to_Keplerian_Orbit_Elements.pdf
func OrbitalElements(r f64.Vec3, v f64.Vec3, m1 float64, m2 float64) (a, e, w, lan, i, m float64) {
	if r[2] == 0 {
		r[2] = Epsilon
	}

	if v[2] == 0 {
		v[2] = Epsilon
	}

	mu := G * (m1 + m2)

	rmag := vec3.Magnitude(r)
	vmag := vec3.Magnitude(v)

	h := vec3.Cross(r, v)

	evec := vec3.Sub(vec3.DivScalar(vec3.Cross(v, h), mu), vec3.DivScalar(r, rmag))

	e = vec3.Magnitude(evec)
	if e == 0 {
		e = Epsilon
	}

	n := f64.Vec3{-h[1], h[0], 0}
	nmag := vec3.Magnitude(n)

	ta := math.Acos(vec3.Dot(evec, r) / (e * rmag))
	if vec3.Dot(r, v) < 0 {
		ta = 2*Pi - ta
	}

	i = math.Acos(h[2] / vec3.Magnitude(h))
	if i == 0 {
		i = Epsilon
	} else if i == Pi {
		i -= Epsilon
	}

	eca := 2 * math.Atan(math.Tan(ta/2)/math.Sqrt((1+e)/(1-e)))

	lan = math.Acos(n[0] / nmag)
	if n[1] < 0 {
		lan = 2*Pi - lan
	}

	w = math.Acos(vec3.Dot(n, evec) / (nmag * e))
	if evec[2] < 0 {
		w = 2*Pi - w
	}

	m = eca - (e * math.Sin(eca))

	a = 1 / ((2 / rmag) - ((vmag * vmag) / mu))

	return
}

// StateVectors at t seconds after epoch from Keplerian Orbital Elements.
//
// accepts:
// a:   semi-major axis                  (m),
// e:   eccentricity                     (0-1),
// w:   argument of periapsis            (rad),
// lan: longitude of ascending node      (rad),
// i:   inclination                      (rad),
// m0:  mean anomaly at epoch            (rad),
// t:   time since epoch                 (seconds),
// m1:  mass of the primary body         (kg),
// m2:  mass of the secondary body       (kg).
//
// returns:
// r:  position relative to primary body (m),
// v:  velocity relative to primary body (m/s).
//
// If the primary body is on-rails then set m2 to 0.
// See OrbitalElements for more details.
//
// https://downloads.rene-schwarz.com/download/M001-Keplerian_Orbit_Elements_to_Cartesian_State_Vectors.pdf
func StateVectors(a, e, w, lan, i, m0, t, m1, m2 float64) (f64.Vec3, f64.Vec3) {
	mu := G * (m1 + m2)

	mT := float64(0)
	if t != 0 {
		mT = m0 + (t * math.Sqrt(mu/(a*a*a)))
	}
	ecaT := EccentricAnomaly(e, mT)
	taT := 2 * (math.Atan2(math.Sqrt(1+e)*math.Sin(ecaT/2), math.Sqrt(1-e)*math.Cos(ecaT/2)))
	rcT := a * (1 - e*math.Cos(ecaT))

	orT := f64.Vec3{
		math.Cos(taT),
		math.Sin(taT),
		0,
	}
	orT = vec3.MulScalar(orT, rcT)

	ovT := f64.Vec3{
		-math.Sin(ecaT),
		math.Sqrt(1-(e*e)) * math.Cos(ecaT),
		0,
	}
	ovT = vec3.MulScalar(ovT, math.Sqrt(mu*a)/rcT)

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
// e: eccentricity    (0-1).
//
// https://en.wikipedia.org/wiki/Apsis
func Periapsis(a, e float64) float64 {
	return a * (1 - e)
}

// Apoapsis distance (m).
//
// a: semi-major axis (m),
// e: eccentricity    (0-1).
//
// https://en.wikipedia.org/wiki/Apsis
func Apoapsis(a, e float64) float64 {
	return a * (1 + e)
}

// Period (s).
//
// a:  semi-major axis            (m),
// m1: mass of the primary body   (kg),
// m2: mass of the secondary body (kg).
//
// If the primary body is on-rails then set m2 to 0.
// See OrbitalElements for more details.
//
// https://en.wikipedia.org/wiki/Orbital_period
func Period(a, m1, m2 float64) float64 {
	mu := G * (m1 + m2)
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
