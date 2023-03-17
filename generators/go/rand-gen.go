package generators

import (
	"math/rand"
	"shared"
)

type LinRegGenerator struct {
	Number int
	Dim    int
	W      shared.Vector
	B      shared.MyFloat
	Sigma  shared.MyFloat
	Area   shared.NDArea
}

func (l *LinRegGenerator) Generate() []*shared.DenseVec {
	// generate returns y = w*x + b + N(0, sigma)
	res := make([]*shared.DenseVec, l.N umber)
	lengths := make([]shared.MyFloat, l.Dim)
	for i := 0; i < l.Dim; i++ {
		lengths[i] = l.Area.Ranges[i].Length()
	}
	for i := 0; i < l.Number; i++ {
		x := shared.NewDenseVec(l.Dim)
		for j := 0; j < l.Dim; j++ {
			x.SetVec(j, shared.MyFloat(l.Area.Ranges[j].Start+shared.MyFloat(rand.Float64())*lengths[j]))
		}
		wx := shared.NewDense(1, 1)
		wx.Dot(l.W.T(), x)
		y := wx.At(0, 0) + l.B + shared.MyFloat(rand.NormFloat64())*l.Sigma
		res[i] = shared.NewDenseVec(l.Dim + 1)
		for j := 0; j < l.Dim; j++ {
			res[i].SetVec(j, x.AtVec(j))
		}
		res[i].SetVec(l.Dim, y)
	}

	return res
}

// func (l *LinRegGenerator) Dump(w io.Writer) {
// 	writer := csv.NewWriter(w)
// 	points := l.Generate()
// 	for _, point := range points {
// 		if err := writer.Write([]string{fmt.Sprintf("%f", point.X), fmt.Sprintf("%f", point.Y)}); err != nil {
// 			log.Fatal(err)
// 		}
// 	}
// 	writer.Flush()
// 	if err := writer.Error(); err != nil {
// 		log.Fatal(err)
// 	}
// }
