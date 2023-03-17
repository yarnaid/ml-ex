package shared

type MyFloat float32

type Generator interface {
	Generate() []*Vector
}

type Point interface {
	Get(x_i int) []MyFloat
	Dim() int
}
