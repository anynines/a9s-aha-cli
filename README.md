# a9s-aha-cli

The command line tool `a9s-aha-cli` is a little helper for the aha.io world.


## Installation

```shell
$ brew tap anynines/tap
$ brew install a9s-aha-cli
```

## Usage

```shell
$ export AHA_USERNAME=username
$ export AHA_PASSWORD=password
$ export SLACK_URL=optional-slack-webhook-url

$ a9s-aha-cli alert
```

## Manual Release Building

```shell
git tag -a stable-0.1.0
GOOS=darwin GOARCH=amd64 go build -ldflags "-X github.com/anynines/a9s-aha-cli/pkg/version.Version=stable-0.1.0"
```
