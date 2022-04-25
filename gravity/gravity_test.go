package gravity_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wafer-bw/gorbit/gravity"
	"golang.org/x/image/math/f64"
)

func TestForce(t *testing.T) {
	t.Run("succeed in calculating gravitational force along the x axis", func(t *testing.T) {
		p1, p2 := f64.Vec3{200, 0, 0}, f64.Vec3{0, 0, 0}
		m1, m2 := 2e+6, 8e+6
		f := gravity.Force(p1, p2, m1, m2)
		fmag := math.Sqrt(f[0]*f[0] + f[1]*f[1] + f[2]*f[2])
		require.Equal(t, "-0.02669", fmt.Sprintf("%.5f", f[0]))
		require.Equal(t, "-0.00000", fmt.Sprintf("%.5f", f[1]))
		require.Equal(t, "-0.00000", fmt.Sprintf("%.5f", f[2]))
		require.Equal(t, "0.02669", fmt.Sprintf("%.5f", fmag))
	})
	t.Run("succeed in calculating gravitational force along the y axis", func(t *testing.T) {
		p1, p2 := f64.Vec3{0, 200, 0}, f64.Vec3{0, 0, 0}
		m1, m2 := 2e+6, 8e+6
		f := gravity.Force(p1, p2, m1, m2)
		fmag := math.Sqrt(f[0]*f[0] + f[1]*f[1] + f[2]*f[2])
		require.Equal(t, "-0.00000", fmt.Sprintf("%.5f", f[0]))
		require.Equal(t, "-0.02669", fmt.Sprintf("%.5f", f[1]))
		require.Equal(t, "-0.00000", fmt.Sprintf("%.5f", f[2]))
		require.Equal(t, "0.02669", fmt.Sprintf("%.5f", fmag))
	})
	t.Run("succeed in calculating gravitational force along the z axis", func(t *testing.T) {
		p1, p2 := f64.Vec3{0, 0, 200}, f64.Vec3{0, 0, 0}
		m1, m2 := 2e+6, 8e+6
		f := gravity.Force(p1, p2, m1, m2)
		fmag := math.Sqrt(f[0]*f[0] + f[1]*f[1] + f[2]*f[2])
		require.Equal(t, "-0.00000", fmt.Sprintf("%.5f", f[0]))
		require.Equal(t, "-0.00000", fmt.Sprintf("%.5f", f[1]))
		require.Equal(t, "-0.02669", fmt.Sprintf("%.5f", f[2]))
		require.Equal(t, "0.02669", fmt.Sprintf("%.5f", fmag))
	})
	t.Run("succeed in calculating gravitational force along all axes", func(t *testing.T) {
		p1, p2 := f64.Vec3{200, 200, 200}, f64.Vec3{0, 0, 0}
		m1, m2 := 2e+6, 8e+6
		f := gravity.Force(p1, p2, m1, m2)
		fmag := math.Sqrt(f[0]*f[0] + f[1]*f[1] + f[2]*f[2])
		require.Equal(t, "-0.00514", fmt.Sprintf("%.5f", f[0]))
		require.Equal(t, "-0.00514", fmt.Sprintf("%.5f", f[1]))
		require.Equal(t, "-0.00514", fmt.Sprintf("%.5f", f[2]))
		require.Equal(t, "0.00890", fmt.Sprintf("%.5f", fmag))
	})
}

