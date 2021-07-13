# Some compiled help-pages

This is probably going to be outdated very fast. Please bear with me and help me keep this file up to date.

`kkpctl --help`

`kkpctl` has two different ways of working, which are similar to `kubectl` with `kubectx` and `kubens` installed.

Option 1)
Always specify `--cloud` and always specify `--project`.
This is similar to use `--namespace` in `kubectl`

Option 2)
Set the cloud and the project using a context using
`kkpctl config cloud set $cloudname`
and
`kkpctl config project set $projectname`
This is similar to use `kubens $namespace` with `kubectl`

## Basic Configuration

### Authentication

Since there is no OIDC-Auth (yet) we have to do authentication with a dirty hack

```
kkpctl ctx set bearer [bearertoken]
```

### Cloud

```
kkpctl add cloud $cloudname $url # Add a new cloud
kkpctl get cloud $cloudname # list a configured cloud
kkpctl delete cloud get $cloudname # delete a cloud
```

### Provider Type

Retreive Information about the providers supported by the KKP instalation

```
kkpctl get providertype

# cmdline options:
# 	--cloud $cloudname	# uses the specified cloud (Default: uses the cloud defined in ctx.cloud

```

### Datacenter

Retrieve Information about which Datacenters are supported by any given providertype in the KKP instance

```
kkpctl get datacenter $providertype

# cmdline options:
# 	--cloud $cloudname	# uses the specified cloud (Default: uses the cloud defined in ctx.cloud
```

### Provider

Define a KKP Provider

```
kkpctl add provider $providername

# cmdline options:
# 	--cloud $cloudname	# uses the specified cloud (Default: uses the cloud defined in ctx.cloud
#	--type	 			# MUST match the provider name in the configured KKP (i. e. if the KKP supports an openstack cloud provider, this field must be set to "openstack" to enable this config for this openstack provider)
#	--username			# The username for this cloud provider
#	--password			# the password for this cloud provider
#	--project			# the project where to deploy the worker nodes of a cluster
#
# example:
# kkpctl add provider optimist_prod --cloud imke-prod --type openstack --username "cedric.kienzler@innovo-cloud.de" --password "superSecurePassword!1337"

kkpctl get provider $providername
kkpctl describe provider $providername
kkpctl delete provider $providername
```

#### Provider's Operating Systems

Specify the OS Images for a given provider

```
kkpctl add osimage $ostype $imagename

# cmdline options:
# 	--cloud $cloudname			# uses the specified cloud (Default: uses the cloud defined in ctx.cloud
#	--provider $providername	# specify the mapping for a provider
#
# example:
# kkpctl add osimage flatcar "Flatcar_Production 2020 - Latest" --cloud imke_prod --provider optimist_prod

```

## Context

```
kkpctl ctx get # get all context variables
kkpctl ctx set cloud $cloudname # set the cloud name to the context, for the cloud to use
kkpctl ctx set project $projectname # set the project name to the context, for kkpctl to use
```

## Login

Since KKP uses OIDC to Login, we simply do something similar to `kubectl oidc-login`

```
kkpctl oidc-login

# cmdline options:
# 	--cloud $cloudname # logs in to the specified cloud (Default: uses the cloud defined in ctx.cloud
```

## Projects

### Get

To get the projects

```
kkpctl get project $project_name

# cmdline options:
# 	--cloud $cloudname	# uses the specified cloud (Default: uses the cloud defined in ctx.cloud
#	--all				# Gets all projects. Only works if your account is an Admin Account in KKP (Default: false)
#	--filter			# Filter by Labels. Specify a comma separated list of Labels to filter for. Multiple Labels are AND. (Example: --filter "project=kkp_test,stage=dev")
#
# example:
# kkpctl get project --cloud imke_prod
```

### Add

To add a project

```
kkpctl add project $project_name

# cmdline options:
# 	--cloud $cloudname	# uses the specified cloud (Default: uses the cloud defined in ctx.cloud
#	--labels			# A comma-separated list of labels to add to the project, in the format `key=value`. (Example: `--labels "stage=dev"`)
#
# example:
# kkpctl add project testproject --cloud imke_prod
```

### Delete

To delete a project

```
kkpctl delete project $project_name

# cmdline options:
# 	--cloud $cloudname	# uses the specified cloud (Default: uses the cloud defined in ctx.cloud
#	--filter			# Delete by Labels. Specify a comma separated list of Labels to filter for. Multiple Labels are AND. (Example: --filter "stage=dev")
```

