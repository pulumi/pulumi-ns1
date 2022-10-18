module github.com/pulumi/pulumi-ns1/provider/v2

go 1.16

replace (
	github.com/hashicorp/go-getter => github.com/hashicorp/go-getter v1.4.0
	github.com/ns1-terraform/terraform-provider-ns1 => github.com/pulumi/terraform-provider-ns1 v1.8.5-0.20220930222250-18f296f2e494
)

require (
	github.com/hashicorp/terraform-plugin-sdk v1.9.1
	github.com/ns1-terraform/terraform-provider-ns1 v1.9.0
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.21.0
	github.com/pulumi/pulumi/sdk/v3 v3.30.0
)
