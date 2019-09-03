package dga

import "strings"

type ResolvedTags struct {
	Tags []string
}

// Add adds an image tag to the list of tags that dga needs to work with.
func (r *ResolvedTags) Add(tag string) {
	if r.Tags == nil {
		r.Tags = []string{}
	}

	r.Tags = append(r.Tags, tag)
}

func TruncateShaHash(hash string) (shortHash string) {
	runes := []rune(hash)

	if len(runes) < 6 {
		return hash
	}

	return string(runes[0:6])
}

func CalculateShortReference(reference string) (shortRef string) {
	parts := strings.Split(reference, "/")

	return parts[len(parts)-1]
}

func CreateTagFormat(prefix string, suffix string) string {
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