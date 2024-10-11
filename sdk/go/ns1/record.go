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

// Provides a NS1 Record resource. This can be used to create, modify, and delete records.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"encoding/json"
//	"fmt"
//
//	"github.com/pulumi/pulumi-external/sdk/go/external"
//	"github.com/pulumi/pulumi-ns1/sdk/v3/go/ns1"
//	"github.com/pulumi/pulumi-std/sdk/go/std"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ns1.NewZone(ctx, "example", &ns1.ZoneArgs{
//				Zone: pulumi.String("terraform.example.io"),
//			})
//			if err != nil {
//				return err
//			}
//			ns1, err := ns1.NewDataSource(ctx, "ns1", &ns1.DataSourceArgs{
//				Name:       pulumi.String("ns1_source"),
//				Sourcetype: pulumi.String("nsone_v1"),
//			})
//			if err != nil {
//				return err
//			}
//			foo, err := ns1.NewDataFeed(ctx, "foo", &ns1.DataFeedArgs{
//				Name:     pulumi.String("foo_feed"),
//				SourceId: ns1.ID(),
//				Config: pulumi.StringMap{
//					"label": pulumi.String("foo"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			bar, err := ns1.NewDataFeed(ctx, "bar", &ns1.DataFeedArgs{
//				Name:     pulumi.String("bar_feed"),
//				SourceId: ns1.ID(),
//				Config: pulumi.StringMap{
//					"label": pulumi.String("bar"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			tmpJSON0, err := json.Marshal([]map[string]interface{}{
//				map[string]interface{}{
//					"job_id":     "abcdef",
//					"bias":       "*0.55",
//					"a5m_cutoff": 0.9,
//				},
//			})
//			if err != nil {
//				return err
//			}
//			json0 := string(tmpJSON0)
//			tmpJSON1, err := json.Marshal(map[string]interface{}{
//				"BR": []string{
//					"SP",
//					"SC",
//				},
//				"DZ": []string{
//					"01",
//					"02",
//					"03",
//				},
//			})
//			if err != nil {
//				return err
//			}
//			json1 := string(tmpJSON1)
//			_, err = ns1.NewRecord(ctx, "www", &ns1.RecordArgs{
//				Zone:   pulumi.Any(tld.Zone),
//				Domain: pulumi.Sprintf("www.%v", tld.Zone),
//				Type:   pulumi.String("CNAME"),
//				Ttl:    pulumi.Int(60),
//				Meta: pulumi.StringMap{
//					"up": pulumi.String("true"),
//				},
//				Regions: ns1.RecordRegionArray{
//					&ns1.RecordRegionArgs{
//						Name: pulumi.String("east"),
//						Meta: pulumi.StringMap{
//							"georegion": pulumi.String("US-EAST"),
//						},
//					},
//					&ns1.RecordRegionArgs{
//						Name: pulumi.String("usa"),
//						Meta: pulumi.StringMap{
//							"country": pulumi.String("US"),
//						},
//					},
//				},
//				Answers: ns1.RecordAnswerArray{
//					&ns1.RecordAnswerArgs{
//						Answer: pulumi.Sprintf("sub1.%v", tld.Zone),
//						Region: pulumi.String("east"),
//						Meta: pulumi.StringMap{
//							"up": foo.ID().ApplyT(func(id string) (string, error) {
//								return fmt.Sprintf("{\"feed\":\"%v\"}", id), nil
//							}).(pulumi.StringOutput),
//						},
//					},
//					&ns1.RecordAnswerArgs{
//						Answer: pulumi.Sprintf("sub2.%v", tld.Zone),
//						Meta: pulumi.StringMap{
//							"up": bar.ID().ApplyT(func(id string) (string, error) {
//								return fmt.Sprintf("{\"feed\":\"%v\"}", id), nil
//							}).(pulumi.StringOutput),
//							"connections": pulumi.String("3"),
//						},
//					},
//					&ns1.RecordAnswerArgs{
//						Answer: pulumi.Sprintf("sub3.%v", tld.Zone),
//						Meta: pulumi.StringMap{
//							"pulsar":       pulumi.String(json0),
//							"subdivisions": pulumi.String(json1),
//						},
//					},
//				},
//				Filters: ns1.RecordFilterArray{
//					&ns1.RecordFilterArgs{
//						Filter: pulumi.String("select_first_n"),
//						Config: pulumi.StringMap{
//							"N": pulumi.String("1"),
//						},
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			// Some other non-NS1 provider that returns a zone with a trailing dot and a domain with a leading dot.
//			_, err = external.NewSource(ctx, "baz", &external.SourceArgs{
//				Zone:   "terraform.example.io.",
//				Domain: ".www.terraform.example.io",
//			})
//			if err != nil {
//				return err
//			}
//			invokeReplace, err := std.Replace(ctx, &std.ReplaceArgs{
//				Text:    zone,
//				Search:  "/(^\\.)|(\\.$)/",
//				Replace: "",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			invokeReplace1, err := std.Replace(ctx, &std.ReplaceArgs{
//				Text:    domain,
//				Search:  "/(^\\.)|(\\.$)/",
//				Replace: "",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			// Basic record showing how to clean a zone or domain field that comes from
//			// another non-NS1 resource. DNS names often end in '.' characters to signify
//			// the root of the DNS tree, but the NS1 provider does not support this.
//			//
//			// In other cases, a domain or zone may be passed in with a preceding dot ('.')
//			// character which would likewise lead the system to fail.
//			_, err = ns1.NewRecord(ctx, "external", &ns1.RecordArgs{
//				Zone:   pulumi.String(invokeReplace.Result),
//				Domain: pulumi.String(invokeReplace1.Result),
//				Type:   pulumi.String("CNAME"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## NS1 Documentation
//
// [Record Api Doc](https://ns1.com/api#records)
//
// ## Import
//
// ```sh
// $ pulumi import ns1:index/record:Record <name> <zone>/<domain>/<type>`
// ```
//
// So for the example above:
//
// ```sh
// $ pulumi import ns1:index/record:Record www terraform.example.io/www.terraform.example.io/CNAME`
// ```
type Record struct {
	pulumi.CustomResourceState

	// One or more NS1 answers for the records' specified type.
	// Answers are documented below.
	Answers     RecordAnswerArrayOutput  `pulumi:"answers"`
	BlockedTags pulumi.StringArrayOutput `pulumi:"blockedTags"`
	// The records' domain. Cannot have leading or trailing
	// dots - see the example above and `FQDN formatting` below.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// One or more NS1 filters for the record(order matters).
	// Filters are documented below.
	Filters RecordFilterArrayOutput `pulumi:"filters"`
	// The target record to link to. This means this record is a
	// 'linked' record, and it inherits all properties from its target.
	Link                   pulumi.StringPtrOutput `pulumi:"link"`
	Meta                   pulumi.StringMapOutput `pulumi:"meta"`
	OverrideAddressRecords pulumi.BoolPtrOutput   `pulumi:"overrideAddressRecords"`
	OverrideTtl            pulumi.BoolPtrOutput   `pulumi:"overrideTtl"`
	// One or more "regions" for the record. These are really
	// just groupings based on metadata, and are called "Answer Groups" in the NS1 UI,
	// but remain `regions` here for legacy reasons. Regions are
	// documented below. Please note the ordering requirement!
	Regions RecordRegionArrayOutput `pulumi:"regions"`
	// Deprecated: short_answers will be deprecated in a future release. It is suggested to migrate to a regular "answers" block.
	ShortAnswers pulumi.StringArrayOutput `pulumi:"shortAnswers"`
	// map of tags in the form of `"key" = "value"` where both key and value are strings
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The records' time to live (in seconds).
	Ttl pulumi.IntOutput `pulumi:"ttl"`
	// The records' RR type.
	Type pulumi.StringOutput `pulumi:"type"`
	// Whether to use EDNS client subnet data when
	// available(in filter chain).
	// * `  meta ` - (Optional) meta is supported at the `record` level. Meta
	//   is documented below.
	UseClientSubnet pulumi.BoolPtrOutput `pulumi:"useClientSubnet"`
	// The zone the record belongs to. Cannot have leading or
	// trailing dots (".") - see the example above and `FQDN formatting` below.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewRecord registers a new resource with the given unique name, arguments, and options.
func NewRecord(ctx *pulumi.Context,
	name string, args *RecordArgs, opts ...pulumi.ResourceOption) (*Record, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Domain == nil {
		return nil, errors.New("invalid value for required argument 'Domain'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	if args.Zone == nil {
		return nil, errors.New("invalid value for required argument 'Zone'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Record
	err := ctx.RegisterResource("ns1:index/record:Record", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRecord gets an existing Record resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRecord(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RecordState, opts ...pulumi.ResourceOption) (*Record, error) {
	var resource Record
	err := ctx.ReadResource("ns1:index/record:Record", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Record resources.
type recordState struct {
	// One or more NS1 answers for the records' specified type.
	// Answers are documented below.
	Answers     []RecordAnswer `pulumi:"answers"`
	BlockedTags []string       `pulumi:"blockedTags"`
	// The records' domain. Cannot have leading or trailing
	// dots - see the example above and `FQDN formatting` below.
	Domain *string `pulumi:"domain"`
	// One or more NS1 filters for the record(order matters).
	// Filters are documented below.
	Filters []RecordFilter `pulumi:"filters"`
	// The target record to link to. This means this record is a
	// 'linked' record, and it inherits all properties from its target.
	Link                   *string           `pulumi:"link"`
	Meta                   map[string]string `pulumi:"meta"`
	OverrideAddressRecords *bool             `pulumi:"overrideAddressRecords"`
	OverrideTtl            *bool             `pulumi:"overrideTtl"`
	// One or more "regions" for the record. These are really
	// just groupings based on metadata, and are called "Answer Groups" in the NS1 UI,
	// but remain `regions` here for legacy reasons. Regions are
	// documented below. Please note the ordering requirement!
	Regions []RecordRegion `pulumi:"regions"`
	// Deprecated: short_answers will be deprecated in a future release. It is suggested to migrate to a regular "answers" block.
	ShortAnswers []string `pulumi:"shortAnswers"`
	// map of tags in the form of `"key" = "value"` where both key and value are strings
	Tags map[string]string `pulumi:"tags"`
	// The records' time to live (in seconds).
	Ttl *int `pulumi:"ttl"`
	// The records' RR type.
	Type *string `pulumi:"type"`
	// Whether to use EDNS client subnet data when
	// available(in filter chain).
	// * `  meta ` - (Optional) meta is supported at the `record` level. Meta
	//   is documented below.
	UseClientSubnet *bool `pulumi:"useClientSubnet"`
	// The zone the record belongs to. Cannot have leading or
	// trailing dots (".") - see the example above and `FQDN formatting` below.
	Zone *string `pulumi:"zone"`
}

type RecordState struct {
	// One or more NS1 answers for the records' specified type.
	// Answers are documented below.
	Answers     RecordAnswerArrayInput
	BlockedTags pulumi.StringArrayInput
	// The records' domain. Cannot have leading or trailing
	// dots - see the example above and `FQDN formatting` below.
	Domain pulumi.StringPtrInput
	// One or more NS1 filters for the record(order matters).
	// Filters are documented below.
	Filters RecordFilterArrayInput
	// The target record to link to. This means this record is a
	// 'linked' record, and it inherits all properties from its target.
	Link                   pulumi.StringPtrInput
	Meta                   pulumi.StringMapInput
	OverrideAddressRecords pulumi.BoolPtrInput
	OverrideTtl            pulumi.BoolPtrInput
	// One or more "regions" for the record. These are really
	// just groupings based on metadata, and are called "Answer Groups" in the NS1 UI,
	// but remain `regions` here for legacy reasons. Regions are
	// documented below. Please note the ordering requirement!
	Regions RecordRegionArrayInput
	// Deprecated: short_answers will be deprecated in a future release. It is suggested to migrate to a regular "answers" block.
	ShortAnswers pulumi.StringArrayInput
	// map of tags in the form of `"key" = "value"` where both key and value are strings
	Tags pulumi.StringMapInput
	// The records' time to live (in seconds).
	Ttl pulumi.IntPtrInput
	// The records' RR type.
	Type pulumi.StringPtrInput
	// Whether to use EDNS client subnet data when
	// available(in filter chain).
	// * `  meta ` - (Optional) meta is supported at the `record` level. Meta
	//   is documented below.
	UseClientSubnet pulumi.BoolPtrInput
	// The zone the record belongs to. Cannot have leading or
	// trailing dots (".") - see the example above and `FQDN formatting` below.
	Zone pulumi.StringPtrInput
}

func (RecordState) ElementType() reflect.Type {
	return reflect.TypeOf((*recordState)(nil)).Elem()
}

type recordArgs struct {
	// One or more NS1 answers for the records' specified type.
	// Answers are documented below.
	Answers     []RecordAnswer `pulumi:"answers"`
	BlockedTags []string       `pulumi:"blockedTags"`
	// The records' domain. Cannot have leading or trailing
	// dots - see the example above and `FQDN formatting` below.
	Domain string `pulumi:"domain"`
	// One or more NS1 filters for the record(order matters).
	// Filters are documented below.
	Filters []RecordFilter `pulumi:"filters"`
	// The target record to link to. This means this record is a
	// 'linked' record, and it inherits all properties from its target.
	Link                   *string           `pulumi:"link"`
	Meta                   map[string]string `pulumi:"meta"`
	OverrideAddressRecords *bool             `pulumi:"overrideAddressRecords"`
	OverrideTtl            *bool             `pulumi:"overrideTtl"`
	// One or more "regions" for the record. These are really
	// just groupings based on metadata, and are called "Answer Groups" in the NS1 UI,
	// but remain `regions` here for legacy reasons. Regions are
	// documented below. Please note the ordering requirement!
	Regions []RecordRegion `pulumi:"regions"`
	// Deprecated: short_answers will be deprecated in a future release. It is suggested to migrate to a regular "answers" block.
	ShortAnswers []string `pulumi:"shortAnswers"`
	// map of tags in the form of `"key" = "value"` where both key and value are strings
	Tags map[string]string `pulumi:"tags"`
	// The records' time to live (in seconds).
	Ttl *int `pulumi:"ttl"`
	// The records' RR type.
	Type string `pulumi:"type"`
	// Whether to use EDNS client subnet data when
	// available(in filter chain).
	// * `  meta ` - (Optional) meta is supported at the `record` level. Meta
	//   is documented below.
	UseClientSubnet *bool `pulumi:"useClientSubnet"`
	// The zone the record belongs to. Cannot have leading or
	// trailing dots (".") - see the example above and `FQDN formatting` below.
	Zone string `pulumi:"zone"`
}

// The set of arguments for constructing a Record resource.
type RecordArgs struct {
	// One or more NS1 answers for the records' specified type.
	// Answers are documented below.
	Answers     RecordAnswerArrayInput
	BlockedTags pulumi.StringArrayInput
	// The records' domain. Cannot have leading or trailing
	// dots - see the example above and `FQDN formatting` below.
	Domain pulumi.StringInput
	// One or more NS1 filters for the record(order matters).
	// Filters are documented below.
	Filters RecordFilterArrayInput
	// The target record to link to. This means this record is a
	// 'linked' record, and it inherits all properties from its target.
	Link                   pulumi.StringPtrInput
	Meta                   pulumi.StringMapInput
	OverrideAddressRecords pulumi.BoolPtrInput
	OverrideTtl            pulumi.BoolPtrInput
	// One or more "regions" for the record. These are really
	// just groupings based on metadata, and are called "Answer Groups" in the NS1 UI,
	// but remain `regions` here for legacy reasons. Regions are
	// documented below. Please note the ordering requirement!
	Regions RecordRegionArrayInput
	// Deprecated: short_answers will be deprecated in a future release. It is suggested to migrate to a regular "answers" block.
	ShortAnswers pulumi.StringArrayInput
	// map of tags in the form of `"key" = "value"` where both key and value are strings
	Tags pulumi.StringMapInput
	// The records' time to live (in seconds).
	Ttl pulumi.IntPtrInput
	// The records' RR type.
	Type pulumi.StringInput
	// Whether to use EDNS client subnet data when
	// available(in filter chain).
	// * `  meta ` - (Optional) meta is supported at the `record` level. Meta
	//   is documented below.
	UseClientSubnet pulumi.BoolPtrInput
	// The zone the record belongs to. Cannot have leading or
	// trailing dots (".") - see the example above and `FQDN formatting` below.
	Zone pulumi.StringInput
}

func (RecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*recordArgs)(nil)).Elem()
}

type RecordInput interface {
	pulumi.Input

	ToRecordOutput() RecordOutput
	ToRecordOutputWithContext(ctx context.Context) RecordOutput
}

func (*Record) ElementType() reflect.Type {
	return reflect.TypeOf((**Record)(nil)).Elem()
}

func (i *Record) ToRecordOutput() RecordOutput {
	return i.ToRecordOutputWithContext(context.Background())
}

func (i *Record) ToRecordOutputWithContext(ctx context.Context) RecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecordOutput)
}

// RecordArrayInput is an input type that accepts RecordArray and RecordArrayOutput values.
// You can construct a concrete instance of `RecordArrayInput` via:
//
//	RecordArray{ RecordArgs{...} }
type RecordArrayInput interface {
	pulumi.Input

	ToRecordArrayOutput() RecordArrayOutput
	ToRecordArrayOutputWithContext(context.Context) RecordArrayOutput
}

type RecordArray []RecordInput

func (RecordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Record)(nil)).Elem()
}

func (i RecordArray) ToRecordArrayOutput() RecordArrayOutput {
	return i.ToRecordArrayOutputWithContext(context.Background())
}

func (i RecordArray) ToRecordArrayOutputWithContext(ctx context.Context) RecordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecordArrayOutput)
}

