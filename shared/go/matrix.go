package shared

type Matrix interface {
	Dims() (int, int)
	At(i, j int) MyFloat
	T() Matrix // Transpose
}

type Transpose struct {
	Matrix Matrix
}

func (t Transpose) At(i, j int) MyFloat {
	return t.Matrix.At(j, i)
}

func (t Transpose) Dims() (r, c int) {
	c, r = t.Matrix.Dims()
	return r, c
}

func (t Transpose) T() Matrix {
	return t.Matrix
}

// Dense is a matrix type that is backed by a slice of data.
type Dense struct {
	Rows, Cols int
	Data       []MyFloat
}

func (d *Dense) Dims() (int, int) {
	return d.Rows, d.Cols
}

func (d *Dense) At(i, j int) MyFloat {
	return d.Data[i*d.Cols+j]
}

func (d *Dense) T() Matrix {
	return &Transpose{d}
}

func NewDense(r, c int) *Dense {
	return &Dense{r, c, make([]MyFloat, r*c)}
}

func (d *Dense) Set(i, j int, v MyFloat) {
	d.Data[i*d.Cols+j] = v
}

func (d *Dense) Dot(a, b Matrix) Matrix {
	br, bc := b.Dims()
	ar, ac := a.Dims()
	if ac != br {
		panic("matices: dimension mismatch")
	}
	for i := 0; i < ar; i++ {
		for j := 0; j < bc; j++ {
			var v MyFloat
			for k := 0; k < ac; k++ {
				v += a.At(i, k) * b.At(k, j)
			}
			d.Set(i, j, v)
		}
	}
	return d
}

func Eq(a, b Matrix) bool {
	ar, ac := a.Dims()
	br, bc := b.Dims()
	if ar != br || ac != bc {
		return false
	}
	for i := 0; i < ar; i++ {
		for j := 0; j < ac; j++ {
			if a.At(i, j) != b.At(i, j) {
				return false
			}
		}
	}
	return true
}

func (d *Dense) Eq(a Matrix) bool {
	return Eq(d, a)
}
