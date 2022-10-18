// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ns1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a NS1 Data Feed resource. This can be used to create, modify, and delete data feeds.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-ns1/sdk/v2/go/ns1"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			example, err := ns1.NewDataSource(ctx, "example", &ns1.DataSourceArgs{
//				Sourcetype: pulumi.String("nsone_v1"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleMonitoring, err := ns1.NewDataSource(ctx, "exampleMonitoring", &ns1.DataSourceArgs{
//				Sourcetype: pulumi.String("nsone_monitoring"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ns1.NewDataFeed(ctx, "uswestFeed", &ns1.DataFeedArgs{
//				SourceId: example.ID(),
//				Config: pulumi.AnyMap{
//					"label": pulumi.Any("uswest"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ns1.NewDataFeed(ctx, "useastFeed", &ns1.DataFeedArgs{
//				SourceId: example.ID(),
//				Config: pulumi.AnyMap{
//					"label": pulumi.Any("useast"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ns1.NewDataFeed(ctx, "useastMonitorFeed", &ns1.DataFeedArgs{
//				SourceId: exampleMonitoring.ID(),
//				Config: pulumi.AnyMap{
//					"jobid": pulumi.Any(ns1_monitoringjob.Example_job.Id),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ## NS1 Documentation
//
// [Datafeed Api Doc](https://ns1.com/api#data-feeds)
type DataFeed struct {
	pulumi.CustomResourceState

	// The feeds configuration matching the specification in
	// `feedConfig` from /data/sourcetypes. `jobid` is required in the `config` for datafeeds connected to NS1 monitoring.
	Config pulumi.MapOutput `pulumi:"config"`
	// The free form name of the data feed.
	Name pulumi.StringOutput `pulumi:"name"`
	// The data source id that this feed is connected to.
	SourceId pulumi.StringOutput `pulumi:"sourceId"`
}

// NewDataFeed registers a new resource with the given unique name, arguments, and options.
func NewDataFeed(ctx *pulumi.Context,
	name string, args *DataFeedArgs, opts ...pulumi.ResourceOption) (*DataFeed, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SourceId == nil {
		return nil, errors.New("invalid value for required argument 'SourceId'")
	}
	var resource DataFeed
	err := ctx.RegisterResource("ns1:index/dataFeed:DataFeed", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataFeed gets an existing DataFeed resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataFeed(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataFeedState, opts ...pulumi.ResourceOption) (*DataFeed, error) {
	var resource DataFeed
	err := ctx.ReadResource("ns1:index/dataFeed:DataFeed", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataFeed resources.
type dataFeedState struct {
	// The feeds configuration matching the specification in
	// `feedConfig` from /data/sourcetypes. `jobid` is required in the `config` for datafeeds connected to NS1 monitoring.
	Config map[string]interface{} `pulumi:"config"`
	// The free form name of the data feed.
	Name *string `pulumi:"name"`
	// The data source id that this feed is connected to.
	SourceId *string `pulumi:"sourceId"`
}

type DataFeedState struct {
	// The feeds configuration matching the specification in
	// `feedConfig` from /data/sourcetypes. `jobid` is required in the `config` for datafeeds connected to NS1 monitoring.
	Config pulumi.MapInput
	// The free form name of the data feed.
	Name pulumi.StringPtrInput
	// The data source id that this feed is connected to.
	SourceId pulumi.StringPtrInput
}

func (DataFeedState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataFeedState)(nil)).Elem()
}

type dataFeedArgs struct {
	// The feeds configuration matching the specification in
	// `feedConfig` from /data/sourcetypes. `jobid` is required in the `config` for datafeeds connected to NS1 monitoring.
	Config map[string]interface{} `pulumi:"config"`
	// The free form name of the data feed.
	Name *string `pulumi:"name"`
	// The data source id that this feed is connected to.
	SourceId string `pulumi:"sourceId"`
}

// The set of arguments for constructing a DataFeed resource.
type DataFeedArgs struct {
	// The feeds configuration matching the specification in
	// `feedConfig` from /data/sourcetypes. `jobid` is required in the `config` for datafeeds connected to NS1 monitoring.
	Config pulumi.MapInput
	// The free form name of the data feed.
	Name pulumi.StringPtrInput
	// The data source id that this feed is connected to.
	SourceId pulumi.StringInput
}

func (DataFeedArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataFeedArgs)(nil)).Elem()
}

