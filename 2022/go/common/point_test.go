package common

import "testing"

func TestPoint_ManhattanDistance(t *testing.T) {
	type fields struct {
		X int
		Y int
	}
	type args struct {
		t Point
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			"distance((2,18), (-2,15)) = 7",
			fields{X: 2, Y: 18},
			args{Point{X: -2, Y: 15}},
			7,
		},
		{
			"distance((8,7), (2,10)) = 9",
			fields{X: 8, Y: 7},
			args{Point{X: 2, Y: 10}},
			9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Point{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := p.ManhattanDistance(&tt.args.t); got != tt.want {
				t.Errorf("ManhattanDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
