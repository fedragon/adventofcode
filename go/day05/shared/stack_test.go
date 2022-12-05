package shared

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_stack_Push(t *testing.T) {
	type fields struct {
		elems []string
	}
	type args struct {
		elem string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		expected []string
	}{
		{
			"pushes an element on top of an empty stack",
			fields{},
			args{"a"},
			[]string{"a"},
		},
		{
			"pushes an element on top of a non-empty stack",
			fields{elems: []string{"b"}},
			args{"a"},
			[]string{"b", "a"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Stack{
				elems: tt.fields.elems,
			}
			s.Push(tt.args.elem)

			assert.Equal(t, Stack{elems: tt.expected}, s)
		})
	}
}

func Test_stack_Pop(t *testing.T) {
	type fields struct {
		elems []string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"returns empty string for empty stack",
			fields{},
			"",
		},
		{
			"returns the element at the top of a non-empty stack",
			fields{elems: []string{"c", "b", "a"}},
			"a",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{}

			for _, elem := range tt.fields.elems {
				s.Push(elem)
			}

			assert.Equalf(t, tt.want, s.Pop(), "Pop()")
		})
	}
}

func Test_stack_PushN(t *testing.T) {
	type fields struct {
		elems []string
	}
	type args struct {
		elems []string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		expected []string
	}{
		{
			"pushes elements on top of an empty stack, preserving their order",
			fields{},
			args{[]string{"a", "b"}},
			[]string{"a", "b"},
		},
		{
			"pushes elements on top of a non-empty stack, preserving their order",
			fields{elems: []string{"c"}},
			args{[]string{"a", "b"}},
			[]string{"c", "a", "b"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Stack{
				elems: tt.fields.elems,
			}
			s.PushN(tt.args.elems...)

			assert.Equal(t, Stack{elems: tt.expected}, s)
		})
	}
}

func Test_stack_PopN(t *testing.T) {
	type fields struct {
		elems []string
	}
	type args struct {
		count int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		{
			"returns empty string for empty stack",
			fields{},
			args{3},
			nil,
		},
		{
			"returns the top N elements of a non-empty stack, in the same order",
			fields{elems: []string{"c", "b", "a"}},
			args{count: 3},
			[]string{"c", "b", "a"},
		},
		{
			"returns the top N elements of a non-empty stack, in the same order",
			fields{elems: []string{"e", "d", "c", "b", "a"}},
			args{count: 3},
			[]string{"c", "b", "a"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{}

			for _, elem := range tt.fields.elems {
				s.Push(elem)
			}

			assert.Equalf(t, tt.want, s.PopN(tt.args.count), "PopN(%v)", tt.args.count)
		})
	}
}