func TestOrbitalElements(t *testing.T) {
	t.Run("succeed in calculating orbital elements for the moon around earth in 3D", func(t *testing.T) {
		p1 := f64.Vec3{0, 0, 0}
		v1 := f64.Vec3{0, 0, 0}

		p2 := f64.Vec3{0, 405400000, gravity.Epsilon}
		v2 := f64.Vec3{1090, 0, gravity.Epsilon}

		r := f64.Vec3{p2[0] - p1[0], p2[1] - p1[1], p2[2] - p1[2]}
		v := f64.Vec3{v2[0] - v1[0], v2[1] - v1[1], v2[2] - v1[2]}

		m1, m2 := 5.972e24, 7.34767309e22

		expectedA := 502989447.71483934
		expectedE := 0.1940188768535872
		expectedI := 180.00000000000017
		expectedLAN := 90.00015405130662
		expectedW := 0.0000011540590637606703
		expectedM := 2.1421319455355324e-16

		expectedPeriod := 3529030.4503167123
		expectedPE := 405400000.0000001
		expectedAP := 600578895.4296786

		a, e, w, lan, i, m := gravity.OrbitalElements(r, v, m1, m2)
		require.Equal(t, fmt.Sprintf("%.5f", expectedA), fmt.Sprintf("%.5f", a))
		require.Equal(t, fmt.Sprintf("%.5f", expectedE), fmt.Sprintf("%.5f", e))
		require.Equal(t, fmt.Sprintf("%.3f", expectedW), fmt.Sprintf("%.3f", gravity.Degrees(w)))
		require.Equal(t, fmt.Sprintf("%.5f", expectedLAN), fmt.Sprintf("%.5f", gravity.Degrees(lan)))
		require.Equal(t, fmt.Sprintf("%.5f", expectedI), fmt.Sprintf("%.5f", gravity.Degrees(i)))
		require.Equal(t, fmt.Sprintf("%.5f", expectedM), fmt.Sprintf("%.5f", m))
		require.Equal(t, fmt.Sprintf("%.5f", expectedPeriod), fmt.Sprintf("%.5f", gravity.Period(a, m1, m2)))
		require.Equal(t, fmt.Sprintf("%.5f", expectedPE), fmt.Sprintf("%.5f", gravity.Periapsis(a, e)))
		require.Equal(t, fmt.Sprintf("%.5f", expectedAP), fmt.Sprintf("%.5f", gravity.Apoapsis(a, e)))
	})
	t.Run("succeed in calculating orbital elements for the moon around earth in 2D", func(t *testing.T) {
		p1 := f64.Vec3{0, 0, 0}
		v1 := f64.Vec3{0, 0, 0}

		p2 := f64.Vec3{0, 405400000, 0}
		v2 := f64.Vec3{1090, 0, 0}

		r := f64.Vec3{p2[0] - p1[0], p2[1] - p1[1], p2[2] - p1[2]}
		v := f64.Vec3{v2[0] - v1[0], v2[1] - v1[1], v2[2] - v1[2]}

		m1, m2 := 5.972e24, 7.34767309e22

		expectedA := 502989447.71483934
		expectedE := 0.1940188768535872
		expectedI := 180.00000000000017
		expectedLAN := float64(0)
		expectedW := float64(0)
		expectedM := 2.1421319455355324e-16

		expectedPeriod := 3529030.4503167123
		expectedPE := 405400000.0000001
		expectedAP := 600578895.4296786

		a, e, w, lan, i, m := gravity.OrbitalElements(r, v, m1, m2)
		require.Equal(t, fmt.Sprintf("%.5f", expectedA), fmt.Sprintf("%.5f", a))
		require.Equal(t, fmt.Sprintf("%.5f", expectedE), fmt.Sprintf("%.5f", e))
		require.Equal(t, fmt.Sprintf("%.3f", expectedW), fmt.Sprintf("%.3f", gravity.Degrees(w)))
		require.Equal(t, fmt.Sprintf("%.5f", expectedLAN), fmt.Sprintf("%.5f", gravity.Degrees(lan)))
		require.Equal(t, fmt.Sprintf("%.5f", expectedI), fmt.Sprintf("%.5f", gravity.Degrees(i)))
		require.Equal(t, fmt.Sprintf("%.5f", expectedM), fmt.Sprintf("%.5f", m))
		require.Equal(t, fmt.Sprintf("%.5f", expectedPeriod), fmt.Sprintf("%.5f", gravity.Period(a, m1, m2)))
		require.Equal(t, fmt.Sprintf("%.5f", expectedPE), fmt.Sprintf("%.5f", gravity.Periapsis(a, e)))
		require.Equal(t, fmt.Sprintf("%.5f", expectedAP), fmt.Sprintf("%.5f", gravity.Apoapsis(a, e)))
	})
}

func TestEccentricAnomoly(t *testing.T) {
	t.Run("high eccentricity", func(t *testing.T) {
		eca := gravity.EccentricAnomoly(0.99999, gravity.Radians(245))
		require.Equal(t, "3.7251043", fmt.Sprintf("%.7f", eca))
	})
	t.Run("low eccentricity", func(t *testing.T) {
		eca := gravity.EccentricAnomoly(0.00001, gravity.Radians(30))
		require.Equal(t, "0.5236038", fmt.Sprintf("%.7f", eca))
	})
}

