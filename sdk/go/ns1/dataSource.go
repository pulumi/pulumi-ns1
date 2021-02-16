// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ns1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a NS1 Data Source resource. This can be used to create, modify, and delete data sources.
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
// 		_, err := ns1.NewDataSource(ctx, "example", &ns1.DataSourceArgs{
// 			Sourcetype: pulumi.String("nsone_v1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## NS1 Documentation
//
// [Datasource Api Doc](https://ns1.com/api#data-sources)
type DataSource struct {
	pulumi.CustomResourceState

	// The data source configuration, determined by its type,
	// matching the specification in `config` from /data/sourcetypes.
	Config pulumi.MapOutput `pulumi:"config"`
	// The free form name of the data source.
	Name pulumi.StringOutput `pulumi:"name"`
	// The data sources type, listed in API endpoint https://api.nsone.net/v1/data/sourcetypes.
	Sourcetype pulumi.StringOutput `pulumi:"sourcetype"`
}

// NewDataSource registers a new resource with the given unique name, arguments, and options.
func NewDataSource(ctx *pulumi.Context,
	name string, args *DataSourceArgs, opts ...pulumi.ResourceOption) (*DataSource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Sourcetype == nil {
		return nil, errors.New("invalid value for required argument 'Sourcetype'")
	}
	var resource DataSource
	err := ctx.RegisterResource("ns1:index/dataSource:DataSource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataSource gets an existing DataSource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataSource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataSourceState, opts ...pulumi.ResourceOption) (*DataSource, error) {
	var resource DataSource
	err := ctx.ReadResource("ns1:index/dataSource:DataSource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataSource resources.
type dataSourceState struct {
	// The data source configuration, determined by its type,
	// matching the specification in `config` from /data/sourcetypes.
	Config map[string]interface{} `pulumi:"config"`
	// The free form name of the data source.
	Name *string `pulumi:"name"`
	// The data sources type, listed in API endpoint https://api.nsone.net/v1/data/sourcetypes.
	Sourcetype *string `pulumi:"sourcetype"`
}

type DataSourceState struct {
	// The data source configuration, determined by its type,
	// matching the specification in `config` from /data/sourcetypes.
	Config pulumi.MapInput
	// The free form name of the data source.
	Name pulumi.StringPtrInput
	// The data sources type, listed in API endpoint https://api.nsone.net/v1/data/sourcetypes.
	Sourcetype pulumi.StringPtrInput
}

func (DataSourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataSourceState)(nil)).Elem()
}

type dataSourceArgs struct {
	// The data source configuration, determined by its type,
	// matching the specification in `config` from /data/sourcetypes.
	Config map[string]interface{} `pulumi:"config"`
	// The free form name of the data source.
	Name *string `pulumi:"name"`
	// The data sources type, listed in API endpoint https://api.nsone.net/v1/data/sourcetypes.
	Sourcetype string `pulumi:"sourcetype"`
}

// The set of arguments for constructing a DataSource resource.
type DataSourceArgs struct {
	// The data source configuration, determined by its type,
	// matching the specification in `config` from /data/sourcetypes.
	Config pulumi.MapInput
	// The free form name of the data source.
	Name pulumi.StringPtrInput
	// The data sources type, listed in API endpoint https://api.nsone.net/v1/data/sourcetypes.
	Sourcetype pulumi.StringInput
}

func (DataSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataSourceArgs)(nil)).Elem()
}

type DataSourceInput interface {
	pulumi.Input

	ToDataSourceOutput() DataSourceOutput
	ToDataSourceOutputWithContext(ctx context.Context) DataSourceOutput
}

func (*DataSource) ElementType() reflect.Type {
	return reflect.TypeOf((*DataSource)(nil))
}

func (i *DataSource) ToDataSourceOutput() DataSourceOutput {
	return i.ToDataSourceOutputWithContext(context.Background())
}

func (i *DataSource) ToDataSourceOutputWithContext(ctx context.Context) DataSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataSourceOutput)
}

func (i *DataSource) ToDataSourcePtrOutput() DataSourcePtrOutput {
	return i.ToDataSourcePtrOutputWithContext(context.Background())
}

func (i *DataSource) ToDataSourcePtrOutputWithContext(ctx context.Context) DataSourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataSourcePtrOutput)
}

type DataSourcePtrInput interface {
	pulumi.Input

	ToDataSourcePtrOutput() DataSourcePtrOutput
	ToDataSourcePtrOutputWithContext(ctx context.Context) DataSourcePtrOutput
}

type dataSourcePtrType DataSourceArgs

func (*dataSourcePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DataSource)(nil))
}

