package common

import (
	"reflect"
	"testing"
)

func TestRange_Merge(t *testing.T) {
	type fields struct {
		Start int
		End   int
	}
	type args struct {
		a *Range
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Range
	}{
		{
			"merges two overlapping ranges",
			fields{Start: 0, End: 100},
			args{&Range{50, 150}},
			&Range{0, 150},
		},
		{
			"returns nil for two non-overlapping ranges",
			fields{Start: 0, End: 10},
			args{&Range{50, 150}},
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Range{
				Start: tt.fields.Start,
				End:   tt.fields.End,
			}
			if got := r.Merge(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
