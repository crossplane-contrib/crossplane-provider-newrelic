# Provider NewRelic

# https://github.com/newrelic/terraform-provider-newrelic/tree/main/newrelic
# https://github.com/upbound/upjet-provider-template
# https://github.com/crossplane/upjet/blob/v0.10.0/docs/add-new-resource-short.md

`crossplane-provider-newrelic` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources for the
NewRelic API.

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://marketplace.upbound.io/providers/upbound/provider-newrelic):
```
up ctp provider install crossplane-contrib/crossplane-provider-newrelic:v0.1.0
```

Alternatively, you can use declarative installation:
```
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-newrelic
spec:
  package: crossplane-contrib/crossplane-provider-newrelic:v0.1.0
EOF
```

Notice that in this example Provider resource is referencing ControllerConfig with debug enabled.

You can see the API reference [here](https://doc.crds.dev/github.com/upbound/provider-newrelic).

## Authentication
See - /examples/providerconfig/providerconfig.yaml
See - /examples/providerconfig/secret.yaml.tmpl

# Currently Supported Resources
* `alert_policy` - /examples/alert/policy.yaml
* `nrql_alert_condition`
* `dashboardjson` - /examples/dashboard/dashboardjson.yaml
* `alert_destination`
* `alert_channel`
* `workflow` - /examples/workflow/workflow.yaml

## Developing

Run code-generation pipeline:
```console
go run cmd/generator/main.go "$PWD"
```

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build binary:

```console
make build
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/upbound/provider-newrelic/issues).
