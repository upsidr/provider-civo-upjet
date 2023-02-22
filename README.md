# Provider Civo Upjet

`provider-civo-upjet` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/upbound/upjet) code
generation tools and exposes XRM-conformant managed resources for the
Civo API.

Right now there are 3 CRDs: Network, Firewall and KubernetesCluster.

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://marketplace.upbound.io/providers/upsidr/provider-civo):

```
up ctp provider install upsidr/provider-civo-upjet:v0.1.0
```

Alternatively, you can use declarative installation:

```
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-civo-upjet
spec:
  package: upsidr/provider-civo-upjet:v0.1.0
EOF
```

Notice that in this example Provider resource is referencing ControllerConfig with debug enabled.

You can see the API reference [here](https://doc.crds.dev/github.com/upsidr/provider-civo-upjet).

## Developing

This was generated using the [Upjet docs](https://github.com/upbound/upjet/blob/main/docs/generating-a-provider.md).

First, don't forget to fetch the build submodule:

```console
make submodules
```

Install Crossplane in a `kind` cluster. Create your `secret.yaml` from the template in `examples/providerconfig/`.

```console
# Create "crossplane-system" namespace if not exists
kubectl create namespace crossplane-system --dry-run=client -o yaml | kubectl apply -f -

kubectl apply -f package/crds

kubectl apply -f examples/providerconfig/secret.yaml
kubectl apply -f examples/providerconfig/providerconfig.yaml
```

Now run the provider locally against a Kubernetes cluster:

```console
make run
```

Generate code and CRDs:

```console
make generate
```

Build, push, and install:

```console
# don't forget to review first
make reviewable test
make all
```

Build binary:

```console
make build
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/upsidr/provider-civo-upjet/issues).
