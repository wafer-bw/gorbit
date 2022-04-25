package gravity_test

import (
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
		gravity.EccentricAnomoly(rand.Float64(), gravity.Radians(float64(rand.Intn(360))))
	}
}
