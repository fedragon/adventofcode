package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMonkey_Act(t *testing.T) {
	monkey2 := Monkey{ID: "2"}
	monkey3 := Monkey{ID: "3"}

	type fields struct {
		ID           string
		Items        []WorryLevel
		Operation    func(WorryLevel) WorryLevel
		TestDividend int
		WhenTrue     func([]*Monkey, WorryLevel)
		WhenFalse    func([]*Monkey, WorryLevel)
	}
	tests := []struct {
		name   string
		fields fields
		verify func()
	}{
		{
			"behaves like Monkey 0 in the example of part 1",
			fields{
				"0",
				[]WorryLevel{79, 98},
				func(level WorryLevel) WorryLevel {
					return level * 19
				},
				23,
				func(mks []*Monkey, item WorryLevel) {
					mks[2].Items = append(mks[2].Items, item)
				},
				func(mks []*Monkey, item WorryLevel) {
					mks[3].Items = append(mks[3].Items, item)
				},
			},
			func() {
				assert.Empty(t, monkey2.Items)
				assert.Equal(t, monkey3.Items, []WorryLevel{500, 620})
			},
		},
	}

	_ = os.Setenv("LOGS_ENABLED", "1")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Monkey{
				ID:            tt.fields.ID,
				Items:         tt.fields.Items,
				IncreaseWorry: tt.fields.Operation,
				TestOperand:   tt.fields.TestDividend,
				WhenTrue:      tt.fields.WhenTrue,
				WhenFalse:     tt.fields.WhenFalse,
			}

			m.Acts([]*Monkey{m, nil, &monkey2, &monkey3})

			tt.verify()
		})
	}
}
