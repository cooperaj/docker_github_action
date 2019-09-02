package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/cooperaj/docker_github_action/internal/app/dga"
	"github.com/spf13/cobra"
)

var gitRef string
var gitSha string
var latest bool
var prefix string
var suffix string

var tagCmd = &cobra.Command{
	Use:   "tag [IMAGE TO TAG]",
	Short: "Tag a docker image",
	Args:  cobra.MinimumNArgs(1),
	Run:   tag,
}

func init() {
	tagCmd.Flags().StringVarP(&gitRef, "ref", "r", "", "Generate a tag using a supplied git reference (defaults to GITHUB_REF when not specified)")
	gitRefEnv := os.Getenv("GITHUB_REF")
	if gitRefEnv != "" {
		tagCmd.Flags().Lookup("ref").NoOptDefVal = gitRefEnv
	}

	tagCmd.Flags().StringVarP(&gitSha, "sha", "s", "", "Generate a tag using a supplied git sha1 hash (defaults to GITHUB_SHA when not specified)")
	gitShaEnv := os.Getenv("GITHUB_SHA")
	if gitShaEnv != "" {
		tagCmd.Flags().Lookup("sha").NoOptDefVal = gitShaEnv
	}

	tagCmd.Flags().BoolVarP(&latest, "latest", "l", false, "Generate a latest tag")

	tagCmd.Flags().StringVarP(&prefix, "prefix", "b", "", "Attach a prefix to the generated tags e.g. repo/image:prefix-latest")
	tagCmd.Flags().StringVarP(&suffix, "suffix", "a", "", "Attach a suffix to the generated tags e.g. repo/image:latest-suffix")

	rootCmd.AddCommand(tagCmd)
}

func tag(cmd *cobra.Command, args []string) {
	if !dga.ImageExists(args[0]) {
		fmt.Printf("Image with name %v does not exist", args[0])
		os.Exit(1)
	}

	fmt.Printf("Tagging image %v\n", args[0])

	if gitRef != "" {
		gitRef = fmt.Sprintf(createTagFormat(), calculateShortReference(gitRef))
		fmt.Printf("  gitref: %v\n", gitRef)
		dga.TagImage(args[0], gitRef)
	}

	if gitSha != "" {
		gitSha = fmt.Sprintf(createTagFormat(), truncateShaHash(gitSha))
		fmt.Printf("  gitsha: %v\n", gitSha)
		dga.TagImage(args[0], gitSha)
	}

	if latest {
		latestStr := fmt.Sprintf(createTagFormat(), "latest")
		fmt.Printf("  latest: %v\n", latestStr)
		dga.TagImage(args[0], latestStr)
	}
}

func truncateShaHash(hash string) (shortHash string) {
	runes := []rune(hash)

	if len(runes) < 6 {
		return hash
	}

	return string(runes[0:6])
}

func calculateShortReference(reference string) (shortRef string) {
	parts := strings.Split(reference, "/")

	return parts[len(parts)-1]
}

func createTagFormat() string {
	var format strings.Builder

	if prefix != "" {
		format.WriteString(prefix)
	}

	format.WriteString("%v")

	if suffix != "" {
		format.WriteString(suffix)
	}

	return format.String()
}
