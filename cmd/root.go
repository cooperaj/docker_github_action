package cmd

import (
	"fmt"
	"github.com/cooperaj/docker_github_action/internal/app/dga"
	"os"

	"github.com/spf13/cobra"
)

var customFlag string
var gitRefFlag string
var gitShaFlag string
var latestFlag bool
var prefixFlag string
var suffixFlag string

var rootCmd = &cobra.Command{
	Use:   "dga",
	Short: "Docker Github Action for simplifying repetitive actions",
	Long: `Sometimes you'll want to put 4 tags on a container... and push 
them all to a registry... after logging in. This tool will aid 
in making the number of steps to do that as small as possible.`,
	Version: "v0.1",
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&gitRefFlag, "ref", "r", "", "Generate a tag using a supplied git reference (defaults to GITHUB_REF when not specified)")
	gitRefEnv := os.Getenv("GITHUB_REF")
	if gitRefEnv != "" {
		rootCmd.PersistentFlags().Lookup("ref").NoOptDefVal = gitRefEnv
	}

	rootCmd.PersistentFlags().StringVarP(&gitShaFlag, "sha", "s", "", "Generate a tag using a supplied git sha1 hash (defaults to GITHUB_SHA when not specified)")
	gitShaEnv := os.Getenv("GITHUB_SHA")
	if gitShaEnv != "" {
		rootCmd.PersistentFlags().Lookup("sha").NoOptDefVal = gitShaEnv
	}

	rootCmd.PersistentFlags().StringVarP(&customFlag, "custom", "c", "", "Generate a tag using the supplied value")

	rootCmd.PersistentFlags().BoolVarP(&latestFlag, "latest", "l", false, "Generate a latest tag")

	rootCmd.PersistentFlags().StringVarP(&prefixFlag, "prefix", "b", "", "Attach a prefix to the generated tags e.g. repo/image:prefix-latest")
	rootCmd.PersistentFlags().StringVarP(&suffixFlag, "suffix", "a", "", "Attach a suffix to the generated tags e.g. repo/image:latest-suffix")
}

// Execute is the entrypoint for the dga (Docker Github Action) program
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func processImageTags() dga.ResolvedTags {
	resolvedTags := dga.ResolvedTags{}

	tagFormat := dga.CreateTagFormat(prefixFlag, suffixFlag)

	if gitRefFlag != "" {
		gitRefFlag = fmt.Sprintf(tagFormat, dga.CalculateShortReference(gitRefFlag))
		resolvedTags.Add(gitRefFlag)
	}

	if gitShaFlag != "" {
		gitShaFlag = fmt.Sprintf(tagFormat, dga.TruncateShaHash(gitShaFlag))
		resolvedTags.Add(gitShaFlag)
	}

	if customFlag != "" {
		resolvedTags.Add(customFlag)
	}

	if latestFlag {
		latestStr := fmt.Sprintf(tagFormat, "latest")
		resolvedTags.Add(latestStr)
	}

	if len(resolvedTags.Tags) == 0 {
		fmt.Println("No tags specified. Use at least one tag flag")
		os.Exit(1)
	}

	return resolvedTags
}