### Describe

To Describe a project. This is similar to `get`, however it includes some additional informations about the project, like clusters and state

```
kkpctl describe project $project_name

# cmdline options:
# 	--cloud $cloudname	# uses the specified cloud (Default: uses the cloud defined in ctx.cloud
#	--filter			# Describe all by Labels. Specify a comma separated list of Labels to filter for. Multiple Labels are AND. (Example: --filter "stage=dev")
```

## Cluster

### Get

To get clusters

```
kkpctl get cluster $cluster_name

# cmdline options:
# 	--cloud $cloudname		# uses the specified cloud (Default: uses the cloud defined in ctx.cloud
#	--project $projectid	# uses the specified project (Default: uses the project defined in ctx.project)
#	--filter				# Filter by Labels. Specify a comma separated list of Labels to filter for. Multiple Labels are AND. (Example: --filter "project=kkp_test,stage=dev")
```

### Add

To add a project

```
kkpctl add cluster $cluster_name

# cmdline options:
# 	--cloud $cloudname		# logs in to the specified cloud (Default: uses the cloud defined in ctx.cloud
#	--project $projectid	# uses the specified project (Default: uses the project defined in ctx.project)
#	--version				# Kubernetes Version to use (Default: uses the cluster version defined as default in the kkp cloud)
#	--labels				# A comma-separated list of labels to add to the cluster, in the format `key=value`. (Example: `--labels "stage=dev"`)
#	--features				# Enable certain features. Currently supported: "AuditLogging, PodSecurityPolicy, PodNodeSelector"
#	--provider				# Use the provider-name you specified in your configuration
#	--datacenter			# Select in wich datacenter to deploy your cluster
#	--node-name				# Define a custom name-prefix for your machine deployment (Optional. Default uses server-side generation)
#	--node-flavor			# Specify the node-flavor to use
#	--node-replica			# Specify how many nodes should be deployed
#	--os-flavor				# Specify which Operating System to use. The Image has to be defined in your clouds config.
#	--node-labels			# A comma-separated list of labels to add to the node-deployment, in the format `key=value`. (Example: `--node-labels "type=compute"`)
#
# example:
# kkpctl add cluster testcluster \
#	--cloud imke_prod \
#	--project abc123de45 \
#	--version '1.18.13' \
#	--labels 'stage=dev' \
#	--features 'AuditLogging' \
#	--provider 'optimist_prod' \
#	--datacenter 'es1' \
#	--node-name 'test' \
#	--node-flavor 'm1.small' \
#	--node-replica 3 \
#	--os-flavor 'flatcar'
```

### Delete

To delete a cluster.
If no cluster-id is specified, and no filter is given, nothing is deleted.
If no cluster-id is specified but `--all`, all clusters in the specified project are deleted.

```
kkpctl delete cluster $cluster_id

# cmdline options:
# 	--cloud $cloudname		# logs in to the specified cloud (Default: uses the cloud defined in ctx.cloud
#	--project $projectid	# uses the specified project (Default: uses the project defined in ctx.project)
#	--filter				# Delete by Labels. Specify a comma separated list of Labels to filter for. Multiple Labels are AND. (Example: --filter "stage=dev")
#	--all					# Deletes all clusters in the given project
```

### Describe

To Describe a cluster. This is similar to `get`, however it includes some additional informations about the cluster, like control plane status, machine deployment status, and events

```
kkpctl describe cluster $cluster_id

# cmdline options:
# 	--cloud $cloudname		# logs in to the specified cloud (Default: uses the cloud defined in ctx.cloud
#	--project $projectid	# uses the specified project (Default: uses the project defined in ctx.project)
#	--filter				# Delete by Labels. Specify a comma separated list of Labels to filter for. Multiple Labels are AND. (Example: --filter "project=kkp_test,stage=dev")
```

### Kubeconfig

To retreive the kube-config of a KKP cluster.

```
kkpctl get kubeconfig $clusterid

# cmdline options:
# 	--cloud $cloudname		# logs in to the specified cloud (Default: uses the cloud defined in ctx.cloud
#	--project $projectid	# uses the specified project (Default: uses the project defined in ctx.project)
#	--set-kubectl 			# Downloads the Kubeconfig and automatically sets up `kubectl` to use this kubeconfig

```