func (i *dataSourcePtrType) ToDataSourcePtrOutput() DataSourcePtrOutput {
	return i.ToDataSourcePtrOutputWithContext(context.Background())
}

func (i *dataSourcePtrType) ToDataSourcePtrOutputWithContext(ctx context.Context) DataSourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataSourcePtrOutput)
}

// DataSourceArrayInput is an input type that accepts DataSourceArray and DataSourceArrayOutput values.
// You can construct a concrete instance of `DataSourceArrayInput` via:
//
//          DataSourceArray{ DataSourceArgs{...} }
type DataSourceArrayInput interface {
	pulumi.Input

	ToDataSourceArrayOutput() DataSourceArrayOutput
	ToDataSourceArrayOutputWithContext(context.Context) DataSourceArrayOutput
}

type DataSourceArray []DataSourceInput

func (DataSourceArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*DataSource)(nil))
}

func (i DataSourceArray) ToDataSourceArrayOutput() DataSourceArrayOutput {
	return i.ToDataSourceArrayOutputWithContext(context.Background())
}

func (i DataSourceArray) ToDataSourceArrayOutputWithContext(ctx context.Context) DataSourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataSourceArrayOutput)
}

// DataSourceMapInput is an input type that accepts DataSourceMap and DataSourceMapOutput values.
// You can construct a concrete instance of `DataSourceMapInput` via:
//
//          DataSourceMap{ "key": DataSourceArgs{...} }
type DataSourceMapInput interface {
	pulumi.Input

	ToDataSourceMapOutput() DataSourceMapOutput
	ToDataSourceMapOutputWithContext(context.Context) DataSourceMapOutput
}

type DataSourceMap map[string]DataSourceInput

func (DataSourceMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*DataSource)(nil))
}

func (i DataSourceMap) ToDataSourceMapOutput() DataSourceMapOutput {
	return i.ToDataSourceMapOutputWithContext(context.Background())
}

func (i DataSourceMap) ToDataSourceMapOutputWithContext(ctx context.Context) DataSourceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataSourceMapOutput)
}

type DataSourceOutput struct {
	*pulumi.OutputState
}

func (DataSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataSource)(nil))
}

func (o DataSourceOutput) ToDataSourceOutput() DataSourceOutput {
	return o
}

func (o DataSourceOutput) ToDataSourceOutputWithContext(ctx context.Context) DataSourceOutput {
	return o
}

func (o DataSourceOutput) ToDataSourcePtrOutput() DataSourcePtrOutput {
	return o.ToDataSourcePtrOutputWithContext(context.Background())
}

func (o DataSourceOutput) ToDataSourcePtrOutputWithContext(ctx context.Context) DataSourcePtrOutput {
	return o.ApplyT(func(v DataSource) *DataSource {
		return &v
	}).(DataSourcePtrOutput)
}

type DataSourcePtrOutput struct {
	*pulumi.OutputState
}

func (DataSourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataSource)(nil))
}

func (o DataSourcePtrOutput) ToDataSourcePtrOutput() DataSourcePtrOutput {
	return o
}

func (o DataSourcePtrOutput) ToDataSourcePtrOutputWithContext(ctx context.Context) DataSourcePtrOutput {
	return o
}

type DataSourceArrayOutput struct{ *pulumi.OutputState }

func (DataSourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataSource)(nil))
}

func (o DataSourceArrayOutput) ToDataSourceArrayOutput() DataSourceArrayOutput {
	return o
}

func (o DataSourceArrayOutput) ToDataSourceArrayOutputWithContext(ctx context.Context) DataSourceArrayOutput {
	return o
}

func (o DataSourceArrayOutput) Index(i pulumi.IntInput) DataSourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataSource {
		return vs[0].([]DataSource)[vs[1].(int)]
	}).(DataSourceOutput)
}

type DataSourceMapOutput struct{ *pulumi.OutputState }

func (DataSourceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DataSource)(nil))
}

func (o DataSourceMapOutput) ToDataSourceMapOutput() DataSourceMapOutput {
	return o
}

func (o DataSourceMapOutput) ToDataSourceMapOutputWithContext(ctx context.Context) DataSourceMapOutput {
	return o
}

func (o DataSourceMapOutput) MapIndex(k pulumi.StringInput) DataSourceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DataSource {
		return vs[0].(map[string]DataSource)[vs[1].(string)]
	}).(DataSourceOutput)
}

func init() {
	pulumi.RegisterOutputType(DataSourceOutput{})
	pulumi.RegisterOutputType(DataSourcePtrOutput{})
	pulumi.RegisterOutputType(DataSourceArrayOutput{})
	pulumi.RegisterOutputType(DataSourceMapOutput{})
}
