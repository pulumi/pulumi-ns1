// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ns1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-ns1/sdk/v3/go/ns1/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// ```sh
//
//	$ pulumi import ns1:index/zone:Zone <name> <zone>`
//
// ```
//
//	So for the example above
//
// ```sh
//
//	$ pulumi import ns1:index/zone:Zone example terraform.example.io`
//
// ```
type Zone struct {
	pulumi.CustomResourceState

	AdditionalPorts pulumi.IntArrayOutput `pulumi:"additionalPorts"`
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
	// List of network IDs for which the zone is
	// available. If no network is provided, the zone will be created in network 0,
	// the primary NS1 Global Network.
	Networks pulumi.IntArrayOutput `pulumi:"networks"`
	// The SOA NX TTL. Conflicts with `primary` and
	// `additionalPrimaries` (default must be accepted).
	NxTtl pulumi.IntOutput `pulumi:"nxTtl"`
	// The primary zones' IPv4 address. This makes the zone a
	// secondary. Conflicts with `secondaries`.
	Primary     pulumi.StringPtrOutput `pulumi:"primary"`
	PrimaryPort pulumi.IntOutput       `pulumi:"primaryPort"`
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
	// TSIG is documented below
	Tsig pulumi.StringMapOutput `pulumi:"tsig"`
	// The SOA TTL.
	Ttl pulumi.IntOutput `pulumi:"ttl"`
	// The domain name of the zone.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewZone registers a new resource with the given unique name, arguments, and options.
func NewZone(ctx *pulumi.Context,
	name string, args *ZoneArgs, opts ...pulumi.ResourceOption) (*Zone, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Zone == nil {
		return nil, errors.New("invalid value for required argument 'Zone'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
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
	AdditionalPorts []int `pulumi:"additionalPorts"`
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
	// List of network IDs for which the zone is
	// available. If no network is provided, the zone will be created in network 0,
	// the primary NS1 Global Network.
	Networks []int `pulumi:"networks"`
	// The SOA NX TTL. Conflicts with `primary` and
	// `additionalPrimaries` (default must be accepted).
	NxTtl *int `pulumi:"nxTtl"`
	// The primary zones' IPv4 address. This makes the zone a
	// secondary. Conflicts with `secondaries`.
	Primary     *string `pulumi:"primary"`
	PrimaryPort *int    `pulumi:"primaryPort"`
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
	// TSIG is documented below
	Tsig map[string]string `pulumi:"tsig"`
	// The SOA TTL.
	Ttl *int `pulumi:"ttl"`
	// The domain name of the zone.
	Zone *string `pulumi:"zone"`
}

type ZoneState struct {
	AdditionalPorts pulumi.IntArrayInput
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
	// List of network IDs for which the zone is
	// available. If no network is provided, the zone will be created in network 0,
	// the primary NS1 Global Network.
	Networks pulumi.IntArrayInput
	// The SOA NX TTL. Conflicts with `primary` and
	// `additionalPrimaries` (default must be accepted).
	NxTtl pulumi.IntPtrInput
	// The primary zones' IPv4 address. This makes the zone a
	// secondary. Conflicts with `secondaries`.
	Primary     pulumi.StringPtrInput
	PrimaryPort pulumi.IntPtrInput
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
	// TSIG is documented below
	Tsig pulumi.StringMapInput
	// The SOA TTL.
	Ttl pulumi.IntPtrInput
	// The domain name of the zone.
	Zone pulumi.StringPtrInput
}

func (ZoneState) ElementType() reflect.Type {
	return reflect.TypeOf((*zoneState)(nil)).Elem()
}

type zoneArgs struct {
	AdditionalPorts []int `pulumi:"additionalPorts"`
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
	// (Computed) The SOA Hostmaster.
	Hostmaster *string `pulumi:"hostmaster"`
	// The target zone(domain name) to link to.
	Link *string `pulumi:"link"`
	// List of network IDs for which the zone is
	// available. If no network is provided, the zone will be created in network 0,
	// the primary NS1 Global Network.
	Networks []int `pulumi:"networks"`
	// The SOA NX TTL. Conflicts with `primary` and
	// `additionalPrimaries` (default must be accepted).
	NxTtl *int `pulumi:"nxTtl"`
	// The primary zones' IPv4 address. This makes the zone a
	// secondary. Conflicts with `secondaries`.
	Primary     *string `pulumi:"primary"`
	PrimaryPort *int    `pulumi:"primaryPort"`
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
	// TSIG is documented below
	Tsig map[string]string `pulumi:"tsig"`
	// The SOA TTL.
	Ttl *int `pulumi:"ttl"`
	// The domain name of the zone.
	Zone string `pulumi:"zone"`
}

// The set of arguments for constructing a Zone resource.
type ZoneArgs struct {
	AdditionalPorts pulumi.IntArrayInput
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
	// (Computed) The SOA Hostmaster.
	Hostmaster pulumi.StringPtrInput
	// The target zone(domain name) to link to.
	Link pulumi.StringPtrInput
	// List of network IDs for which the zone is
	// available. If no network is provided, the zone will be created in network 0,
	// the primary NS1 Global Network.
	Networks pulumi.IntArrayInput
	// The SOA NX TTL. Conflicts with `primary` and
	// `additionalPrimaries` (default must be accepted).
	NxTtl pulumi.IntPtrInput
	// The primary zones' IPv4 address. This makes the zone a
	// secondary. Conflicts with `secondaries`.
	Primary     pulumi.StringPtrInput
	PrimaryPort pulumi.IntPtrInput
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
	// TSIG is documented below
	Tsig pulumi.StringMapInput
	// The SOA TTL.
	Ttl pulumi.IntPtrInput
	// The domain name of the zone.
	Zone pulumi.StringInput
}

func (ZoneArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*zoneArgs)(nil)).Elem()
}

type ZoneInput interface {
	pulumi.Input

	ToZoneOutput() ZoneOutput
	ToZoneOutputWithContext(ctx context.Context) ZoneOutput
}

func (*Zone) ElementType() reflect.Type {
	return reflect.TypeOf((**Zone)(nil)).Elem()
}

func (i *Zone) ToZoneOutput() ZoneOutput {
	return i.ToZoneOutputWithContext(context.Background())
}

func (i *Zone) ToZoneOutputWithContext(ctx context.Context) ZoneOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneOutput)
}

