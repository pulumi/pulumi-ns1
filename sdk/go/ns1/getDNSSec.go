// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ns1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-ns1/sdk/v3/go/ns1/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides DNSSEC details about a NS1 Zone.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-ns1/sdk/v3/go/ns1"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			// Get DNSSEC details about a NS1 Zone.
//			exampleZone, err := ns1.NewZone(ctx, "example", &ns1.ZoneArgs{
//				Zone:   pulumi.String("terraform.example.io"),
//				Dnssec: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_ = ns1.GetDNSSecOutput(ctx, ns1.GetDNSSecOutputArgs{
//				Zone: exampleZone.Zone,
//			}, nil)
//			return nil
//		})
//	}
//
// ```
func GetDNSSec(ctx *pulumi.Context, args *GetDNSSecArgs, opts ...pulumi.InvokeOption) (*GetDNSSecResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDNSSecResult
	err := ctx.Invoke("ns1:index/getDNSSec:getDNSSec", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDNSSec.
type GetDNSSecArgs struct {
	// The name of the zone to get DNSSEC details for.
	Zone string `pulumi:"zone"`
}

// A collection of values returned by getDNSSec.
type GetDNSSecResult struct {
	// (Computed) - Delegation field is documented
	// below.
	Delegations []GetDNSSecDelegation `pulumi:"delegations"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// (Computed) - Keys field is documented below.
	Keys []GetDNSSecKey `pulumi:"keys"`
	Zone string         `pulumi:"zone"`
}

func GetDNSSecOutput(ctx *pulumi.Context, args GetDNSSecOutputArgs, opts ...pulumi.InvokeOption) GetDNSSecResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetDNSSecResultOutput, error) {
			args := v.(GetDNSSecArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ns1:index/getDNSSec:getDNSSec", args, GetDNSSecResultOutput{}, options).(GetDNSSecResultOutput), nil
		}).(GetDNSSecResultOutput)
}

// A collection of arguments for invoking getDNSSec.
type GetDNSSecOutputArgs struct {
	// The name of the zone to get DNSSEC details for.
	Zone pulumi.StringInput `pulumi:"zone"`
}

func (GetDNSSecOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDNSSecArgs)(nil)).Elem()
}

// A collection of values returned by getDNSSec.
type GetDNSSecResultOutput struct{ *pulumi.OutputState }

func (GetDNSSecResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDNSSecResult)(nil)).Elem()
}

func (o GetDNSSecResultOutput) ToGetDNSSecResultOutput() GetDNSSecResultOutput {
	return o
}

func (o GetDNSSecResultOutput) ToGetDNSSecResultOutputWithContext(ctx context.Context) GetDNSSecResultOutput {
	return o
}

// (Computed) - Delegation field is documented
// below.
func (o GetDNSSecResultOutput) Delegations() GetDNSSecDelegationArrayOutput {
	return o.ApplyT(func(v GetDNSSecResult) []GetDNSSecDelegation { return v.Delegations }).(GetDNSSecDelegationArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDNSSecResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDNSSecResult) string { return v.Id }).(pulumi.StringOutput)
}

// (Computed) - Keys field is documented below.
func (o GetDNSSecResultOutput) Keys() GetDNSSecKeyArrayOutput {
	return o.ApplyT(func(v GetDNSSecResult) []GetDNSSecKey { return v.Keys }).(GetDNSSecKeyArrayOutput)
}

func (o GetDNSSecResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v GetDNSSecResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDNSSecResultOutput{})
}
