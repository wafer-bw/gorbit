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
PASS
coverage: 92.4% of statements
ok  	github.com/wafer-bw/gorbit/gravity	12.646s
```

### MacOS
```
# TODO
```

## References
- http://orbitsimulator.com/formulas/OrbitalElements.html
- https://downloads.rene-schwarz.com/download/M002-Cartesian_State_Vectors_to_Keplerian_Orbit_Elements.pdf
- https://downloads.rene-schwarz.com/download/M001-Keplerian_Orbit_Elements_to_Cartesian_State_Vectors.pdf
- https://elainecoe.github.io/orbital-mechanics-calculator/formulas.html#two-body
- http://www.csun.edu/~hcmth017/master/node16.html

## Further Reading
- https://articles.adsabs.harvard.edu/full/seri/AJ.../0067//0000010.000.html
