module github.com/pulumi/pulumi-ns1/provider/v2

go 1.16

replace (
	github.com/hashicorp/go-getter v1.5.0 => github.com/hashicorp/go-getter v1.4.0
	github.com/ns1-terraform/terraform-provider-ns1 => github.com/pulumi/terraform-provider-ns1 v1.8.5-0.20210428113532-9ca93c8eccfa
)

require (
	github.com/hashicorp/terraform-plugin-sdk v1.9.1
	github.com/ns1-terraform/terraform-provider-ns1 v1.9.0
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.0.0
	github.com/pulumi/pulumi/pkg/v3 v3.0.0
	github.com/pulumi/pulumi/sdk/v3 v3.0.0
)