func TestStateVectors(t *testing.T) {
	t.Run("succeed in calculating state vectors from orbital elements at epoch (t0)", func(t *testing.T) {
		a := 502989447.71483934
		e := 0.1940188768535872
		w := 0.0000011540590637606703
		lan := 90.00015405130662
		i := 180.00000000000017
		m := 2.1421319455355324e-16
		T := float64(0)
		m1, m2 := 5.972e24, 7.34767309e22

		expectedRX := -1081.834380268568
		expectedRY := 4.053999999985566e+08
		expectedRZ := float64(0)
		expectedVX := 1089.999999996119
		expectedVY := 0.002908730820160692
		expectedVZ := float64(0)

		r, v := gravity.StateVectors(
			a,
			e,
			gravity.Radians(w),
			gravity.Radians(lan),
			gravity.Radians(i),
			gravity.Radians(m),
			T,
			m1,
			m2,
		)
		require.Equal(t, fmt.Sprintf("%.6f", expectedRX), fmt.Sprintf("%.6f", r[0]))
		require.Equal(t, fmt.Sprintf("%.6f", expectedRY), fmt.Sprintf("%.6f", r[1]))
		require.Equal(t, fmt.Sprintf("%.6f", expectedRZ), fmt.Sprintf("%.6f", r[2]))
		require.Equal(t, fmt.Sprintf("%.6f", expectedVX), fmt.Sprintf("%.6f", v[0]))
		require.Equal(t, fmt.Sprintf("%.6f", expectedVY), fmt.Sprintf("%.6f", v[1]))
		require.Equal(t, fmt.Sprintf("%.6f", expectedVZ), fmt.Sprintf("%.6f", v[2]))
	})
	t.Run("succeed in calculating state vectors for the moon around earth in 2D at t+100", func(t *testing.T) {
		a := 502989447.71483934
		e := 0.1940188768535872
		w := 0.0000011540590637606703
		lan := 90.00015405130662
		i := 180.00000000000017
		m := 2.1421319455355324e-16
		m1, m2 := 5.972e24, 7.34767309e22
		T := gravity.Period(a, m1, m2) / 2

		expectedRX := 1602.6810503687375
		expectedRY := -6.005788954275403e+08
		expectedRZ := float64(0)
		expectedVX := -735.76677995367
		expectedVY := -0.0019634380843619485
		expectedVZ := float64(0)

		r, v := gravity.StateVectors(
			a,
			e,
			gravity.Radians(w),
			gravity.Radians(lan),
			gravity.Radians(i),
			gravity.Radians(m),
			T,
			m1,
			m2,
		)
		require.Equal(t, fmt.Sprintf("%.6f", expectedRX), fmt.Sprintf("%.6f", r[0]))
		require.Equal(t, fmt.Sprintf("%.6f", expectedRY), fmt.Sprintf("%.6f", r[1]))
		require.Equal(t, fmt.Sprintf("-%.6f", expectedRZ), fmt.Sprintf("%.6f", r[2]))
		require.Equal(t, fmt.Sprintf("%.6f", expectedVX), fmt.Sprintf("%.6f", v[0]))
		require.Equal(t, fmt.Sprintf("%.6f", expectedVY), fmt.Sprintf("%.6f", v[1]))
		require.Equal(t, fmt.Sprintf("-%.6f", expectedVZ), fmt.Sprintf("%.6f", v[2]))
	})
}

func TestSvToOeToSv(t *testing.T) {
	p1 := f64.Vec3{0, 0, 0}
	v1 := f64.Vec3{0, 0, 0}

	p2 := f64.Vec3{0, 405400000, 100}
	v2 := f64.Vec3{1090, 0, 10}

	r := f64.Vec3{p2[0] - p1[0], p2[1] - p1[1], p2[2] - p1[2]}
	v := f64.Vec3{v2[0] - v1[0], v2[1] - v1[1], v2[2] - v1[2]}

	m1, m2 := 5.972e24, 7.34767309e22

	a, e, w, lan, i, m := gravity.OrbitalElements(r, v, m1, m2)
	R, V := gravity.StateVectors(a, e, w, lan, i, m, 0, m1, m2)

	a, e, w, lan, i, m = gravity.OrbitalElements(R, V, m1, m2)
	R2, V2 := gravity.StateVectors(a, e, w, lan, i, m, 0, m1, m2)
	require.Equal(t, fmt.Sprintf("%.6f", R[0]), fmt.Sprintf("%.6f", R2[0]))
	require.Equal(t, fmt.Sprintf("%.6f", R[1]), fmt.Sprintf("%.6f", R2[1]))
	require.Equal(t, fmt.Sprintf("%.6f", R[2]), fmt.Sprintf("%.6f", R2[2]))
	require.Equal(t, fmt.Sprintf("%.6f", V[0]), fmt.Sprintf("%.6f", V2[0]))
	require.Equal(t, fmt.Sprintf("%.6f", R[1]), fmt.Sprintf("%.6f", R2[1]))
	require.Equal(t, fmt.Sprintf("%.6f", R[2]), fmt.Sprintf("%.6f", R2[2]))
}
