package shared

import (
	"testing"
)

func TestDense_Dot(t *testing.T) {
	type args struct {
		a Matrix
		b Matrix
	}
	tests := []struct {
		name string
		args args
		want Matrix
	}{
		{
			name: "2x2 * 2x2",
			args: args{
				a: &Dense{
					Rows: 2,
					Cols: 2,
					Data: []MyFloat{1, 2, 3, 4},
				},
				b: &Dense{
					Rows: 2,
					Cols: 2,
					Data: []MyFloat{1, 2, 3, 4},
				},
			},
			want: &Dense{
				Rows: 2,
				Cols: 2,
				Data: []MyFloat{7, 10, 15, 22},
			},
		},
		{
			name: "2x2 * 2x3",
			args: args{
				a: &Dense{
					Rows: 2,
					Cols: 2,
					Data: []MyFloat{1, 2, 3, 4},
				},
				b: &Dense{
					Rows: 2,
					Cols: 3,
					Data: []MyFloat{1, 2, 3, 4, 5, 6},
				},
			},
			want: &Dense{
				Rows: 2,
				Cols: 3,
				Data: []MyFloat{9, 12, 15, 19, 26, 33},
			},
		},
		{
			name: "1x1 * 1x1",
			args: args{
				a: &Dense{
					Rows: 1,
					Cols: 1,
					Data: []MyFloat{2},
				},
				b: &Dense{
					Rows: 1,
					Cols: 1,
					Data: []MyFloat{3},
				},
			},
			want: &Dense{
				Rows: 1,
				Cols: 1,
				Data: []MyFloat{6},
			},
		},
		{
			name: "1x2 * 2x1",
			args: args{
				a: &Dense{
					Rows: 1,
					Cols: 2,
					Data: []MyFloat{1, 2},
				},
				b: &Dense{
					Rows: 2,
					Cols: 1,
					Data: []MyFloat{3, 4},
				},
			},
			want: &Dense{
				Rows: 1,
				Cols: 1,
				Data: []MyFloat{11},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ac, _ := tt.args.a.Dims()
			_, br := tt.args.b.Dims()
			d := NewDense(ac, br)
			d.Dot(tt.args.a, tt.args.b)
			if !d.Eq(tt.want) {
				t.Errorf("Dense.Dot() = %v, want %v", d, tt.want)
			}
		})
	}
}
