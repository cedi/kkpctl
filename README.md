# kkpctl

This tool aims to implement the [KKP](github.com/kubermatic/kubermatic) API as a useful CLI tool.
The usage should remind of kubectl.

kkpctl is written in Go and uses the [cobra](github.com/spf13/cobra) framework.

## Status

This is currently WIP

## Usage

For the full usage documentation, see the [docs](docs/commandline-usage.md)

## Quick-Start

1. Setup `kkpctl`
```
# Add your first cloud
$ kkpctl add cloud imke_prod https://imke.cloud/

# Add your first cloudprovider
$ kkpctl add provider optimist_prod --cloud imke_prod --type openstack --username 'cedric.kienzler@innovo-cloud.de' --password 'superSecurePassword!1337'

# Add the OSType to OSImage Mapping
$ kkpctl add osimage flatcar "Flatcar_Production 2020 - Latest" --cloud imke_prod --provider optimist_prod
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
$ kkpctl kubeconfig testcluster --cloud imke_prod --project testproject --set-kubectl
$ kubectl get pods -A
```
