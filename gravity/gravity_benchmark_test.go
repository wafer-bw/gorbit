package gravity_test

import (
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/wafer-bw/gorbit/gravity"
	"golang.org/x/image/math/f64"
)

func BenchmarkForce(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		gravity.Force(
			f64.Vec3{
				rand.NormFloat64(),
				rand.NormFloat64(),
				rand.NormFloat64(),
			},
			f64.Vec3{
				rand.NormFloat64(),
				rand.NormFloat64(),
				rand.NormFloat64(),
			},
			rand.NormFloat64(),
			rand.NormFloat64(),
		)
	}
}

func BenchmarkEccentricAnomaly(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		gravity.EccentricAnomaly(rand.Float64(), gravity.Radians(float64(rand.Intn(360))))
	}
}

func BenchmarkOrbitalElements(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		gravity.OrbitalElements(
			f64.Vec3{
				rand.NormFloat64(),
				rand.NormFloat64(),
				rand.NormFloat64(),
			},
			f64.Vec3{
				rand.NormFloat64(),
				rand.NormFloat64(),
				rand.NormFloat64(),
			},
			rand.NormFloat64(),
			rand.NormFloat64(),
		)
	}
}

func BenchmarkStateVectors(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		gravity.StateVectors(
			rand.NormFloat64(),
			rand.Float64(),
			gravity.Radians(float64(rand.Intn(360))),
			gravity.Radians(float64(rand.Intn(360))),
			gravity.Radians(float64(rand.Intn(360))),
			gravity.Radians(float64(rand.Intn(360))),
			math.Abs(rand.NormFloat64()),
			math.Abs(rand.NormFloat64()),
			math.Abs(rand.NormFloat64()),
		)
	}
}

func BenchmarkPeriapsis(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		gravity.Periapsis(rand.NormFloat64(), rand.Float64())
	}
}

func BenchmarkApoapsis(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		gravity.Apoapsis(rand.NormFloat64(), rand.Float64())
	}
}

func BenchmarkPeriod(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		gravity.Period(rand.NormFloat64(), math.Abs(rand.NormFloat64()), math.Abs(rand.NormFloat64()))
	}
}

func BenchmarkDegrees(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		gravity.Degrees(rand.NormFloat64())
	}
}

func BenchmarkRadians(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		gravity.Radians(rand.NormFloat64())
	}
}
