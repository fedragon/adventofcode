package day04

import "testing"

func TestRange_Overlaps(t *testing.T) {
	type args struct {
		a *Range
		b *Range
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"2-4, 5-7",
			args{a: &Range{Start: 2, End: 4}, b: &Range{Start: 5, End: 7}},
			false,
		},
		{
			"5-7, 2-4",
			args{a: &Range{Start: 5, End: 7}, b: &Range{Start: 2, End: 4}},
			false,
		},
		{
			"2-4, 4-6",
			args{a: &Range{Start: 2, End: 4}, b: &Range{Start: 4, End: 6}},
			true,
		},
		{
			"2-4, 3-3",
			args{a: &Range{Start: 2, End: 4}, b: &Range{Start: 3, End: 3}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.a.Overlaps(tt.args.b); got != tt.want {
				t.Errorf("Overlaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