// RecordMapInput is an input type that accepts RecordMap and RecordMapOutput values.
// You can construct a concrete instance of `RecordMapInput` via:
//
//	RecordMap{ "key": RecordArgs{...} }
type RecordMapInput interface {
	pulumi.Input

	ToRecordMapOutput() RecordMapOutput
	ToRecordMapOutputWithContext(context.Context) RecordMapOutput
}

type RecordMap map[string]RecordInput

func (RecordMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Record)(nil)).Elem()
}

func (i RecordMap) ToRecordMapOutput() RecordMapOutput {
	return i.ToRecordMapOutputWithContext(context.Background())
}

func (i RecordMap) ToRecordMapOutputWithContext(ctx context.Context) RecordMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecordMapOutput)
}

type RecordOutput struct{ *pulumi.OutputState }

func (RecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Record)(nil)).Elem()
}

func (o RecordOutput) ToRecordOutput() RecordOutput {
	return o
}

func (o RecordOutput) ToRecordOutputWithContext(ctx context.Context) RecordOutput {
	return o
}

// One or more NS1 answers for the records' specified type.
// Answers are documented below.
func (o RecordOutput) Answers() RecordAnswerArrayOutput {
	return o.ApplyT(func(v *Record) RecordAnswerArrayOutput { return v.Answers }).(RecordAnswerArrayOutput)
}

