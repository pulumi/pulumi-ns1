// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ns1

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The provider type for the ns1 package. By default, resources use package-wide configuration
// settings, however an explicit `Provider` instance may be created and passed during resource
// construction to achieve fine-grained programmatic control over provider settings. See the
// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
type Provider struct {
	pulumi.ProviderResourceState
}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		args = &ProviderArgs{}
	}
	if args.Apikey == nil {
		args.Apikey = pulumi.StringPtr(getEnvOrDefault("", nil, "NS1_APIKEY").(string))
	}
	if args.Endpoint == nil {
		args.Endpoint = pulumi.StringPtr(getEnvOrDefault("", nil, "NS1_ENDPOINT").(string))
	}
	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:ns1", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	// The ns1 API key, this is required
	Apikey               *string `pulumi:"apikey"`
	EnableDdi            *bool   `pulumi:"enableDdi"`
	Endpoint             *string `pulumi:"endpoint"`
	IgnoreSsl            *bool   `pulumi:"ignoreSsl"`
	RateLimitParallelism *int    `pulumi:"rateLimitParallelism"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	// The ns1 API key, this is required
	Apikey               pulumi.StringPtrInput
	EnableDdi            pulumi.BoolPtrInput
	Endpoint             pulumi.StringPtrInput
	IgnoreSsl            pulumi.BoolPtrInput
	RateLimitParallelism pulumi.IntPtrInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}
