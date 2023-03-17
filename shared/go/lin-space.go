package shared

type LinRange struct {
	Start MyFloat
	End   MyFloat
}

func (l *LinRange) Length() MyFloat {
	return l.End - l.Start
}

type NDArea struct {
	Ranges []LinRange
}

type NDPoint struct {
	X []MyFloat
}