func (o RecordOutput) BlockedTags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Record) pulumi.StringArrayOutput { return v.BlockedTags }).(pulumi.StringArrayOutput)
}

// The records' domain. Cannot have leading or trailing
// dots - see the example above and `FQDN formatting` below.
func (o RecordOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v *Record) pulumi.StringOutput { return v.Domain }).(pulumi.StringOutput)
}

// One or more NS1 filters for the record(order matters).
// Filters are documented below.
func (o RecordOutput) Filters() RecordFilterArrayOutput {
	return o.ApplyT(func(v *Record) RecordFilterArrayOutput { return v.Filters }).(RecordFilterArrayOutput)
}

// The target record to link to. This means this record is a
// 'linked' record, and it inherits all properties from its target.
func (o RecordOutput) Link() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Record) pulumi.StringPtrOutput { return v.Link }).(pulumi.StringPtrOutput)
}

func (o RecordOutput) Meta() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Record) pulumi.StringMapOutput { return v.Meta }).(pulumi.StringMapOutput)
}

func (o RecordOutput) OverrideAddressRecords() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Record) pulumi.BoolPtrOutput { return v.OverrideAddressRecords }).(pulumi.BoolPtrOutput)
}

func (o RecordOutput) OverrideTtl() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Record) pulumi.BoolPtrOutput { return v.OverrideTtl }).(pulumi.BoolPtrOutput)
}

