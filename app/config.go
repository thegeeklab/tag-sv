package app

import "fmt"

var ErrMissingVersion = fmt.Errorf("version argument is required")

type Config struct {
	OutputFile   string
	Suffix       string
	SuffixStrict bool
	ExtraTags    string
	ForceLatest  bool
	IgnorePre    bool
}
