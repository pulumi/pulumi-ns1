// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ns1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a NS1 Record resource. This can be used to create, modify, and delete records.
//
// ## NS1 Documentation
//
// [Record Api Doc](https://ns1.com/api#records)
//
// ## Import
//
// ```sh
//  $ pulumi import ns1:index/record:Record <name> <zone>/<domain>/<type>`
// ```
//
//  So for the example above
//
// ```sh
//  $ pulumi import ns1:index/record:Record www terraform.example.io/www.terraform.example.io/CNAME`
// ```
type Record struct {
	pulumi.CustomResourceState

	// One or more NS1 answers for the records' specified type.
	// Answers are documented below.
	Answers RecordAnswerArrayOutput `pulumi:"answers"`
	// The records' domain. Cannot have leading or trailing
	// dots - see the example above and `FQDN formatting` below.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// One or more NS1 filters for the record(order matters).
	// Filters are documented below.
	Filters RecordFilterArrayOutput `pulumi:"filters"`
	// The target record to link to. This means this record is a
	// 'linked' record, and it inherits all properties from its target.
	Link        pulumi.StringPtrOutput `pulumi:"link"`
	Meta        pulumi.MapOutput       `pulumi:"meta"`
	OverrideTtl pulumi.BoolPtrOutput   `pulumi:"overrideTtl"`
	// One or more "regions" for the record. These are really
	// just groupings based on metadata, and are called "Answer Groups" in the NS1 UI,
	// but remain `regions` here for legacy reasons. Regions are
	// documented below. Please note the ordering requirement!
	Regions RecordRegionArrayOutput `pulumi:"regions"`
	// Deprecated: short_answers will be deprecated in a future release. It is suggested to migrate to a regular "answers" block.
	ShortAnswers pulumi.StringArrayOutput `pulumi:"shortAnswers"`
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
	Answers []RecordAnswer `pulumi:"answers"`
	// The records' domain. Cannot have leading or trailing
	// dots - see the example above and `FQDN formatting` below.
	Domain *string `pulumi:"domain"`
	// One or more NS1 filters for the record(order matters).
	// Filters are documented below.
	Filters []RecordFilter `pulumi:"filters"`
	// The target record to link to. This means this record is a
	// 'linked' record, and it inherits all properties from its target.
	Link        *string                `pulumi:"link"`
	Meta        map[string]interface{} `pulumi:"meta"`
	OverrideTtl *bool                  `pulumi:"overrideTtl"`
	// One or more "regions" for the record. These are really
	// just groupings based on metadata, and are called "Answer Groups" in the NS1 UI,
	// but remain `regions` here for legacy reasons. Regions are
	// documented below. Please note the ordering requirement!
	Regions []RecordRegion `pulumi:"regions"`
	// Deprecated: short_answers will be deprecated in a future release. It is suggested to migrate to a regular "answers" block.
	ShortAnswers []string `pulumi:"shortAnswers"`
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
	Answers RecordAnswerArrayInput
	// The records' domain. Cannot have leading or trailing
	// dots - see the example above and `FQDN formatting` below.
	Domain pulumi.StringPtrInput
	// One or more NS1 filters for the record(order matters).
	// Filters are documented below.
	Filters RecordFilterArrayInput
	// The target record to link to. This means this record is a
	// 'linked' record, and it inherits all properties from its target.
	Link        pulumi.StringPtrInput
	Meta        pulumi.MapInput
	OverrideTtl pulumi.BoolPtrInput
	// One or more "regions" for the record. These are really
	// just groupings based on metadata, and are called "Answer Groups" in the NS1 UI,
	// but remain `regions` here for legacy reasons. Regions are
	// documented below. Please note the ordering requirement!
	Regions RecordRegionArrayInput
	// Deprecated: short_answers will be deprecated in a future release. It is suggested to migrate to a regular "answers" block.
	ShortAnswers pulumi.StringArrayInput
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
	Answers []RecordAnswer `pulumi:"answers"`
	// The records' domain. Cannot have leading or trailing
	// dots - see the example above and `FQDN formatting` below.
	Domain string `pulumi:"domain"`
	// One or more NS1 filters for the record(order matters).
	// Filters are documented below.
	Filters []RecordFilter `pulumi:"filters"`
	// The target record to link to. This means this record is a
	// 'linked' record, and it inherits all properties from its target.
	Link        *string                `pulumi:"link"`
	Meta        map[string]interface{} `pulumi:"meta"`
	OverrideTtl *bool                  `pulumi:"overrideTtl"`
	// One or more "regions" for the record. These are really
	// just groupings based on metadata, and are called "Answer Groups" in the NS1 UI,
	// but remain `regions` here for legacy reasons. Regions are
	// documented below. Please note the ordering requirement!
	Regions []RecordRegion `pulumi:"regions"`
	// Deprecated: short_answers will be deprecated in a future release. It is suggested to migrate to a regular "answers" block.
	ShortAnswers []string `pulumi:"shortAnswers"`
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
	Answers RecordAnswerArrayInput
	// The records' domain. Cannot have leading or trailing
	// dots - see the example above and `FQDN formatting` below.
	Domain pulumi.StringInput
	// One or more NS1 filters for the record(order matters).
	// Filters are documented below.
	Filters RecordFilterArrayInput
	// The target record to link to. This means this record is a
	// 'linked' record, and it inherits all properties from its target.
	Link        pulumi.StringPtrInput
	Meta        pulumi.MapInput
	OverrideTtl pulumi.BoolPtrInput
	// One or more "regions" for the record. These are really
	// just groupings based on metadata, and are called "Answer Groups" in the NS1 UI,
	// but remain `regions` here for legacy reasons. Regions are
	// documented below. Please note the ordering requirement!
	Regions RecordRegionArrayInput
	// Deprecated: short_answers will be deprecated in a future release. It is suggested to migrate to a regular "answers" block.
	ShortAnswers pulumi.StringArrayInput
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
//          RecordArray{ RecordArgs{...} }
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
//          RecordMap{ "key": RecordArgs{...} }
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