// ZoneArrayInput is an input type that accepts ZoneArray and ZoneArrayOutput values.
// You can construct a concrete instance of `ZoneArrayInput` via:
//
//	ZoneArray{ ZoneArgs{...} }
type ZoneArrayInput interface {
	pulumi.Input

	ToZoneArrayOutput() ZoneArrayOutput
	ToZoneArrayOutputWithContext(context.Context) ZoneArrayOutput
}

type ZoneArray []ZoneInput

func (ZoneArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Zone)(nil)).Elem()
}

func (i ZoneArray) ToZoneArrayOutput() ZoneArrayOutput {
	return i.ToZoneArrayOutputWithContext(context.Background())
}

func (i ZoneArray) ToZoneArrayOutputWithContext(ctx context.Context) ZoneArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneArrayOutput)
}

// ZoneMapInput is an input type that accepts ZoneMap and ZoneMapOutput values.
// You can construct a concrete instance of `ZoneMapInput` via:
//
//	ZoneMap{ "key": ZoneArgs{...} }
type ZoneMapInput interface {
	pulumi.Input

	ToZoneMapOutput() ZoneMapOutput
	ToZoneMapOutputWithContext(context.Context) ZoneMapOutput
}

type ZoneMap map[string]ZoneInput

func (ZoneMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Zone)(nil)).Elem()
}

func (i ZoneMap) ToZoneMapOutput() ZoneMapOutput {
	return i.ToZoneMapOutputWithContext(context.Background())
}

func (i ZoneMap) ToZoneMapOutputWithContext(ctx context.Context) ZoneMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneMapOutput)
}

type ZoneOutput struct{ *pulumi.OutputState }

func (ZoneOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Zone)(nil)).Elem()
}

func (o ZoneOutput) ToZoneOutput() ZoneOutput {
	return o
}

func (o ZoneOutput) ToZoneOutputWithContext(ctx context.Context) ZoneOutput {
	return o
}

func (o ZoneOutput) AdditionalPorts() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *Zone) pulumi.IntArrayOutput { return v.AdditionalPorts }).(pulumi.IntArrayOutput)
}

// List of additional IPv4 addresses for the primary
// zone. Conflicts with `secondaries`.
func (o ZoneOutput) AdditionalPrimaries() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Zone) pulumi.StringArrayOutput { return v.AdditionalPrimaries }).(pulumi.StringArrayOutput)
}

func (o ZoneOutput) AutogenerateNsRecord() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Zone) pulumi.BoolPtrOutput { return v.AutogenerateNsRecord }).(pulumi.BoolPtrOutput)
}

// (Computed) Authoritative Name Servers.
func (o ZoneOutput) DnsServers() pulumi.StringOutput {
	return o.ApplyT(func(v *Zone) pulumi.StringOutput { return v.DnsServers }).(pulumi.StringOutput)
}

// Whether or not DNSSEC is enabled for the zone.
// Note that DNSSEC must be enabled on the account by support for this to be set
// to `true`.
func (o ZoneOutput) Dnssec() pulumi.BoolOutput {
	return o.ApplyT(func(v *Zone) pulumi.BoolOutput { return v.Dnssec }).(pulumi.BoolOutput)
}

