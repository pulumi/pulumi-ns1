module github.com/pulumi/pulumi-ns1/provider

go 1.16

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/hashicorp/go-getter v1.5.0 => github.com/hashicorp/go-getter v1.4.0
	github.com/ns1-terraform/terraform-provider-ns1 => github.com/pulumi/terraform-provider-ns1 v1.8.5-0.20210330152659-729f7130349e
)

require (
	github.com/hashicorp/terraform-plugin-sdk v1.9.1
	github.com/ns1-terraform/terraform-provider-ns1 v1.9.0
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.23.0
	github.com/pulumi/pulumi/sdk/v2 v2.24.1
)
