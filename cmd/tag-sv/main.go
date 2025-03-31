package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/thegeeklab/tag-sv/app"
	"github.com/urfave/cli/v3"
)

const FilePermLax = 0o644

//nolint:gochecknoglobals
var (
	BuildVersion = "devel"
	BuildDate    = "00000000"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	cli.VersionPrinter = func(c *cli.Command) {
		fmt.Printf("%s version=%s date=%s\n", c.Name, c.Version, BuildDate)
	}

	ctx := context.Background()
	config := app.Config{}

	app := &cli.Command{
		Name:      "tag-sv",
		Usage:     "Create tags from SemVer version string",
		ArgsUsage: "VERSION",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "output-file",
				Sources:     cli.EnvVars("TAG_SV_OUTPUT_FILE"),
				Destination: &config.OutputFile,
				Usage:       "path to write the tags output (default: stdout)",
			},
			&cli.StringFlag{
				Name:        "suffix",
				Sources:     cli.EnvVars("TAG_SV_SUFFIX"),
				Destination: &config.Suffix,
				Usage:       "add a suffix to all tags",
			},
			&cli.BoolFlag{
				Name:        "suffix-strict",
				Sources:     cli.EnvVars("TAG_SV_SUFFIX_STRICT"),
				Destination: &config.SuffixStrict,
				Usage:       "only output tags with suffixes when suffix is set",
			},
			&cli.StringFlag{
				Name:        "extra-tags",
				Sources:     cli.EnvVars("TAG_SV_EXTRA_TAGS"),
				Destination: &config.ExtraTags,
				Usage:       "additional tags to include, comma-separated",
			},
			&cli.BoolFlag{
				Name:        "force-latest",
				Sources:     cli.EnvVars("TAG_SV_FORCE_LATEST"),
				Destination: &config.ForceLatest,
				Usage:       "always include 'latest' tag in output",
			},
			&cli.BoolFlag{
				Name:        "ignore-pre",
				Sources:     cli.EnvVars("TAG_SV_IGNORE_PRERELEASE"),
				Destination: &config.IgnorePre,
				Usage:       "ignore pre-release and always get the full tag list",
			},
		},

		Action: func(_ context.Context, c *cli.Command) error {
			if c.NArg() < 1 {
				return app.ErrMissingVersion
			}
			v := c.Args().First()

			return run(config, v)
		},
	}

	if err := app.Run(ctx, os.Args); err != nil {
		log.Fatal().Err(err).Msg("Execution error")
	}
}

func run(a app.Config, v string) error {
	tags := a.DefaultTags(v)
	tags = a.TagSuffix(tags)
	tags = a.TagExtra(tags)

	output := strings.Join(tags, ",")

	if a.OutputFile != "" {
		return os.WriteFile(filepath.Clean(a.OutputFile), []byte(output), FilePermLax)
	}

	fmt.Println(output)

	return nil
}
