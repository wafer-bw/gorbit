# gorbit
Keplerian orbital mechanics helper package

![tests](https://github.com/wafer-bw/gorbit/workflows/tests/badge.svg)
![lint](https://github.com/wafer-bw/gorbit/workflows/lint/badge.svg)

## Benchmarks

### Windows
```
goos: windows
goarch: amd64
pkg: github.com/wafer-bw/gorbit/gravity
cpu: Intel(R) Core(TM) i9-9900K CPU @ 3.60GHz
BenchmarkForce-16               	 8622380	       140.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkEccentricAnomaly-16    	 7895198	       153.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkOrbitalElements-16     	 4305946	       280.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkStateVectors-16        	 1316767	       917.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkPeriapsis-16           	47008496	        25.18 ns/op	       0 B/op	       0 allocs/op
BenchmarkApoapsis-16            	46990088	        25.50 ns/op	       0 B/op	       0 allocs/op
BenchmarkPeriod-16              	30735972	        39.41 ns/op	       0 B/op	       0 allocs/op
BenchmarkDegrees-16             	92124859	        13.58 ns/op	       0 B/op	       0 allocs/op
BenchmarkRadians-16             	84480270	        13.57 ns/op	       0 B/op	       0 allocs/op
```

### MacOS
```
goos: darwin
goarch: amd64
pkg: github.com/wafer-bw/gorbit/gravity
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkForce-12                        6555768               170.1 ns/op             0 B/op          0 allocs/op
BenchmarkEccentricAnomaly-12             4374888               275.6 ns/op             0 B/op          0 allocs/op
BenchmarkOrbitalElements-12              2930158               455.9 ns/op             0 B/op          0 allocs/op
BenchmarkStateVectors-12                  865819              1386 ns/op               0 B/op          0 allocs/op
BenchmarkPeriapsis-12                   33178585                32.43 ns/op            0 B/op          0 allocs/op
BenchmarkApoapsis-12                    34882042                29.81 ns/op            0 B/op          0 allocs/op
BenchmarkPeriod-12                      21805924                48.01 ns/op            0 B/op          0 allocs/op
BenchmarkDegrees-12                     61156219                17.98 ns/op            0 B/op          0 allocs/op
BenchmarkRadians-12                     69422544                17.25 ns/op            0 B/op          0 allocs/op
```

## References
- http://orbitsimulator.com/formulas/OrbitalElements.html
- https://downloads.rene-schwarz.com/download/M002-Cartesian_State_Vectors_to_Keplerian_Orbit_Elements.pdf
- https://downloads.rene-schwarz.com/download/M001-Keplerian_Orbit_Elements_to_Cartesian_State_Vectors.pdf
- https://elainecoe.github.io/orbital-mechanics-calculator/formulas.html#two-body
- http://www.csun.edu/~hcmth017/master/node16.html

## Further Reading
- https://articles.adsabs.harvard.edu/full/seri/AJ.../0067//0000010.000.html
