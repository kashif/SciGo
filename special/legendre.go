package special

/*
  Legendre polynomials: Legendre(n, x)
*/

// Special cases are:
//      Legendre(0, x) = 1.0
//      Legendre(1, x) = x
func Legendre(n int, x float64) float64 {
  //special cases
	switch {
	case n == 0:
		return 1.0
	case n == 1:
		return x
	}

	p0 := float64(1.0)
	p1 := x
	pn := p1

	for ii := 2; ii <= n; ii++ {
		i := float64(ii)
		pn = ((2.0*i-1.0)*x*p1 - (i-1.0)*p0) / i
		p0 = p1
		p1 = pn
	}

	return pn
}
