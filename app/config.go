package app

import "errors"

var ErrMissingVersion = errors.New("version argument is required")

type Config struct {
	OutputFile   string
	Suffix       string
	SuffixStrict bool
	ExtraTags    string
	ForceLatest  bool
	IgnorePre    bool
}
