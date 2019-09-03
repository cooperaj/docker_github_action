package dga

import (
	"context"
	"fmt"

	"github.com/docker/distribution/reference"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

var dockerClient *client.Client

func init() {
	var err error
	dockerClient, err = client.NewEnvClient()
	if err != nil {
		panic(err)
	}
}

// ImageExists Checks that a specified image exists in the docker host.
func ImageExists(imageName string) bool {
	sourceRef, err := reference.ParseNormalizedNamed(imageName)
	if err != nil {
		return false
	}

	filter := filters.NewArgs()
	filter.Add("reference", reference.FamiliarString(sourceRef))
	images, err := dockerClient.ImageList(context.Background(), types.ImageListOptions{
		All:     false,
		Filters: filter,
	})

	if err != nil {
		panic(err)
	}

	if len(images) > 0 {
		return true
	}

	return false
}

// TagImage Tags a given image with a new tag
func TagImage(baseImage string, tag string) {
	if !ImageExists(baseImage) {
		panic(fmt.Errorf("Image %v does not exist, cannot tag", baseImage))
	}

	sourceRef, err := reference.ParseNormalizedNamed(baseImage)
	if err != nil {
		panic(err)
	}

	var targetRef reference.Named
	targetRef, err = reference.WithTag(sourceRef, tag)
	if err != nil {
		panic(err)
	}

	err = dockerClient.ImageTag(context.Background(), sourceRef.String(), targetRef.String())
	if err != nil {
		panic(err)
	}
}
