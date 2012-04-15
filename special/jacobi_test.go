package special

import "testing"

var jacobi = []float64{
	1.830168625887056e+06,
	2.6431748399463676e+07,
	0.6596293075491676,
	1.9019429125281873e+06,
	9.910913256270967e+07,
	69652.66946089824,
	2.4652203034949177e+06,
	44849.54643817527,
	3334.0292843868647,
	5.3022890790198795e+07,
}

func TestJacobi(t *testing.T) {
	for i := 0; i < len(vf); i++ {
		if f := Jacobi(6, 2, 2, vf[i]); !soclose(jacobi[i], f, 4e-14) {
			t.Errorf("Jacobi(6,2,2,%g) = %g, want %g", vf[i], f, jacobi[i])
		}
	}
}