// One or more "regions" for the record. These are really
// just groupings based on metadata, and are called "Answer Groups" in the NS1 UI,
// but remain `regions` here for legacy reasons. Regions are
// documented below. Please note the ordering requirement!
func (o RecordOutput) Regions() RecordRegionArrayOutput {
	return o.ApplyT(func(v *Record) RecordRegionArrayOutput { return v.Regions }).(RecordRegionArrayOutput)
}

// Deprecated: short_answers will be deprecated in a future release. It is suggested to migrate to a regular "answers" block.
func (o RecordOutput) ShortAnswers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Record) pulumi.StringArrayOutput { return v.ShortAnswers }).(pulumi.StringArrayOutput)
}

// map of tags in the form of `"key" = "value"` where both key and value are strings
func (o RecordOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Record) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The records' time to live (in seconds).
func (o RecordOutput) Ttl() pulumi.IntOutput {
	return o.ApplyT(func(v *Record) pulumi.IntOutput { return v.Ttl }).(pulumi.IntOutput)
}

// The records' RR type.
func (o RecordOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Record) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Whether to use EDNS client subnet data when
// available(in filter chain).
//   - `  meta ` - (Optional) meta is supported at the `record` level. Meta
//     is documented below.
func (o RecordOutput) UseClientSubnet() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Record) pulumi.BoolPtrOutput { return v.UseClientSubnet }).(pulumi.BoolPtrOutput)
}

