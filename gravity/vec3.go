package gravity

import (
	"math"

	"golang.org/x/image/math/f64"
)

func sub(v1, v2 f64.Vec3) f64.Vec3 {
	v1[0] = v1[0] - v2[0]
	v1[1] = v1[1] - v2[1]
	v1[2] = v1[2] - v2[2]
	return v1
}

func cross(v1, v2 f64.Vec3) f64.Vec3 {
	return f64.Vec3{
		v1[1]*v2[2] - v1[2]*v2[1],
		v1[2]*v2[0] - v1[0]*v2[2],
		v1[0]*v2[1] - v1[1]*v2[0],
	}
}

func mulScalar(v f64.Vec3, s float64) f64.Vec3 {
	v[0] = v[0] * s
	v[1] = v[1] * s
	v[2] = v[2] * s
	return v
}

func divScalar(v f64.Vec3, s float64) f64.Vec3 {
	v[0] = v[0] / s
	v[1] = v[1] / s
	v[2] = v[2] / s
	return v
}

func magnitude(v f64.Vec3) float64 {
	return math.Sqrt(v[0]*v[0] + v[1]*v[1] + v[2]*v[2])
}

func dot(v1, v2 f64.Vec3) float64 {
	return v1[0]*v2[0] + v1[1]*v2[1] + v1[2]*v2[2]
}
