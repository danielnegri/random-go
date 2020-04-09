package random

import (
	"testing"
)

func TestSelectWeightedIndex(t *testing.T) {
	tests := []struct {
		name    string
		choices []float64
		seed    int64
		want    int
		wantErr bool
	}{
		{
			name:    "test1",
			choices: []float64{0.35, 0.20, 0.40, 0.05},
			seed:    1,
			want:    2,
		},
		{
			name:    "test2",
			choices: []float64{0.35, 0.20, 0.40, 0.05},
			seed:    2,
			want:    0,
		},
		{
			name:    "test3",
			choices: []float64{0.35, 0.20, 0.40, 0.05, 0.35, 0.20, 0.40, 0.05, 0.35, 0.20, 0.40, 0.05, 0.35, 0.20, 0.40, 0.05, 0.35, 0.20, 0.40, 0.05, 0.35, 0.20, 0.40, 0.05, 0.35, 0.20, 0.40, 0.05},
			seed:    1572644808, // 11/01/2019 @ 9:46pm (UTC)
			want:    26,
		},
		{
			name:    "single",
			choices: []float64{1},
			seed:    0,
			want:    0,
		},
		{
			name:    "nil",
			choices: nil,
			seed:    0,
			want:    0,
		},
		{
			name:    "empty",
			choices: []float64{},
			seed:    0,
			want:    0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Seed(tt.seed)
			if got := SelectWeightedIndex(tt.choices); got != tt.want {
				t.Errorf("SelectWeightedIndex() got = %v, want %v", got, tt.want)
			}
		})
	}
}
