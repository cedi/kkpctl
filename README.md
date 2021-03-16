[![GitHub license](https://img.shields.io/github/license/cedi/kkpctl.svg)](https://github.com/cedi/kkpctl/blob/main/LICENSE)
[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/cedi/kkpctl.svg)](https://github.com/cedi/kkpctl)
[![GoDoc reference example](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/cedi/kkpctl)
[![GoReportCard example](https://goreportcard.com/badge/github.com/cedi/kkpctl)](https://goreportcard.com/report/github.com/cedi/kkpctl)
[![Total alerts](https://img.shields.io/lgtm/alerts/g/cedi/kkpctl.svg?logo=lgtm&logoWidth=18)](https://lgtm.com/projects/g/cedi/kkpctl/alerts/)
[![workflow status](https://github.com/cedi/kkpctl/actions/workflows/go.yml/badge.svg)](https://github.com/cedi/kkpctl/actions)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com)


# kkpctl

This tool aims to implement the [KKP](https://github.com/kubermatic/kubermatic) API as a useful CLI tool.
The usage should remind of kubectl.


# Install from source

Pre-Requirement:
* Having the `go` installed
* your `$GOPATH` environment variable is set
* `$GOPATH/bin` is part of your `$PATH` environment variable
* Having `git` installed

```
mkdir -p $GOPATH/src/github.com/cedi/
git clone https://github.com/cedi/kkpctl.git $GOPATH/src/github.com/cedi/kkpctl
cd $GOPATH/src/github.com/cedi/kkpctl
make install_release
```

## Shell Completion

`kkpctl` comes with auto-completion for bash, zsh, fish, and PowerShell.
For more information see `kkpctl completion --help`

### Bash
```
$ source <(kkpctl completion bash)
```

To load completions for each session, execute once:
```
# Linux:
$ kkpctl completion bash > /etc/bash_completion.d/kkpctl

# macOS:
$ kkpctl completion bash > /usr/local/etc/bash_completion.d/kkpctl
```

### ZSH
If shell completion is not already enabled in your environment, you will need to enable it.
You can execute the following once
```
$ echo "autoload -U compinit; compinit" >> ~/.zshrc
```

To load completions for each session, execute once:
```
$ kkpctl completion zsh > "${fpath[1]}/_kkpctl"
```

You will need to start a new shell for this setup to take effect.

### fish
```
$ kkpctl completion fish | source
```

To load completions for each session, execute once:
```
kkpctl completion fish > ~/.config/fish/completions/kkpctl.fis
```

### PowerShell
```
PS> kkpctl completion powershell | Out-String | Invoke-Expression
```

To load completions for every new session, run:
```
PS> kkpctl completion powershell > kkpctl.ps1
```

and source this file from your PowerShell profile.

# Usage

For the full usage documentation see the [docs](docs/commandline-usage.md)

## Quick-Start

1. Setup `kkpctl`
```
# Add your first cloud
$ kkpctl config add cloud imke_prod https://imke.cloud/

# Add your first cloudprovider
$ kkpctl config add provider openstack --username "user@email.de" --password "my-super-secure-password" --tenant "internal-openstack-tenant" optimist

# Set your context to use the freshly added cloud
$ kkpctl ctx set cloud imke_prod
```

2. Login to KKP
```
# Note: This is a work around, until we have oidc-login available in kkpctl
$ kkpctl ctx set bearer 'ey03b....'
```

3. Create your first project
```
$ kkpctl add project testproject
```

4. Display your newly created project
```
$ kkpctl describe project 6tmbnhdl7h
```

5. Create your first cluster
```
$ kkpctl add cluster --project 6tmbnhdl7h --datacenter ix2 --provider optimist --version 1.18.13 --labels stage=dev kkpctltest
```

6. Describe your first cluster
```
$ kkpctl describe cluster --project 6tmbnhdl7h qvjdddt72t
```

7. Connect to your cluster, once it's ready
```
$ kkpctl get kubeconfig --project testproject qvjdddt72t -w 
$ export KUBECONFIG=./kubeconfig-admin-qvjdddt72t
$ kubectl get pods -A
```

# Contributing

## devcontainer
The easiest way to get your development enviroment up and running is using the [devcontainer](https://code.visualstudio.com/docs/remote/containers-tutorial).
Simply clone the repository, open the folder in your VSCode and accept the popup which asks if VSCode should restart in the dev-container. 

## Repository layout
```
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

## Makefile 
The repository ships with a makefile which makes it easier to build and install the application.
Useful Makefile targets are `build`, `release`, `test`, `test_all`, `install`, `install_release`, `clean`, and `vet`.

Most of them are self-explaining. I just want to point out the difference between a "development" and a "release" build.
* The development build is a regular `go build` with the `-race` flag enabled to detect race conditions easier.
* The release build is a regular `go build` withouth the `-race` flag, but with `-ldflags "-s -w"` to strip the debug symbols from the binary.

The `build` and `release` targets depend on `fmt` and `tidy`, so your code is always formated and your `go.mod` file is always tidy.

## Pull requests
We welcome pull requests. Feel free to dig through the [issues](https://github.com/cedi/kkpctl/issues) and jump in.