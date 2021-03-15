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

## Status

This is currently WIP

## Development

### Acknowledgements

This project proudly uses [cobra](https://github.com/spf13/cobra) framework to handle the CLI parsing and [viper](https://github.com/spf13/viper) to read from environment variables

## Usage

For the full usage documentation, see the [docs](docs/commandline-usage.md)

### Quick-Start

1. Setup `kkpctl`
```
# Add your first cloud
$ kkpctl config add cloud imke_prod https://imke.cloud/

# Add your first cloudprovider
$ kkpctl config add provider openstack optimist --username "user@email.de" --password "my-super-secure-password" --tenant "internal-openstack-tenant"

# Set your context to use the freshly added cloud
$ kkpctl ctx set cloud imke_prod
```

2. Login to KKP
```
$ kkpctl oidc-login --cloud imke_prod
```

3. Create your first project
```
$ kkpctl add project testproject --cloud imke_prod
```

4. Display your newly created project
```
$ kkpctl get project testproject --cloud imke_prod
```

5. Create your first cluster
```
$ kkpctl add cluster testcluster --cloud imke_prod --project testproject --version '1.18.13' --labels 'stage=dev' --features 'AuditLogging' --provider 'optimist_prod' --datacenter 'es1' --node-name 'test' --node-flavor 'm1.small' --node-replica 3 --os-flavor 'flatcar'
```

6. Describe your first cluster
```
$ kkpctl describe cluster testcluster
```

7. Connect to your cluster, once it's ready
```
$ kkpctl get kubeconfig testcluster --cloud imke_prod --project testproject --set-kubectl
$ kubectl get pods -A
```