// The zone the record belongs to. Cannot have leading or
// trailing dots (".") - see the example above and `FQDN formatting` below.
func (o RecordOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *Record) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type RecordArrayOutput struct{ *pulumi.OutputState }

func (RecordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Record)(nil)).Elem()
}

func (o RecordArrayOutput) ToRecordArrayOutput() RecordArrayOutput {
	return o
}

func (o RecordArrayOutput) ToRecordArrayOutputWithContext(ctx context.Context) RecordArrayOutput {
	return o
}

func (o RecordArrayOutput) Index(i pulumi.IntInput) RecordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Record {
		return vs[0].([]*Record)[vs[1].(int)]
	}).(RecordOutput)
}

type RecordMapOutput struct{ *pulumi.OutputState }

func (RecordMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Record)(nil)).Elem()
}

func (o RecordMapOutput) ToRecordMapOutput() RecordMapOutput {
	return o
}

func (o RecordMapOutput) ToRecordMapOutputWithContext(ctx context.Context) RecordMapOutput {
	return o
}

func (o RecordMapOutput) MapIndex(k pulumi.StringInput) RecordOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Record {
		return vs[0].(map[string]*Record)[vs[1].(string)]
	}).(RecordOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RecordInput)(nil)).Elem(), &Record{})
	pulumi.RegisterInputType(reflect.TypeOf((*RecordArrayInput)(nil)).Elem(), RecordArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RecordMapInput)(nil)).Elem(), RecordMap{})
	pulumi.RegisterOutputType(RecordOutput{})
	pulumi.RegisterOutputType(RecordArrayOutput{})
	pulumi.RegisterOutputType(RecordMapOutput{})
}
