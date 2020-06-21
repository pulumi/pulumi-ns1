// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ns1

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a NS1 Data Source resource. This can be used to create, modify, and delete data sources.
//
//
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
	if args == nil || args.Sourcetype == nil {
		return nil, errors.New("missing required argument 'Sourcetype'")
	}
	if args == nil {
		args = &DataSourceArgs{}
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
