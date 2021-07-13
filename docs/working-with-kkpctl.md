# Working with kkpctl

## Configuration

Currently there are two ways to configure your `kkpctl`.
The easiest option is to use `kkpctl config`, however due to the early stage of this project, this only works for the openstack cloudprovider for now.

### Openstack

```bash
# Add your first cloudprovider
kkpctl config add provider openstack --username "user@email.de" --password "my-super-secure-password" --tenant "internal-openstack-tenant" optimist
```

### Configure `kkpctl` manualy

```bash
# Create an empty configuration
kkpctl config generate -w

# Edit the just created configuration with your favorite Editor and fill in the details yourself
$EDITOR ~/.config/kkpctl/config.yaml
```

## Add your KKP Cloud

### Retrieve OIDC ClientID and Secret from your KKP installation

__NOTE:__ Make sure, that `http://localhost:8000` is a valid RedirectURI in your dex configuration for the `kubermatic` client if you use this method.

```bash
kubectl get configmap -n oauth dex -ojson | jq '.data."config.yaml"' --raw-output | yq eval --tojson | jq '.staticClients | [ .[] | select( .id | contains("kubermatic")) ] | .[].secret' --raw-output
```

__Security Advise:__ It is better, if you register a separate OIDC Application for `kkpctl` that only allows redirect to `http://localhost:8080`. This is just meant a quick demo! Never do this in production!

### Configuring kkpctl

```bash
# Add the kkp cloud with a name
kkpctl config add cloud kubermatic_dev --url https://dev.kubermatic.io --client_id kubermatic --client_secret dGVzdDEyMw==

# Set your context to use the freshly added cloud
kkpctl ctx set cloud kubermatic_dev
```

### Login to your kkp cloud

```bash
$ kkpctl oidc-login
You will now be taken to your browser for authentication
Authentication URL: https://dev.kubermatic.io/dex/auth?access_type=offline&client_id=kubermatic&redirect_uri=http%3A%2F%2Flocalhost%3A8000&response_type=code&scope=openid+email+profile&state=state
Authentication successful
```

## Work with projects

1. Create your first project

```bash
kkpctl add project testproject
```

2. List your projects

```bash
kkpctl get project
```

3. Display your newly created project

```bash
kkpctl describe project 6tmbnhdl7h
```

## Working with clusters

1. Create your first cluster

```bash
kkpctl add cluster --project 6tmbnhdl7h --datacenter ix2 --provider optimist --version 1.18.13 --labels stage=dev kkpctltest
```

2. List your clusters

```bash
kkpctl get cluster --project 6tmbnhdl7h
```

3. Describe your first cluster

```bash
kkpctl describe cluster --project 6tmbnhdl7h qvjdddt72t
```

## Retrieve your kubeconfig and configure it for use

1. Connect to your cluster, once it's ready

```bash
kkpctl get kubeconfig --project 6tmbnhdl7h qvjdddt72t -w
export KUBECONFIG=./kubeconfig-admin-qvjdddt72t
kubectl get pods -A
```
