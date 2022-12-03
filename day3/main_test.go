package main

import (
	"testing"
)

func Test_priority(t *testing.T) {
	tests := []struct {
		name string
		item rune
		want int
	}{
		{
			"a = 1",
			'a',
			1,
		},
		{
			"z = 26",
			'z',
			26,
		},
		{
			"A = 27",
			'A',
			27,
		},
		{
			"Z = 52",
			'Z',
			52,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := priority(tt.item); got != tt.want {
				t.Errorf("priority() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sumBadges(t *testing.T) {
	tests := []struct {
		name  string
		lines []string
		want  int
	}{
		{

			"find badge in single group",
			[]string{
				"zBBtHnnHtwwHplmlRlzPLCpp",
				"vvhJccJFGFcNsdNNJbhJsJQplQMRLQMlfdfTPCLfQQCT",
				"GPhjcjhZDjWtnSVH",
			},
			42,
		},
		{

			"find and sum badges in multiple groups",
			[]string{
				"zBBtHnnHtwwHplmlRlzPLCpp",
				"vvhJccJFGFcNsdNNJbhJsJQplQMRLQMlfdfTPCLfQQCT",
				"GPhjcjhZDjWtnSVH",
				"BNhHVhrGNVTbDHdDJdJRPJdSQQSJwPjR",
				"lvtsfbsqzwSnJcvjSm",
				"MftttFLftZMLgtgMbltMqZzbDNrTpVGhNWrDTrpTGNpZGZhD",
				"VSSHcTgTtTdtllZlzmmbljTn",
				"RqMqsFfQLLFLQFMMfRLPZLvPpCfWrbpmCbjCnfjlWmnrmmnm",
				"hqRDqPDRsqNHwtHSNBZtJd",
			},
			96,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumBadges(tt.lines); got != tt.want {
				t.Errorf("sumBadges() = %v, want %v", got, tt.want)
			}
		})
	}
}
