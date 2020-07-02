// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ns1

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type Zone struct {
	pulumi.CustomResourceState

	// List of additional IPv4 addresses for the primary
	// zone. Conflicts with `secondaries`.
	AdditionalPrimaries  pulumi.StringArrayOutput `pulumi:"additionalPrimaries"`
	AutogenerateNsRecord pulumi.BoolPtrOutput     `pulumi:"autogenerateNsRecord"`
	// (Computed) Authoritative Name Servers.
	DnsServers pulumi.StringOutput `pulumi:"dnsServers"`
	// Whether or not DNSSEC is enabled for the zone.
	// Note that DNSSEC must be enabled on the account by support for this to be set
	// to `true`.
	Dnssec pulumi.BoolOutput `pulumi:"dnssec"`
	// The SOA Expiry. Conflicts with `primary` and
	// `additionalPrimaries` (default must be accepted).
	Expiry pulumi.IntOutput `pulumi:"expiry"`
	// (Computed) The SOA Hostmaster.
	Hostmaster pulumi.StringOutput `pulumi:"hostmaster"`
	// The target zone(domain name) to link to.
	Link pulumi.StringPtrOutput `pulumi:"link"`
	// - List of network IDs (`int`) for which the zone
	//   should be made available. Default is network 0, the primary NSONE Global
	//   Network. Normally, you should not have to worry about this.
	Networks pulumi.IntArrayOutput `pulumi:"networks"`
	// The SOA NX TTL. Conflicts with `primary` and
	// `additionalPrimaries` (default must be accepted).
	NxTtl pulumi.IntOutput `pulumi:"nxTtl"`
	// The primary zones' IPv4 address. This makes the zone a
	// secondary. Conflicts with `secondaries`.
	Primary pulumi.StringPtrOutput `pulumi:"primary"`
	// The SOA Refresh. Conflicts with `primary` and
	// `additionalPrimaries` (default must be accepted).
	Refresh pulumi.IntOutput `pulumi:"refresh"`
	// The SOA Retry. Conflicts with `primary` and
	// `additionalPrimaries` (default must be accepted).
	Retry pulumi.IntOutput `pulumi:"retry"`
	// List of secondary servers. This makes the zone a
	// primary. Conflicts with `primary` and `additionalPrimaries`.
	// Secondaries is documented below.
	Secondaries ZoneSecondaryArrayOutput `pulumi:"secondaries"`
	// The SOA TTL.
	Ttl pulumi.IntOutput `pulumi:"ttl"`
	// The domain name of the zone.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewZone registers a new resource with the given unique name, arguments, and options.
func NewZone(ctx *pulumi.Context,
	name string, args *ZoneArgs, opts ...pulumi.ResourceOption) (*Zone, error) {
	if args == nil || args.Zone == nil {
		return nil, errors.New("missing required argument 'Zone'")
	}
	if args == nil {
		args = &ZoneArgs{}
	}
	var resource Zone
	err := ctx.RegisterResource("ns1:index/zone:Zone", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetZone gets an existing Zone resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetZone(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ZoneState, opts ...pulumi.ResourceOption) (*Zone, error) {
	var resource Zone
	err := ctx.ReadResource("ns1:index/zone:Zone", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Zone resources.
type zoneState struct {
	// List of additional IPv4 addresses for the primary
	// zone. Conflicts with `secondaries`.
	AdditionalPrimaries  []string `pulumi:"additionalPrimaries"`
	AutogenerateNsRecord *bool    `pulumi:"autogenerateNsRecord"`
	// (Computed) Authoritative Name Servers.
	DnsServers *string `pulumi:"dnsServers"`
	// Whether or not DNSSEC is enabled for the zone.
	// Note that DNSSEC must be enabled on the account by support for this to be set
	// to `true`.
	Dnssec *bool `pulumi:"dnssec"`
	// The SOA Expiry. Conflicts with `primary` and
	// `additionalPrimaries` (default must be accepted).
	Expiry *int `pulumi:"expiry"`
	// (Computed) The SOA Hostmaster.
	Hostmaster *string `pulumi:"hostmaster"`
	// The target zone(domain name) to link to.
	Link *string `pulumi:"link"`
	// - List of network IDs (`int`) for which the zone
	//   should be made available. Default is network 0, the primary NSONE Global
	//   Network. Normally, you should not have to worry about this.
	Networks []int `pulumi:"networks"`
	// The SOA NX TTL. Conflicts with `primary` and
	// `additionalPrimaries` (default must be accepted).
	NxTtl *int `pulumi:"nxTtl"`
	// The primary zones' IPv4 address. This makes the zone a
	// secondary. Conflicts with `secondaries`.
	Primary *string `pulumi:"primary"`
	// The SOA Refresh. Conflicts with `primary` and
	// `additionalPrimaries` (default must be accepted).
	Refresh *int `pulumi:"refresh"`
	// The SOA Retry. Conflicts with `primary` and
	// `additionalPrimaries` (default must be accepted).
	Retry *int `pulumi:"retry"`
	// List of secondary servers. This makes the zone a
	// primary. Conflicts with `primary` and `additionalPrimaries`.
	// Secondaries is documented below.
	Secondaries []ZoneSecondary `pulumi:"secondaries"`
	// The SOA TTL.
	Ttl *int `pulumi:"ttl"`
	// The domain name of the zone.
	Zone *string `pulumi:"zone"`
}

type ZoneState struct {
	// List of additional IPv4 addresses for the primary
	// zone. Conflicts with `secondaries`.
	AdditionalPrimaries  pulumi.StringArrayInput
	AutogenerateNsRecord pulumi.BoolPtrInput
	// (Computed) Authoritative Name Servers.
	DnsServers pulumi.StringPtrInput
	// Whether or not DNSSEC is enabled for the zone.
	// Note that DNSSEC must be enabled on the account by support for this to be set
	// to `true`.
	Dnssec pulumi.BoolPtrInput
	// The SOA Expiry. Conflicts with `primary` and
	// `additionalPrimaries` (default must be accepted).
	Expiry pulumi.IntPtrInput
	// (Computed) The SOA Hostmaster.
	Hostmaster pulumi.StringPtrInput
	// The target zone(domain name) to link to.
	Link pulumi.StringPtrInput
	// - List of network IDs (`int`) for which the zone
	//   should be made available. Default is network 0, the primary NSONE Global
	//   Network. Normally, you should not have to worry about this.
	Networks pulumi.IntArrayInput
	// The SOA NX TTL. Conflicts with `primary` and
	// `additionalPrimaries` (default must be accepted).
	NxTtl pulumi.IntPtrInput
	// The primary zones' IPv4 address. This makes the zone a
	// secondary. Conflicts with `secondaries`.
	Primary pulumi.StringPtrInput
	// The SOA Refresh. Conflicts with `primary` and
	// `additionalPrimaries` (default must be accepted).
	Refresh pulumi.IntPtrInput
	// The SOA Retry. Conflicts with `primary` and
	// `additionalPrimaries` (default must be accepted).
	Retry pulumi.IntPtrInput
	// List of secondary servers. This makes the zone a
	// primary. Conflicts with `primary` and `additionalPrimaries`.
	// Secondaries is documented below.
	Secondaries ZoneSecondaryArrayInput
	// The SOA TTL.
	Ttl pulumi.IntPtrInput
	// The domain name of the zone.
	Zone pulumi.StringPtrInput
}

func (ZoneState) ElementType() reflect.Type {
	return reflect.TypeOf((*zoneState)(nil)).Elem()
}

type zoneArgs struct {
	// List of additional IPv4 addresses for the primary
	// zone. Conflicts with `secondaries`.
	AdditionalPrimaries  []string `pulumi:"additionalPrimaries"`
	AutogenerateNsRecord *bool    `pulumi:"autogenerateNsRecord"`
	// Whether or not DNSSEC is enabled for the zone.
	// Note that DNSSEC must be enabled on the account by support for this to be set
	// to `true`.
	Dnssec *bool `pulumi:"dnssec"`
	// The SOA Expiry. Conflicts with `primary` and
	// `additionalPrimaries` (default must be accepted).
	Expiry *int `pulumi:"expiry"`
	// The target zone(domain name) to link to.
	Link *string `pulumi:"link"`
	// - List of network IDs (`int`) for which the zone
	//   should be made available. Default is network 0, the primary NSONE Global
	//   Network. Normally, you should not have to worry about this.
	Networks []int `pulumi:"networks"`
	// The SOA NX TTL. Conflicts with `primary` and
	// `additionalPrimaries` (default must be accepted).
	NxTtl *int `pulumi:"nxTtl"`
	// The primary zones' IPv4 address. This makes the zone a
	// secondary. Conflicts with `secondaries`.
	Primary *string `pulumi:"primary"`
	// The SOA Refresh. Conflicts with `primary` and
	// `additionalPrimaries` (default must be accepted).
	Refresh *int `pulumi:"refresh"`
	// The SOA Retry. Conflicts with `primary` and
	// `additionalPrimaries` (default must be accepted).
	Retry *int `pulumi:"retry"`
	// List of secondary servers. This makes the zone a
	// primary. Conflicts with `primary` and `additionalPrimaries`.
	// Secondaries is documented below.
	Secondaries []ZoneSecondary `pulumi:"secondaries"`
	// The SOA TTL.
	Ttl *int `pulumi:"ttl"`
	// The domain name of the zone.
	Zone string `pulumi:"zone"`
}

// The set of arguments for constructing a Zone resource.
type ZoneArgs struct {
	// List of additional IPv4 addresses for the primary
	// zone. Conflicts with `secondaries`.
	AdditionalPrimaries  pulumi.StringArrayInput
	AutogenerateNsRecord pulumi.BoolPtrInput
	// Whether or not DNSSEC is enabled for the zone.
	// Note that DNSSEC must be enabled on the account by support for this to be set
	// to `true`.
	Dnssec pulumi.BoolPtrInput
	// The SOA Expiry. Conflicts with `primary` and
	// `additionalPrimaries` (default must be accepted).
	Expiry pulumi.IntPtrInput
	// The target zone(domain name) to link to.
	Link pulumi.StringPtrInput
	// - List of network IDs (`int`) for which the zone
	//   should be made available. Default is network 0, the primary NSONE Global
	//   Network. Normally, you should not have to worry about this.
	Networks pulumi.IntArrayInput
	// The SOA NX TTL. Conflicts with `primary` and
	// `additionalPrimaries` (default must be accepted).
	NxTtl pulumi.IntPtrInput
	// The primary zones' IPv4 address. This makes the zone a
	// secondary. Conflicts with `secondaries`.
	Primary pulumi.StringPtrInput
	// The SOA Refresh. Conflicts with `primary` and
	// `additionalPrimaries` (default must be accepted).
	Refresh pulumi.IntPtrInput
	// The SOA Retry. Conflicts with `primary` and
	// `additionalPrimaries` (default must be accepted).
	Retry pulumi.IntPtrInput
	// List of secondary servers. This makes the zone a
	// primary. Conflicts with `primary` and `additionalPrimaries`.
	// Secondaries is documented below.
	Secondaries ZoneSecondaryArrayInput
	// The SOA TTL.
	Ttl pulumi.IntPtrInput
	// The domain name of the zone.
	Zone pulumi.StringInput
}

func (ZoneArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*zoneArgs)(nil)).Elem()
}