type DataFeedInput interface {
	pulumi.Input

	ToDataFeedOutput() DataFeedOutput
	ToDataFeedOutputWithContext(ctx context.Context) DataFeedOutput
}

func (*DataFeed) ElementType() reflect.Type {
	return reflect.TypeOf((**DataFeed)(nil)).Elem()
}

func (i *DataFeed) ToDataFeedOutput() DataFeedOutput {
	return i.ToDataFeedOutputWithContext(context.Background())
}

func (i *DataFeed) ToDataFeedOutputWithContext(ctx context.Context) DataFeedOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataFeedOutput)
}

// DataFeedArrayInput is an input type that accepts DataFeedArray and DataFeedArrayOutput values.
// You can construct a concrete instance of `DataFeedArrayInput` via:
//
//	DataFeedArray{ DataFeedArgs{...} }
type DataFeedArrayInput interface {
	pulumi.Input

	ToDataFeedArrayOutput() DataFeedArrayOutput
	ToDataFeedArrayOutputWithContext(context.Context) DataFeedArrayOutput
}

type DataFeedArray []DataFeedInput

func (DataFeedArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DataFeed)(nil)).Elem()
}

func (i DataFeedArray) ToDataFeedArrayOutput() DataFeedArrayOutput {
	return i.ToDataFeedArrayOutputWithContext(context.Background())
}

func (i DataFeedArray) ToDataFeedArrayOutputWithContext(ctx context.Context) DataFeedArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataFeedArrayOutput)
}

// DataFeedMapInput is an input type that accepts DataFeedMap and DataFeedMapOutput values.
// You can construct a concrete instance of `DataFeedMapInput` via:
//
//	DataFeedMap{ "key": DataFeedArgs{...} }
type DataFeedMapInput interface {
	pulumi.Input

	ToDataFeedMapOutput() DataFeedMapOutput
	ToDataFeedMapOutputWithContext(context.Context) DataFeedMapOutput
}

type DataFeedMap map[string]DataFeedInput

func (DataFeedMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DataFeed)(nil)).Elem()
}

func (i DataFeedMap) ToDataFeedMapOutput() DataFeedMapOutput {
	return i.ToDataFeedMapOutputWithContext(context.Background())
}

func (i DataFeedMap) ToDataFeedMapOutputWithContext(ctx context.Context) DataFeedMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataFeedMapOutput)
}

type DataFeedOutput struct{ *pulumi.OutputState }

func (DataFeedOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataFeed)(nil)).Elem()
}

func (o DataFeedOutput) ToDataFeedOutput() DataFeedOutput {
	return o
}

func (o DataFeedOutput) ToDataFeedOutputWithContext(ctx context.Context) DataFeedOutput {
	return o
}

type DataFeedArrayOutput struct{ *pulumi.OutputState }

func (DataFeedArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DataFeed)(nil)).Elem()
}

func (o DataFeedArrayOutput) ToDataFeedArrayOutput() DataFeedArrayOutput {
	return o
}

func (o DataFeedArrayOutput) ToDataFeedArrayOutputWithContext(ctx context.Context) DataFeedArrayOutput {
	return o
}

func (o DataFeedArrayOutput) Index(i pulumi.IntInput) DataFeedOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DataFeed {
		return vs[0].([]*DataFeed)[vs[1].(int)]
	}).(DataFeedOutput)
}

type DataFeedMapOutput struct{ *pulumi.OutputState }

func (DataFeedMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DataFeed)(nil)).Elem()
}

func (o DataFeedMapOutput) ToDataFeedMapOutput() DataFeedMapOutput {
	return o
}

func (o DataFeedMapOutput) ToDataFeedMapOutputWithContext(ctx context.Context) DataFeedMapOutput {
	return o
}

func (o DataFeedMapOutput) MapIndex(k pulumi.StringInput) DataFeedOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DataFeed {
		return vs[0].(map[string]*DataFeed)[vs[1].(string)]
	}).(DataFeedOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataFeedInput)(nil)).Elem(), &DataFeed{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataFeedArrayInput)(nil)).Elem(), DataFeedArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataFeedMapInput)(nil)).Elem(), DataFeedMap{})
	pulumi.RegisterOutputType(DataFeedOutput{})
	pulumi.RegisterOutputType(DataFeedArrayOutput{})
	pulumi.RegisterOutputType(DataFeedMapOutput{})
}
