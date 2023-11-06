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

// Provides a NS1 Data Source resource. This can be used to create, modify, and delete data sources.
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
//			_, err := ns1.NewDataSource(ctx, "example", &ns1.DataSourceArgs{
//				Sourcetype: pulumi.String("nsone_v1"),
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
// [Datasource Api Doc](https://ns1.com/api#data-sources)
//
// ## Import
//
// ```sh
//
//	$ pulumi import ns1:index/dataSource:DataSource <name> <datasource_id>`
//
// ```
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
	opts = internal.PkgResourceDefaultOpts(opts)
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
	return reflect.TypeOf((**DataSource)(nil)).Elem()
}

func (i *DataSource) ToDataSourceOutput() DataSourceOutput {
	return i.ToDataSourceOutputWithContext(context.Background())
}

func (i *DataSource) ToDataSourceOutputWithContext(ctx context.Context) DataSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataSourceOutput)
}

// DataSourceArrayInput is an input type that accepts DataSourceArray and DataSourceArrayOutput values.
// You can construct a concrete instance of `DataSourceArrayInput` via:
//
//	DataSourceArray{ DataSourceArgs{...} }
type DataSourceArrayInput interface {
	pulumi.Input

	ToDataSourceArrayOutput() DataSourceArrayOutput
	ToDataSourceArrayOutputWithContext(context.Context) DataSourceArrayOutput
}

type DataSourceArray []DataSourceInput

func (DataSourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DataSource)(nil)).Elem()
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
//	DataSourceMap{ "key": DataSourceArgs{...} }
type DataSourceMapInput interface {
	pulumi.Input

	ToDataSourceMapOutput() DataSourceMapOutput
	ToDataSourceMapOutputWithContext(context.Context) DataSourceMapOutput
}

type DataSourceMap map[string]DataSourceInput

func (DataSourceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DataSource)(nil)).Elem()
}

func (i DataSourceMap) ToDataSourceMapOutput() DataSourceMapOutput {
	return i.ToDataSourceMapOutputWithContext(context.Background())
}

func (i DataSourceMap) ToDataSourceMapOutputWithContext(ctx context.Context) DataSourceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataSourceMapOutput)
}

type DataSourceOutput struct{ *pulumi.OutputState }

func (DataSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataSource)(nil)).Elem()
}

func (o DataSourceOutput) ToDataSourceOutput() DataSourceOutput {
	return o
}

func (o DataSourceOutput) ToDataSourceOutputWithContext(ctx context.Context) DataSourceOutput {
	return o
}

// The data source configuration, determined by its type,
// matching the specification in `config` from /data/sourcetypes.
func (o DataSourceOutput) Config() pulumi.MapOutput {
	return o.ApplyT(func(v *DataSource) pulumi.MapOutput { return v.Config }).(pulumi.MapOutput)
}

// The free form name of the data source.
func (o DataSourceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DataSource) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The data sources type, listed in API endpoint https://api.nsone.net/v1/data/sourcetypes.
func (o DataSourceOutput) Sourcetype() pulumi.StringOutput {
	return o.ApplyT(func(v *DataSource) pulumi.StringOutput { return v.Sourcetype }).(pulumi.StringOutput)
}

type DataSourceArrayOutput struct{ *pulumi.OutputState }

func (DataSourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DataSource)(nil)).Elem()
}

func (o DataSourceArrayOutput) ToDataSourceArrayOutput() DataSourceArrayOutput {
	return o
}

func (o DataSourceArrayOutput) ToDataSourceArrayOutputWithContext(ctx context.Context) DataSourceArrayOutput {
	return o
}

func (o DataSourceArrayOutput) Index(i pulumi.IntInput) DataSourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DataSource {
		return vs[0].([]*DataSource)[vs[1].(int)]
	}).(DataSourceOutput)
}

type DataSourceMapOutput struct{ *pulumi.OutputState }

func (DataSourceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DataSource)(nil)).Elem()
}

func (o DataSourceMapOutput) ToDataSourceMapOutput() DataSourceMapOutput {
	return o
}

func (o DataSourceMapOutput) ToDataSourceMapOutputWithContext(ctx context.Context) DataSourceMapOutput {
	return o
}

func (o DataSourceMapOutput) MapIndex(k pulumi.StringInput) DataSourceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DataSource {
		return vs[0].(map[string]*DataSource)[vs[1].(string)]
	}).(DataSourceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataSourceInput)(nil)).Elem(), &DataSource{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataSourceArrayInput)(nil)).Elem(), DataSourceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataSourceMapInput)(nil)).Elem(), DataSourceMap{})
	pulumi.RegisterOutputType(DataSourceOutput{})
	pulumi.RegisterOutputType(DataSourceArrayOutput{})
	pulumi.RegisterOutputType(DataSourceMapOutput{})
}
