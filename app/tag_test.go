package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultTags(t *testing.T) {
	tests := []struct {
		name       string
		versionStr string
		config     Config
		expected   []string
	}{
		{
			name:       "empty version",
			versionStr: "",
			config:     Config{ForceLatest: false},
			expected:   []string{"latest"},
		},
		{
			name:       "force latest",
			versionStr: "1.0.0",
			config:     Config{ForceLatest: true},
			expected:   []string{"latest", "1.0.0", "1.0", "1"},
		},
		{
			name:       "valid version",
			versionStr: "v1.2.3",
			config:     Config{ForceLatest: false},
			expected:   []string{"1.2.3", "1.2", "1"},
		},
		{
			name:       "prerelease version",
			versionStr: "v1.2.3-alpha",
			config:     Config{ForceLatest: false, IgnorePre: false},
			expected:   []string{"1.2.3-alpha"},
		},
		{
			name:       "ignoring prerelease",
			versionStr: "v1.2.3-alpha",
			config:     Config{ForceLatest: false, IgnorePre: true},
			expected:   []string{"1.2.3-alpha", "1.2.3", "1.2", "1"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.config.DefaultTags(tt.versionStr)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestTagSuffix(t *testing.T) {
	tests := []struct {
		name     string
		config   Config
		tags     []string
		expected []string
	}{
		{
			name:     "no suffix",
			config:   Config{SuffixStrict: false},
			tags:     []string{"1.0.0", "1.0"},
			expected: []string{"1.0.0", "1.0"},
		},
		{
			name:     "with suffix, not strict",
			config:   Config{Suffix: "beta", SuffixStrict: false},
			tags:     []string{"1.0.0", "1.0"},
			expected: []string{"1.0.0", "1.0", "1.0.0-beta", "1.0-beta"},
		},
		{
			name:     "with suffix, strict",
			config:   Config{Suffix: "beta", SuffixStrict: true},
			tags:     []string{"1.0", "latest"},
			expected: []string{"1.0-beta", "beta"},
		},
		{
			name:     "only latest tag with suffix",
			config:   Config{Suffix: "stable", SuffixStrict: false},
			tags:     []string{"1.0", "latest"},
			expected: []string{"1.0", "latest", "1.0-stable", "stable"},
		},
		{
			name:     "multiple tags with suffix",
			config:   Config{Suffix: "rc", SuffixStrict: false},
			tags:     []string{"1.0", "1.1"},
			expected: []string{"1.0", "1.1", "1.0-rc", "1.1-rc"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.config.TagSuffix(tt.tags)
			assert.Equal(t, tt.expected, result)
		})
	}
}
