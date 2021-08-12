# kkpctl

[![GitHub license](https://img.shields.io/github/license/cedi/kkpctl.svg)](https://github.com/cedi/kkpctl/blob/main/LICENSE)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/cedi/kkpctl.svg)](https://github.com/cedi/kkpctl)
[![GoDoc reference example](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/cedi/kkpctl)
[![GoReportCard example](https://goreportcard.com/badge/github.com/cedi/kkpctl)](https://goreportcard.com/report/github.com/cedi/kkpctl)
[![Total alerts](https://img.shields.io/lgtm/alerts/g/cedi/kkpctl.svg?logo=lgtm&logoWidth=18)](https://lgtm.com/projects/g/cedi/kkpctl/alerts/)
[![workflow status](https://github.com/cedi/kkpctl/actions/workflows/go.yml/badge.svg)](https://github.com/cedi/kkpctl/actions)

This tool aims to implement the [KKP](https://github.com/kubermatic/kubermatic) API as a useful CLI tool.
The usage should remind of kubectl.

## Usage

The usage of `kkpctl` should remind of `kubectl`.
For the full usage documentation see the [docs](docs/).

`kkpctl` comes with auto-completion right out of the box for bash, zsh, fish, and PowerShell.

```bash
kkpctl completion --help
```

## Quick-Start

### Download `kkpctl` and install it to your `$GOPATH/bin` folder

```bash
pushd /tmp

# Get the Download URL for your system
DOWNLOAD_PATH=$(curl -s https://api.github.com/repos/cedi/kkpctl/releases/latest | jq -r ".assets[]?.browser_download_url" | grep --color=never --ignore-case $(uname -s) | grep --color=never $(uname -m | sed 's/x86_64/amd64/g'))

FILENAME=$(echo $DOWNLOAD_PATH | awk -F'/' '{print $NF}')
FOLDER_NAME=$(echo $FILENAME | sed 's/.tar.gz//g')

# Download the tar.gz archive
curl -s -L $DOWNLOAD_PATH -o $FILENAME

# unpack the tar.gz archive
mkdir $FOLDER_NAME
tar -xzf $FILENAME -C $FOLDER_NAME

# install kkpctl to $GOPATH/bin/
cp $FOLDER_NAME/kkpctl $GOPATH/bin/kkpctl

popd
```

### Configure your KKP Cloud

#### Service Account Token

The simplest way to access your KKP Cloud is trough Service Account Token.
Please see the [KKP Documentation](https://docs.kubermatic.com/kubermatic/v2.17/guides/service_account/using_service_account/) for how to retrieve a service account token.

Once you got your Token, you can configure `kkpctl` to use the service account token instead of OIDC authentication using

```bash
kkpctl config add cloud imke --url https://imke.cloud --auth_token akdfjhklqwerhli2uh=
```

#### OIDC

Retrieve OIDC ClientID and Secret from your KKP installation

> __NOTE:__ Make sure, that `http://localhost:8000` is a valid RedirectURI in your dex configuration for the `kubermatic` client if you use this method.

> __Security Advise:__ It is better, if you register a separate OIDC Application for `kkpctl` that only allows redirect to `http://localhost:8080`. This is just meant a quick demo! Never do this in production!

```bash
# get the kubermatic client-secret
CLIENT_SECRET=$(kubectl get configmap -n oauth dex -ojson | jq '.data."config.yaml"' --raw-output | yq eval --tojson | jq '.staticClients | [ .[] | select( .id | contains("kubermatic")) ] | .[].secret' --raw-output)

# Add the kkp cloud with a name
kkpctl config add cloud kubermatic_dev --url https://dev.kubermatic.io --client_id kubermatic --client_secret $CLIENT_SECRET

# Set your context to use the freshly added cloud
kkpctl ctx set cloud kubermatic_dev
```

### Login to kkp

```bash
kkpctl oidc-login
```

And you're done!
Now, let's head over to the [working with kkpctl](docs/working-with-kkpctl.md) document where we go into more detail.

## Contributing

### devcontainer

The easiest way to get your development enviroment up and running is using the [devcontainer](https://code.visualstudio.com/docs/remote/containers-tutorial).
Simply clone the repository, open the folder in your VSCode and accept the popup which asks if VSCode should restart in the dev-container.

### Install from source

Pre-Requirement:

* Having the `go` installed
* your `$GOPATH` environment variable is set
* `$GOPATH/bin` is part of your `$PATH` environment variable
* Having `git` installed

```bash
mkdir -p $GOPATH/src/github.com/cedi/
git clone https://github.com/cedi/kkpctl.git $GOPATH/src/github.com/cedi/kkpctl
cd $GOPATH/src/github.com/cedi/kkpctl
make install_release
```

### Makefile

The repository ships with a makefile which makes it easier to build and install the application.
Useful Makefile targets are `build`, `release`, `test`, `test_all`, `install`, `install_release`, `clean`, and `vet`.

Most of them are self-explaining. I just want to point out the difference between a "development" and a "release" build.

* The development build is a regular `go build` with the `-race` flag enabled to detect race conditions easier.
* The release build is a regular `go build` withouth the `-race` flag, but with `-ldflags "-s -w"` to strip the debug symbols from the binary.

The `build` and `release` targets depend on `fmt` and `tidy`, so your code is always formated and your `go.mod` file is always tidy.

### Repository layout

```bash
├── .devcontainer   # the kkpctl repository comes with a devcontainer, so you can easily get started using VSCode
├── .github         # all github related configuration lays here
│   └── workflows   # contains the CI pipelines for kkpctl
├── .vscode         # contains a launch.json to get started with debugging the code
├── Makefile        # all the usefull aliases to build and test the project
├── cmd             # everything related to command line parsing is located in here. This is where you probably wanna start looking at
├── docs            # contains documentation
├── hack            # contains scripts for development
├── main.go         # the main entry point to the application
├── pkg             # most of the code is located here
│   ├── client      # the code that connects to the KKP API is here
│   ├── config      # contains the logic around the configuration of kkpctl
│   ├── describe    # the code that displays advanced information (describe) of a KKP API object
│   ├── model       # some additional data models we defined
│   ├── output      # similar as describe, but focuses on a simple output of an object
│   └── utils       # some utility functions which are usefull :)
└── tests           # contains mocks and test-files

```

## Pull requests

I warmly welcome pull requests. Feel free to dig through the [issues](https://github.com/cedi/kkpctl/issues) and jump in with whatever you feel comfortable with.
If you have new feature ideas, feel free to open a new issue and we can have a discussion.

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com)
