module github.com/pulumi/pulumi-ns1/provider

go 1.14

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/ns1-terraform/terraform-provider-ns1 => github.com/pulumi/terraform-provider-ns1 v1.8.5-0.20210107202504-2ffd42104c82
)

require (
	github.com/hashicorp/terraform-plugin-sdk v1.9.1
	github.com/ns1-terraform/terraform-provider-ns1 v1.9.0
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.13.2
	github.com/pulumi/pulumi/sdk/v2 v2.13.3-0.20201109230029-a6f8b9b205cd
)
