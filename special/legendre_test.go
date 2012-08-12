package special

import "testing"

var legendre = []float64{
	208026.681799869,
	3.0311795466441787e+06,
	0.08139663958056097,
	216226.67891243123,
	1.1390700088223575e+07,
	7679.284657380014,
	280611.42678106716,
	4908.0426079885165,
	336.97034551754865,
	6.088386439050518e+06,
}

func TestLegendre(t *testing.T) {
	for i := 0; i < len(vf); i++ {
		if f := Legendre(6, vf[i]); !soclose(legendre[i], f, 4e-14) {
			t.Errorf("Legendre(6,%g) = %g, want %g", vf[i], f, legendre[i])
		}
	}
}