// The SOA Expiry. Conflicts with `primary` and
// `additionalPrimaries` (default must be accepted).
func (o ZoneOutput) Expiry() pulumi.IntOutput {
	return o.ApplyT(func(v *Zone) pulumi.IntOutput { return v.Expiry }).(pulumi.IntOutput)
}

// (Computed) The SOA Hostmaster.
func (o ZoneOutput) Hostmaster() pulumi.StringOutput {
	return o.ApplyT(func(v *Zone) pulumi.StringOutput { return v.Hostmaster }).(pulumi.StringOutput)
}

// The target zone(domain name) to link to.
func (o ZoneOutput) Link() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Zone) pulumi.StringPtrOutput { return v.Link }).(pulumi.StringPtrOutput)
}

// List of network IDs for which the zone is
// available. If no network is provided, the zone will be created in network 0,
// the primary NS1 Global Network.
func (o ZoneOutput) Networks() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *Zone) pulumi.IntArrayOutput { return v.Networks }).(pulumi.IntArrayOutput)
}

// The SOA NX TTL. Conflicts with `primary` and
// `additionalPrimaries` (default must be accepted).
func (o ZoneOutput) NxTtl() pulumi.IntOutput {
	return o.ApplyT(func(v *Zone) pulumi.IntOutput { return v.NxTtl }).(pulumi.IntOutput)
}

// The primary zones' IPv4 address. This makes the zone a
// secondary. Conflicts with `secondaries`.
func (o ZoneOutput) Primary() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Zone) pulumi.StringPtrOutput { return v.Primary }).(pulumi.StringPtrOutput)
}

func (o ZoneOutput) PrimaryPort() pulumi.IntOutput {
	return o.ApplyT(func(v *Zone) pulumi.IntOutput { return v.PrimaryPort }).(pulumi.IntOutput)
}

// The SOA Refresh. Conflicts with `primary` and
// `additionalPrimaries` (default must be accepted).
func (o ZoneOutput) Refresh() pulumi.IntOutput {
	return o.ApplyT(func(v *Zone) pulumi.IntOutput { return v.Refresh }).(pulumi.IntOutput)
}

// The SOA Retry. Conflicts with `primary` and
// `additionalPrimaries` (default must be accepted).
func (o ZoneOutput) Retry() pulumi.IntOutput {
	return o.ApplyT(func(v *Zone) pulumi.IntOutput { return v.Retry }).(pulumi.IntOutput)
}

// List of secondary servers. This makes the zone a
// primary. Conflicts with `primary` and `additionalPrimaries`.
// Secondaries is documented below.
func (o ZoneOutput) Secondaries() ZoneSecondaryArrayOutput {
	return o.ApplyT(func(v *Zone) ZoneSecondaryArrayOutput { return v.Secondaries }).(ZoneSecondaryArrayOutput)
}

// TSIG is documented below
func (o ZoneOutput) Tsig() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Zone) pulumi.StringMapOutput { return v.Tsig }).(pulumi.StringMapOutput)
}

// The SOA TTL.
func (o ZoneOutput) Ttl() pulumi.IntOutput {
	return o.ApplyT(func(v *Zone) pulumi.IntOutput { return v.Ttl }).(pulumi.IntOutput)
}

// The domain name of the zone.
func (o ZoneOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *Zone) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type ZoneArrayOutput struct{ *pulumi.OutputState }

func (ZoneArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Zone)(nil)).Elem()
}

func (o ZoneArrayOutput) ToZoneArrayOutput() ZoneArrayOutput {
	return o
}

func (o ZoneArrayOutput) ToZoneArrayOutputWithContext(ctx context.Context) ZoneArrayOutput {
	return o
}

func (o ZoneArrayOutput) Index(i pulumi.IntInput) ZoneOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Zone {
		return vs[0].([]*Zone)[vs[1].(int)]
	}).(ZoneOutput)
}

type ZoneMapOutput struct{ *pulumi.OutputState }

func (ZoneMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Zone)(nil)).Elem()
}

func (o ZoneMapOutput) ToZoneMapOutput() ZoneMapOutput {
	return o
}

func (o ZoneMapOutput) ToZoneMapOutputWithContext(ctx context.Context) ZoneMapOutput {
	return o
}

func (o ZoneMapOutput) MapIndex(k pulumi.StringInput) ZoneOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Zone {
		return vs[0].(map[string]*Zone)[vs[1].(string)]
	}).(ZoneOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ZoneInput)(nil)).Elem(), &Zone{})
	pulumi.RegisterInputType(reflect.TypeOf((*ZoneArrayInput)(nil)).Elem(), ZoneArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ZoneMapInput)(nil)).Elem(), ZoneMap{})
	pulumi.RegisterOutputType(ZoneOutput{})
	pulumi.RegisterOutputType(ZoneArrayOutput{})
	pulumi.RegisterOutputType(ZoneMapOutput{})
}
