// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ns1

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
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
// 	"github.com/pulumi/pulumi-ns1/sdk/go/ns1"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ns1.LookupZone(ctx, &ns1.LookupZoneArgs{
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
