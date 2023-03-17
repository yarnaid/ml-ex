package shared

type Vector interface {
	Matrix
	Len() int
	AtVec(i int) MyFloat
}

type TransposeVec struct {
	Vector Vector
}

func (t TransposeVec) At(i, j int) MyFloat {
	return t.Vector.At(j, i)
}

func (t TransposeVec) AtVec(i int) MyFloat {
	return t.Vector.AtVec(i)
}

func (t TransposeVec) Len() int {
	return t.Vector.Len()
}

func (t TransposeVec) Dims() (r, c int) {
	c, r = t.Vector.Dims()
	return r, c
}

func (t TransposeVec) T() Matrix {
	return t.Vector
}

func (t TransposeVec) TVec() Vector {
	return t.Vector
}

type DenseVec struct {
	len  int
	data []MyFloat
}

func NewDenseVec(l int) *DenseVec {
	if l <= 0 {
		panic("vector: non-positive length")
	}
	return &DenseVec{l, make([]MyFloat, l)}
}

func (d *DenseVec) Len() int {
	return d.len
}

func (d *DenseVec) AtVec(i int) MyFloat {
	return d.data[i]
}

func (d *DenseVec) At(i, j int) MyFloat {
	if i != 0 {
		panic("vector: row index out of range")
	}
	return d.data[j]
}

func (d *DenseVec) Dims() (int, int) {
	return d.len, 1
}

func (d *DenseVec) T() Matrix {
	return &Transpose{d}
}

func (d *DenseVec) TVec() Vector {
	return &TransposeVec{d}
}

func (d *DenseVec) SetVec(i int, v MyFloat) {
	d.data[i] = v
}
