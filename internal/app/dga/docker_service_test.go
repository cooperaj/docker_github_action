package dga

import (
	"context"
	"os"
	"testing"

	"github.com/docker/distribution/reference"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/stretchr/testify/assert"
)

var dockerTestClient *client.Client

func TestMain(m *testing.M) {
	var err error
	dockerTestClient, err = client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	// ensure busybox:latest exists for all the tests needed.
	path, err := reference.ParseNormalizedNamed("library/busybox:latest")
	if err != nil {
		panic(err)
	}

	io, err := dockerTestClient.ImagePull(context.Background(), path.String(), types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	_ = io.Close()

	os.Exit(m.Run())
}

func TestImageExists(t *testing.T) {
	type args struct {
		imageName string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Image exists when it does",
			args: args{"busybox:latest"},
			want: true,
		},
		{
			name: "Short image exists when it does",
			args: args{"busybox"},
			want: true,
		},
		{
			name: "Long image exists when it does",
			args: args{"library/busybox:latest"},
			want: true,
		},
		{
			name: "Image doesn't exist when it doesn't",
			args: args{"cooperaj/wontexist:latest"},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exists := ImageExists(tt.args.imageName)

			assert.Equal(t, tt.want, exists)
		})
	}
}

func TestTagImage(t *testing.T) {
	type args struct {
		baseImage string
		tag       string
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Can tag full image",
			args: args{
				baseImage: "library/busybox:latest",
				tag:       "newtag",
			},
		},
		{
			name: "Can tag short image",
			args: args{
				baseImage: "busybox:latest",
				tag:       "newtag",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TagImage(tt.args.baseImage, tt.args.tag)

			baseImagePath, err := reference.ParseNormalizedNamed(tt.args.baseImage)
			if err != nil {
				panic(err)
			}

			newImagePath, err := reference.WithTag(baseImagePath, tt.args.tag)
			if err != nil {
				panic(err)
			}

			filter := filters.NewArgs()
			filter.Add("reference", reference.FamiliarString(newImagePath))
			images, err := dockerClient.ImageList(context.Background(), types.ImageListOptions{
				All:     false,
				Filters: filter,
			})

			if err != nil {
				panic(err)
			}

			assert.NotEqual(t, len(images), 0)

			_, err = dockerClient.ImageRemove(context.Background(), reference.FamiliarString(newImagePath), types.ImageRemoveOptions{})
			if err != nil {
				panic(err)
			}
		})
	}
}
