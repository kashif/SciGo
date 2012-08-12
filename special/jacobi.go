package special

/*
  Jacobi polynomials: Jacobi(n, a, b, x)
*/

// Special cases are:
//      Jacobi(0, a, b, x) = 1.0
//      Jacobi(1, a, b, x) = 0.5 * (a - b + (a + b + 2.0)*x)
func Jacobi(n int, a float64, b float64, x float64) float64 {
	// special cases
	switch {
	case n == 0:
		return 1.0
	case n == 1:
		return 0.5 * (a - b + (a+b+2.0)*x)
	}

	p0 := float64(1.0)
	p1 := float64(0.5 * (a - b + (a+b+2.0)*x))
	p2 := float64(0.0)

	var a1, a2, a3, a4 float64

	for ii := 1; ii < n; ii++ {
		i := float64(ii)
		a1 = 2.0 * (i + 1.0) * (i + a + b + 1.0) * (2.0*i + a + b)
		a2 = (2.0*i + a + b + 1.0) * (a*a - b*b)
		a3 = (2.0*i + a + b) * (2.0*i + a + b + 1.0) * (2.0*i + a + b + 2.0)
		a4 = 2.0 * (i + a) * (i + b) * (2.0*i + a + b + 2.0)
		p2 = 1.0 / a1 * ((a2+a3*x)*p1 - a4*p0)

		p0 = p1
		p1 = p2
	}

	return p2
}
