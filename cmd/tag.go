package cmd

import (
	"fmt"
	"github.com/cooperaj/docker_github_action/internal/app/dga"
	"github.com/spf13/cobra"
	"os"
)

var tagCmd = &cobra.Command{
	Use:   "tag [IMAGE TO TAG]",
	Short: "Tag a docker image",
	Args:  cobra.MinimumNArgs(1),
	Run:   tag,
}

func init() {
	rootCmd.AddCommand(tagCmd)
}

func tag(cmd *cobra.Command, args []string) {
	if !dga.ImageExists(args[0]) {
		fmt.Printf("Image with name %v does not exist\n", args[0])
		os.Exit(1)
	}

	resolvedTags := processImageTags()

	fmt.Printf("Tagging image %v\n", args[0])

	for _, tag := range resolvedTags.Tags {
		fmt.Printf("  %v\n", tag)
		dga.TagImage(args[0], tag)
	}
}