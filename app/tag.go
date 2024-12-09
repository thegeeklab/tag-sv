package app

import (
	"fmt"
	"strings"

	"github.com/Masterminds/semver/v3"
)

func (c *Config) DefaultTags(versionStr string) []string {
	defaultTag := []string{"latest"}

	var tags []string

	if c.ForceLatest {
		tags = append(tags, "latest")
	}

	if versionStr == "" {
		return defaultTag
	}

	versionStr = strings.TrimPrefix(versionStr, "refs/tags/")
	versionStr = strings.TrimPrefix(versionStr, "v")

	version, err := semver.NewVersion(versionStr)
	if err != nil {
		return defaultTag
	}

	if version.Prerelease() != "" {
		preTag := fmt.Sprintf("%d.%d.%d-%s", version.Major(), version.Minor(), version.Patch(), version.Prerelease())
		tags = append(tags, preTag)

		if !c.IgnorePre {
			return tags
		}
	}

	tags = append(tags, fmt.Sprintf("%d.%d.%d", version.Major(), version.Minor(), version.Patch()))
	tags = append(tags, fmt.Sprintf("%d.%d", version.Major(), version.Minor()))

	if version.Major() > 0 {
		tags = append(tags, fmt.Sprintf("%d", version.Major()))
	}

	return tags
}

func (c *Config) TagSuffix(tags []string) []string {
	if c.Suffix == "" {
		return tags
	}

	var result []string
	if !c.SuffixStrict {
		result = append(result, tags...)
	}

	for _, t := range tags {
		if t == "latest" {
			result = append(result, c.Suffix)
		} else {
			result = append(result, fmt.Sprintf("%s-%s", t, c.Suffix))
		}
	}

	return result
}

func (c *Config) TagExtra(tags []string) []string {
	if c.ExtraTags == "" {
		return tags
	}

	extraTags := strings.Split(c.ExtraTags, ",")
	for i := range extraTags {
		extraTags[i] = strings.TrimSpace(extraTags[i])
	}

	return append(tags, extraTags...)
}
