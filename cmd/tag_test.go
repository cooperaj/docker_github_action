package cmd

import (
	"testing"
)

func Test_truncateShaHash(t *testing.T) {
	type args struct {
		hash string
	}

	tests := []struct {
		name          string
		args          args
		wantShortHash string
	}{
		{
			name:          "Truncates long string",
			args:          args{hash: "AABBCCDDEEFF"},
			wantShortHash: "AABBCC",
		},
		{
			name:          "Returns short string",
			args:          args{hash: "AABB"},
			wantShortHash: "AABB",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotShortHash := truncateShaHash(tt.args.hash); gotShortHash != tt.wantShortHash {
				t.Errorf("truncateShaHash() = %v, want %v", gotShortHash, tt.wantShortHash)
			}
		})
	}
}

func Test_calculateShortReference(t *testing.T) {
	type args struct {
		reference string
	}
	tests := []struct {
		name         string
		args         args
		wantShortRef string
	}{
		{
			name:         "Returns branch name from full ref",
			args:         args{reference: "refs/head/master"},
			wantShortRef: "master",
		},
		{
			name:         "Returns tag name from full tag ref",
			args:         args{reference: "refs/tags/v1.1"},
			wantShortRef: "v1.1",
		},
		{
			name:         "Returns branch name from full remote ref",
			args:         args{reference: "refs/remotes/origin/master"},
			wantShortRef: "master",
		},
		{
			name:         "Returns branch name from full remote ref",
			args:         args{reference: "refs/remotes/origin/master"},
			wantShortRef: "master",
		},
		{
			name:         "Returns branch name from short ref",
			args:         args{reference: "master"},
			wantShortRef: "master",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotShortRef := calculateShortReference(tt.args.reference); gotShortRef != tt.wantShortRef {
				t.Errorf("calculateShortReference() = %v, want %v", gotShortRef, tt.wantShortRef)
			}
		})
	}
}
