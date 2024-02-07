# Authenticating to NewRelic API

## Table of Contents

- [Authenticating to NewRelic API](#authenticating-to-newrelic-api)
  - [Table of Contents](#table-of-contents)
  - [Overview](#overview)
  - [Using `ProviderConfig`](#using-providerconfig)

## Overview

`crossplane-provider-newrelic` requires credentials to be provided in order to authenticate to the
NewRelic API. As this provider is upjet based, the authentication mechanism mirrors the terraform provider configuration.
This is described in detail in the
  [provider_configuration guide](https://registry.terraform.io/providers/newrelic/newrelic/latest/docs/guides/provider_configuration#configuration-via-the-provider-block).

## Using `ProviderConfig`

`crossplane-provider-newrelic` will be use credentials stored in a secret and referenced from a `ProviderConfig`.

The first thing that needs to be done is to create a `ProviderConfig`:

```console
$ cat <<EOF | kubectl apply -f -
apiVersion: provider-newrelic.upbound.io/v1beta1
kind: ProviderConfig
metadata:
  name: default-us
spec:
  credentials:
    source: Secret
    secretRef:
      name: example-creds-us
      namespace: crossplane-system
      key: credentials
EOF
```

Next create a secret as referenced from the `ProviderConfig`:

```console
$ cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Secret
metadata:
  name: example-creds-us
  namespace: crossplane-system
type: Opaque
stringData:
  credentials: |
    {
      "account_id": "NEW_RELIC_ACCOUNT_ID", # required.
      "api_key": "NEW_RELIC_API_KEY", # required.  Usually prefixed with 'NRAK'
      "region": "US" # optional.  Valid regions are US and EU
    }
EOF
```

You may now reference the `ProviderConfig` in your resources:
```console
$ cat <<EOF | kubectl apply -f -
---
apiVersion: alert.newrelic.upbound.io/v1alpha1
kind: Policy
metadata:
  name: example-policy
spec:
  forProvider:
    incidentPreference: PER_POLICY
    name: example-policy
  providerConfigRef:
    name: default-us
EOF
```

Multiple `ProviderConfigs` can be used to switch between credentials when more than
one target account or region is being reconciled by the provider.
