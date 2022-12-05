package main

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
			s := stack{
				elems: tt.fields.elems,
			}
			s.Push(tt.args.elem)

			assert.Equal(t, stack{elems: tt.expected}, s)
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
			s := &stack{}

			for _, elem := range tt.fields.elems {
				s.Push(elem)
			}

			assert.Equalf(t, tt.want, s.Pop(), "Pop()")
		})
	}
}
