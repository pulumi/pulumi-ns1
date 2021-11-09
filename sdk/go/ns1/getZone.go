// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ns1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides details about a NS1 Zone. Use this if you would simply like to read
// information from NS1 into your configurations. For read/write operations, you
// should use a resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-ns1/sdk/v2/go/ns1"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ns1.LookupZone(ctx, &GetZoneArgs{
// 			Zone: "terraform.example.io",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupZone(ctx *pulumi.Context, args *LookupZoneArgs, opts ...pulumi.InvokeOption) (*LookupZoneResult, error) {
	var rv LookupZoneResult
	err := ctx.Invoke("ns1:index/getZone:getZone", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getZone.
type LookupZoneArgs struct {
	// List of additional IPv4 addresses for the primary
	// zone.
	AdditionalPrimaries []string `pulumi:"additionalPrimaries"`
	// The domain name of the zone.
	Zone string `pulumi:"zone"`
}

// A collection of values returned by getZone.
type LookupZoneResult struct {
	// List of additional IPv4 addresses for the primary
	// zone.
	AdditionalPrimaries []string `pulumi:"additionalPrimaries"`
	// Authoritative Name Servers.
	DnsServers string `pulumi:"dnsServers"`
	// Whether or not DNSSEC is enabled for the zone.
	Dnssec bool `pulumi:"dnssec"`
	// The SOA Expiry.
	Expiry int `pulumi:"expiry"`
	// The SOA Hostmaster.
	Hostmaster string `pulumi:"hostmaster"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The linked target zone.
	Link string `pulumi:"link"`
	// List of network IDs (`int`) for which the zone should be made
	// available. Default is network 0, the primary NSONE Global Network.
	Networks []int `pulumi:"networks"`
	// The SOA NX TTL.
	NxTtl int `pulumi:"nxTtl"`
	// The primary zones' IPv4 address.
	Primary string `pulumi:"primary"`
	// The SOA Refresh.
	Refresh int `pulumi:"refresh"`
	// The SOA Retry.
	Retry int `pulumi:"retry"`
	// List of secondary servers. Secondaries is
	// documented below.
	Secondaries []GetZoneSecondary `pulumi:"secondaries"`
	// The SOA TTL.
	Ttl  int    `pulumi:"ttl"`
	Zone string `pulumi:"zone"`
}

func LookupZoneOutput(ctx *pulumi.Context, args LookupZoneOutputArgs, opts ...pulumi.InvokeOption) LookupZoneResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupZoneResult, error) {
			args := v.(LookupZoneArgs)
			r, err := LookupZone(ctx, &args, opts...)
			return *r, err
		}).(LookupZoneResultOutput)
}

// A collection of arguments for invoking getZone.
type LookupZoneOutputArgs struct {
	// List of additional IPv4 addresses for the primary
	// zone.
	AdditionalPrimaries pulumi.StringArrayInput `pulumi:"additionalPrimaries"`
	// The domain name of the zone.
	Zone pulumi.StringInput `pulumi:"zone"`
}

func (LookupZoneOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupZoneArgs)(nil)).Elem()
}

// A collection of values returned by getZone.
type LookupZoneResultOutput struct{ *pulumi.OutputState }

func (LookupZoneResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupZoneResult)(nil)).Elem()
}

func (o LookupZoneResultOutput) ToLookupZoneResultOutput() LookupZoneResultOutput {
	return o
}

func (o LookupZoneResultOutput) ToLookupZoneResultOutputWithContext(ctx context.Context) LookupZoneResultOutput {
	return o
}

// List of additional IPv4 addresses for the primary
// zone.
func (o LookupZoneResultOutput) AdditionalPrimaries() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupZoneResult) []string { return v.AdditionalPrimaries }).(pulumi.StringArrayOutput)
}

// Authoritative Name Servers.
func (o LookupZoneResultOutput) DnsServers() pulumi.StringOutput {
	return o.ApplyT(func(v LookupZoneResult) string { return v.DnsServers }).(pulumi.StringOutput)
}

// Whether or not DNSSEC is enabled for the zone.
func (o LookupZoneResultOutput) Dnssec() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupZoneResult) bool { return v.Dnssec }).(pulumi.BoolOutput)
}

// The SOA Expiry.
func (o LookupZoneResultOutput) Expiry() pulumi.IntOutput {
	return o.ApplyT(func(v LookupZoneResult) int { return v.Expiry }).(pulumi.IntOutput)
}

// The SOA Hostmaster.
func (o LookupZoneResultOutput) Hostmaster() pulumi.StringOutput {
	return o.ApplyT(func(v LookupZoneResult) string { return v.Hostmaster }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupZoneResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupZoneResult) string { return v.Id }).(pulumi.StringOutput)
}

// The linked target zone.
func (o LookupZoneResultOutput) Link() pulumi.StringOutput {
	return o.ApplyT(func(v LookupZoneResult) string { return v.Link }).(pulumi.StringOutput)
}

// List of network IDs (`int`) for which the zone should be made
// available. Default is network 0, the primary NSONE Global Network.
func (o LookupZoneResultOutput) Networks() pulumi.IntArrayOutput {
	return o.ApplyT(func(v LookupZoneResult) []int { return v.Networks }).(pulumi.IntArrayOutput)
}

// The SOA NX TTL.
func (o LookupZoneResultOutput) NxTtl() pulumi.IntOutput {
	return o.ApplyT(func(v LookupZoneResult) int { return v.NxTtl }).(pulumi.IntOutput)
}

// The primary zones' IPv4 address.
func (o LookupZoneResultOutput) Primary() pulumi.StringOutput {
	return o.ApplyT(func(v LookupZoneResult) string { return v.Primary }).(pulumi.StringOutput)
}

// The SOA Refresh.
func (o LookupZoneResultOutput) Refresh() pulumi.IntOutput {
	return o.ApplyT(func(v LookupZoneResult) int { return v.Refresh }).(pulumi.IntOutput)
}

// The SOA Retry.
func (o LookupZoneResultOutput) Retry() pulumi.IntOutput {
	return o.ApplyT(func(v LookupZoneResult) int { return v.Retry }).(pulumi.IntOutput)
}

// List of secondary servers. Secondaries is
// documented below.
func (o LookupZoneResultOutput) Secondaries() GetZoneSecondaryArrayOutput {
	return o.ApplyT(func(v LookupZoneResult) []GetZoneSecondary { return v.Secondaries }).(GetZoneSecondaryArrayOutput)
}

// The SOA TTL.
func (o LookupZoneResultOutput) Ttl() pulumi.IntOutput {
	return o.ApplyT(func(v LookupZoneResult) int { return v.Ttl }).(pulumi.IntOutput)
}

func (o LookupZoneResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupZoneResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupZoneResultOutput{})
}
