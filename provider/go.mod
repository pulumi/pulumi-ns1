module github.com/pulumi/pulumi-ns1/provider

go 1.14

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/terraform-providers/terraform-provider-ns1 => github.com/pulumi/terraform-provider-ns1 v1.8.5-0.20200702083701-74e4b64874e0
)

require (
	github.com/hashicorp/terraform-plugin-sdk v1.9.1
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.5.3
	github.com/pulumi/pulumi/sdk/v2 v2.5.1-0.20200701223250-45d2fa95d60b
	github.com/terraform-providers/terraform-provider-ns1 v1.8.3
)
