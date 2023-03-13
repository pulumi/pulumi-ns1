[![Actions Status](https://github.com/pulumi/pulumi-ns1/workflows/master/badge.svg)](https://github.com/pulumi/pulumi-ns1/actions)
[![Slack](http://www.pulumi.com/images/docs/badges/slack.svg)](https://slack.pulumi.com)
[![NPM version](https://badge.fury.io/js/%40pulumi%2Fns1.svg)](https://www.npmjs.com/package/@pulumi/ns1)
[![Python version](https://badge.fury.io/py/pulumi-ns1.svg)](https://pypi.org/project/pulumi-ns1)
[![NuGet version](https://badge.fury.io/nu/pulumi.ns1.svg)](https://badge.fury.io/nu/pulumi.ns1)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/pulumi/pulumi-ns1/sdk/v3/go)](https://pkg.go.dev/github.com/pulumi/pulumi-ns1/sdk/v3/go)
[![License](https://img.shields.io/npm/l/%40pulumi%2Fpulumi.svg)](https://github.com/pulumi/pulumi-ns1/blob/master/LICENSE)

# NS1 Resource Provider

The NS1 resource provider for Pulumi lets you manage NS1
resources in your cloud programs. To use this package, please [install the
Pulumi CLI first](https://pulumi.io/).

## Installing

This package is available in many languages in the standard packaging formats.

### Node.js (Java/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

    $ npm install @pulumi/ns1

or `yarn`:

    $ yarn add @pulumi/ns1

### Python

To use from Python, install using `pip`:

    $ pip install pulumi_ns1

### Go

To use from Go, use `go get` to grab the latest version of the library

    $ go get github.com/pulumi/pulumi-ns1/sdk/v3

### .NET

To use from .NET, install using `dotnet add package`:

    $ dotnet add package Pulumi.Ns1

## Configuration

The following configuration points are available:

* ns1:apikey - (Required) NS1 API token. It must be provided, but it can also be sourced from the `NS1_APIKEY` 
  environment variable.
* ns1:endpoint - (Optional) NS1 API endpoint. For managed clients, this normally should not be set. Can also be sources
  via the `NS1_ENDPOINT` environment variable.
* ns1:enableDdi - (Optional) This sets the permission schema to a DDI-compatible schema. Users of the managed SaaS 
  product should not need to set this. Users of DDI should set this to true if managing teams, users, or API 
  keys through this provider.
* ns1:rateLimitParallelism - (Optional) Integer for parallelism amount (default is `10`). NS1 uses a token-based method
  for rate limiting API requests.

## Reference

For further information, please visit [the NS1 provider docs](https://www.pulumi.com/docs/intro/cloud-providers/ns1) 
or for detailed reference documentation, please visit [the API docs](https://www.pulumi.com/docs/reference/pkg/ns1).
