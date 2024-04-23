// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ns1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-ns1/sdk/v3/go/ns1/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides details of all available monitoring regions.
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
//			// Get details of all available monitoring regions.
//			_, err := ns1.GetMonitoringRegions(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetMonitoringRegions(ctx *pulumi.Context, args *GetMonitoringRegionsArgs, opts ...pulumi.InvokeOption) (*GetMonitoringRegionsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetMonitoringRegionsResult
	err := ctx.Invoke("ns1:index/getMonitoringRegions:getMonitoringRegions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMonitoringRegions.
type GetMonitoringRegionsArgs struct {
	// A set of the available monitoring regions. Regions is
	// documented below.
	Regions []GetMonitoringRegionsRegion `pulumi:"regions"`
}

// A collection of values returned by getMonitoringRegions.
type GetMonitoringRegionsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A set of the available monitoring regions. Regions is
	// documented below.
	Regions []GetMonitoringRegionsRegion `pulumi:"regions"`
}

func GetMonitoringRegionsOutput(ctx *pulumi.Context, args GetMonitoringRegionsOutputArgs, opts ...pulumi.InvokeOption) GetMonitoringRegionsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetMonitoringRegionsResult, error) {
			args := v.(GetMonitoringRegionsArgs)
			r, err := GetMonitoringRegions(ctx, &args, opts...)
			var s GetMonitoringRegionsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetMonitoringRegionsResultOutput)
}

// A collection of arguments for invoking getMonitoringRegions.
type GetMonitoringRegionsOutputArgs struct {
	// A set of the available monitoring regions. Regions is
	// documented below.
	Regions GetMonitoringRegionsRegionArrayInput `pulumi:"regions"`
}

func (GetMonitoringRegionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMonitoringRegionsArgs)(nil)).Elem()
}

// A collection of values returned by getMonitoringRegions.
type GetMonitoringRegionsResultOutput struct{ *pulumi.OutputState }

func (GetMonitoringRegionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMonitoringRegionsResult)(nil)).Elem()
}

func (o GetMonitoringRegionsResultOutput) ToGetMonitoringRegionsResultOutput() GetMonitoringRegionsResultOutput {
	return o
}

func (o GetMonitoringRegionsResultOutput) ToGetMonitoringRegionsResultOutputWithContext(ctx context.Context) GetMonitoringRegionsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetMonitoringRegionsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetMonitoringRegionsResult) string { return v.Id }).(pulumi.StringOutput)
}

// A set of the available monitoring regions. Regions is
// documented below.
func (o GetMonitoringRegionsResultOutput) Regions() GetMonitoringRegionsRegionArrayOutput {
	return o.ApplyT(func(v GetMonitoringRegionsResult) []GetMonitoringRegionsRegion { return v.Regions }).(GetMonitoringRegionsRegionArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetMonitoringRegionsResultOutput{})
}
