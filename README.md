# gorbit
2D & 3D orbital mechanics

![tests](https://github.com/wafer-bw/gorbit/workflows/tests/badge.svg)
![lint](https://github.com/wafer-bw/gorbit/workflows/lint/badge.svg)

## Benchmarks

### Windows
```
goos: windows
goarch: amd64
pkg: github.com/wafer-bw/gorbit/gravity
cpu: Intel(R) Core(TM) i9-9900K CPU @ 3.60GHz
BenchmarkForce-16                        8586116               139.7 ns/op             0 B/op          0 allocs/op
BenchmarkEccentricAnomaly-16             8142841               146.7 ns/op             0 B/op          0 allocs/op
```

## References
- http://orbitsimulator.com/formulas/OrbitalElements.html
- https://downloads.rene-schwarz.com/download/M002-Cartesian_State_Vectors_to_Keplerian_Orbit_Elements.pdf
- https://downloads.rene-schwarz.com/download/M001-Keplerian_Orbit_Elements_to_Cartesian_State_Vectors.pdf
- https://elainecoe.github.io/orbital-mechanics-calculator/formulas.html#two-body
- http://www.csun.edu/~hcmth017/master/node16.html
