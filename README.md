# tag-sv

Create tags from SemVer version string

[![Build Status](https://ci.thegeeklab.de/api/badges/thegeeklab/tag-sv/status.svg)](https://ci.thegeeklab.de/repos/thegeeklab/tag-sv)
[![Go Report Card](https://goreportcard.com/badge/github.com/thegeeklab/tag-sv)](https://goreportcard.com/report/github.com/thegeeklab/tag-sv)
[![GitHub contributors](https://img.shields.io/github/contributors/thegeeklab/tag-sv)](https://github.com/thegeeklab/tag-sv/graphs/contributors)
[![License: MIT](https://img.shields.io/github/license/thegeeklab/tag-sv)](https://github.com/thegeeklab/tag-sv/blob/main/LICENSE)

Simple tool to create a list of tags from a given SemVer version string.

## Installation

Prebuilt multiarch binaries are available for Linux only.

```Shell
curl -SsfL https://github.com/thegeeklab/tag-sv/releases/latest/download/tag-sv-linux-amd64 -o /usr/local/bin/tag-sv
chmod +x /usr/local/bin/tag-sv
```

## Build

Build the binary from source with the following command:

```Shell
make build
```

## Usage

```Shell
$ tag-sv --help
NAME:
   tag-sv - Create tags from SemVer version string

USAGE:
   tag-sv [global options] command [command options] VERSION

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --output-file value  path to write the tags output (default: stdout) [$TAG_SV_OUTPUT_FILE]
   --suffix value       add a suffix to all tags [$TAG_SV_SUFFIX]
   --suffix-strict      only output tags with suffixes when suffix is set (default: false) [$TAG_SV_SUFFIX_STRICT]
   --extra-tags value   additional tags to include, comma-separated [$TAG_SV_EXTRA_TAGS]
   --force-latest       always include 'latest' tag in output (default: false) [$TAG_SV_FORCE_LATEST]
   --ignore-pre         ignore pre-release and always get the full tag list (default: false) [$TAG_SV_IGNORE_PRERELEASE]
   --help, -h           show help
```

## Examples

```Shell
$ tag-sv "1.0.1"
# 1.0.1,1.0,1

$ tag-sv "0.1.0"
# 0.1.0, 0.1

## 'v' prefixes e.g. from git tags will be removed
$ tag-sv "v1.0.1"
# 1.0.1,1.0,1

## unsufficient semver version strings will be tried to convert automatically
## if conversion doesn't work return 'latest'
$ tag-sv "1.0"
# 1.0.0,1.0,1

$ tag-sv "1.0.0-beta"
# 1.0.0-beta

## ignore prerelease to always get a full list of tags
$ tag-sv --ignore-pre "1.0.0-beta"
# 1.0.0-beta,1.0.0,1.0,1

$ tag-sv --suffix amd64 "1.0.0"
# 1.0.0,1.0,1,1.0.0-amd64,1.0-amd64,1-amd64

$ tag-sv --suffix amd64 --suffix-strict "1.0.0"
# 1.0.0-amd64,1.0-amd64,1-amd64

$ tag-sv --extra-tags extra1,extra2 "1.0.0"
# 1.0.0,1.0,1,extra1,extra2
```

## Contributors

Special thanks to all [contributors](https://github.com/thegeeklab/tag-sv/graphs/contributors). If you would like to contribute, please see the [instructions](https://github.com/thegeeklab/tag-sv/blob/main/CONTRIBUTING.md).

## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/thegeeklab/tag-sv/blob/main/LICENSE) file for details.
