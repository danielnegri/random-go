package random

import (
	"testing"
)

func TestSelectWeightedIndex(t *testing.T) {
	type args struct {
		choices []float64
		seed    int64
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "seed1",
			args: args{
				choices: []float64{0.35, 0.20, 0.40, 0.05},
				seed:    1,
			},
			want:    2,
			wantErr: false,
		},
		{
			name: "seed2",
			args: args{
				choices: []float64{0.35, 0.20, 0.40, 0.05},
				seed:    2,
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "seed2",
			args: args{
				choices: []float64{0.35, 0.20, 0.40, 0.05, 0.35, 0.20, 0.40, 0.05, 0.35, 0.20, 0.40, 0.05, 0.35, 0.20, 0.40, 0.05, 0.35, 0.20, 0.40, 0.05, 0.35, 0.20, 0.40, 0.05, 0.35, 0.20, 0.40, 0.05},
				seed:    1572644808, // 11/01/2019 @ 9:46pm (UTC)
			},
			want:    26,
			wantErr: false,
		},
		{
			name: "single choice",
			args: args{
				choices: []float64{1},
				seed:    0,
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "nil",
			args: args{
				choices: nil,
				seed:    0,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "empty",
			args: args{
				choices: []float64{},
				seed:    0,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SelectWeightedIndex(tt.args.choices, tt.args.seed)
			if (err != nil) != tt.wantErr {
				t.Errorf("SelectWeightedIndex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SelectWeightedIndex() got = %v, want %v", got, tt.want)
			}
		})
	}
}
