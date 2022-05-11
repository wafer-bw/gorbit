package theta_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wafer-bw/gorbit/theta"
)

func TestRadians(t *testing.T) {
	r := theta.Radians(180)
	require.Equal(t, theta.Pi, r)
}

func TestDegrees(t *testing.T) {
	d := theta.Degrees(theta.Pi)
	require.Equal(t, float64(180), d)
}
