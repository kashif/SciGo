package special

import "testing"

func tolerance(a, b, e float64) bool {
	d := a - b
	if d < 0 {
		d = -d
	}

	if a != 0 {
		e = e * a
		if e < 0 {
			e = -e
		}
	}
	return d < e
}

func soclose(a, b, e float64) bool { return tolerance(a, b, e) }

var vf = []float64{
	4.9790119248836735e+00,
	7.7388724745781045e+00,
	-2.7688005719200159e-01,
	-5.0106036182710749e+00,
	9.6362937071984173e+00,
	2.9263772392439646e+00,
	5.2290834314593066e+00,
	2.7279399104360102e+00,
	1.8253080916808550e+00,
	-8.6859247685756013e+00,
}

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
