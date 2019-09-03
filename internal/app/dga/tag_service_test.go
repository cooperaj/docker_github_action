package dga

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateShortReference(t *testing.T) {
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
			shortRef := CalculateShortReference(tt.args.reference)

			assert.Equal(t, tt.wantShortRef, shortRef)
		})
	}
}

func TestCreateTagFormat(t *testing.T) {
	type args struct {
		prefix string
		suffix string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Creates a correct prefix format",
			args: args {
				prefix: "test",
				suffix: "",
			},
			want: "test%v",
		},
		{
			name: "Creates a correct suffix format",
			args: args {
				prefix: "",
				suffix: "test",
			},
			want: "%vtest",
		},
		{
			name: "Creates a correct combined format",
			args: args {
				prefix: "test",
				suffix: "test",
			},
			want: "test%vtest",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			format := CreateTagFormat(tt.args.prefix, tt.args.suffix)

			assert.Equal(t, tt.want, format)
		})
	}
}

func TestResolvedTags_Add(t *testing.T) {
	type fields struct {
		Tags []string
	}
	type args struct {
		tag string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantCount int
	}{
		{
			name:   "Adds a tag correctly",
			fields: fields{
				Tags: nil,
			},
			args:   args{
				tag: "test",
			},
			wantCount: 1,
		},
		{
			name:   "Adds multiple tags correctly",
			fields: fields{
				Tags: []string{
					"test",
				},
			},
			args:   args{
				tag: "test2",
			},
			wantCount: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ResolvedTags{
				Tags: tt.fields.Tags,
			}

			r.Add(tt.args.tag)

			assert.Len(t, r.Tags, tt.wantCount)
			assert.Contains(t, r.Tags, tt.args.tag)
		})
	}
}

func TestTruncateShaHash(t *testing.T) {
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
			shortHash := TruncateShaHash(tt.args.hash)

			assert.Equal(t, tt.wantShortHash, shortHash)
		})
	}
}