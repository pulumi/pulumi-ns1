// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ns1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a NS1 Pulsar application resource. This can be used to create, modify, and delete applications.
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
// 		_, err := ns1.NewApplication(ctx, "ns1App", &ns1.ApplicationArgs{
// 			DefaultConfig: &ns1.ApplicationDefaultConfigArgs{
// 				Http:                   pulumi.Bool(true),
// 				Https:                  pulumi.Bool(false),
// 				Job_timeout_millis:     pulumi.Float64(100),
// 				Request_timeout_millis: pulumi.Float64(100),
// 				Static_values:          pulumi.Bool(true),
// 			},
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
// [Application Api Docs](https://ns1.com/api#get-list-pulsar-applications)
//
// ## Import
//
// ```sh
//  $ pulumi import ns1:index/application:Application `ns1_application`
// ```
//
//  So for the example above
//
// ```sh
//  $ pulumi import ns1:index/application:Application example terraform.example.io`
// ```
type Application struct {
	pulumi.CustomResourceState

	// Indicates whether or not this application is currently active and usable for traffic
	// steering.
	Active pulumi.BoolPtrOutput `pulumi:"active"`
	// The amount of time (in milliseconds) the browser should wait before running
	// measurements.
	BrowserWaitMillis pulumi.IntPtrOutput `pulumi:"browserWaitMillis"`
	// -(Optional) Default job configuration. If a field is present here and not on a specific job
	// associated with this application, the default value specified here is used..
	DefaultConfig ApplicationDefaultConfigPtrOutput `pulumi:"defaultConfig"`
	// -(Optional) Number of jobs to measure per user impression.
	JobsPerTransaction pulumi.IntPtrOutput `pulumi:"jobsPerTransaction"`
	// Descriptive name for this Pulsar app.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewApplication registers a new resource with the given unique name, arguments, and options.
func NewApplication(ctx *pulumi.Context,
	name string, args *ApplicationArgs, opts ...pulumi.ResourceOption) (*Application, error) {
	if args == nil {
		args = &ApplicationArgs{}
	}

	var resource Application
	err := ctx.RegisterResource("ns1:index/application:Application", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplication gets an existing Application resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationState, opts ...pulumi.ResourceOption) (*Application, error) {
	var resource Application
	err := ctx.ReadResource("ns1:index/application:Application", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Application resources.
type applicationState struct {
	// Indicates whether or not this application is currently active and usable for traffic
	// steering.
	Active *bool `pulumi:"active"`
	// The amount of time (in milliseconds) the browser should wait before running
	// measurements.
	BrowserWaitMillis *int `pulumi:"browserWaitMillis"`
	// -(Optional) Default job configuration. If a field is present here and not on a specific job
	// associated with this application, the default value specified here is used..
	DefaultConfig *ApplicationDefaultConfig `pulumi:"defaultConfig"`
	// -(Optional) Number of jobs to measure per user impression.
	JobsPerTransaction *int `pulumi:"jobsPerTransaction"`
	// Descriptive name for this Pulsar app.
	Name *string `pulumi:"name"`
}

type ApplicationState struct {
	// Indicates whether or not this application is currently active and usable for traffic
	// steering.
	Active pulumi.BoolPtrInput
	// The amount of time (in milliseconds) the browser should wait before running
	// measurements.
	BrowserWaitMillis pulumi.IntPtrInput
	// -(Optional) Default job configuration. If a field is present here and not on a specific job
	// associated with this application, the default value specified here is used..
	DefaultConfig ApplicationDefaultConfigPtrInput
	// -(Optional) Number of jobs to measure per user impression.
	JobsPerTransaction pulumi.IntPtrInput
	// Descriptive name for this Pulsar app.
	Name pulumi.StringPtrInput
}

func (ApplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationState)(nil)).Elem()
}

type applicationArgs struct {
	// Indicates whether or not this application is currently active and usable for traffic
	// steering.
	Active *bool `pulumi:"active"`
	// The amount of time (in milliseconds) the browser should wait before running
	// measurements.
	BrowserWaitMillis *int `pulumi:"browserWaitMillis"`
	// -(Optional) Default job configuration. If a field is present here and not on a specific job
	// associated with this application, the default value specified here is used..
	DefaultConfig *ApplicationDefaultConfig `pulumi:"defaultConfig"`
	// -(Optional) Number of jobs to measure per user impression.
	JobsPerTransaction *int `pulumi:"jobsPerTransaction"`
	// Descriptive name for this Pulsar app.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a Application resource.
type ApplicationArgs struct {
	// Indicates whether or not this application is currently active and usable for traffic
	// steering.
	Active pulumi.BoolPtrInput
	// The amount of time (in milliseconds) the browser should wait before running
	// measurements.
	BrowserWaitMillis pulumi.IntPtrInput
	// -(Optional) Default job configuration. If a field is present here and not on a specific job
	// associated with this application, the default value specified here is used..
	DefaultConfig ApplicationDefaultConfigPtrInput
	// -(Optional) Number of jobs to measure per user impression.
	JobsPerTransaction pulumi.IntPtrInput
	// Descriptive name for this Pulsar app.
	Name pulumi.StringPtrInput
}

func (ApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationArgs)(nil)).Elem()
}

type ApplicationInput interface {
	pulumi.Input

	ToApplicationOutput() ApplicationOutput
	ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput
}

func (*Application) ElementType() reflect.Type {
	return reflect.TypeOf((*Application)(nil))
}

func (i *Application) ToApplicationOutput() ApplicationOutput {
	return i.ToApplicationOutputWithContext(context.Background())
}

func (i *Application) ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationOutput)
}

func (i *Application) ToApplicationPtrOutput() ApplicationPtrOutput {
	return i.ToApplicationPtrOutputWithContext(context.Background())
}

func (i *Application) ToApplicationPtrOutputWithContext(ctx context.Context) ApplicationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationPtrOutput)
}

type ApplicationPtrInput interface {
	pulumi.Input

	ToApplicationPtrOutput() ApplicationPtrOutput
	ToApplicationPtrOutputWithContext(ctx context.Context) ApplicationPtrOutput
}

type applicationPtrType ApplicationArgs

func (*applicationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Application)(nil))
}

func (i *applicationPtrType) ToApplicationPtrOutput() ApplicationPtrOutput {
	return i.ToApplicationPtrOutputWithContext(context.Background())
}

func (i *applicationPtrType) ToApplicationPtrOutputWithContext(ctx context.Context) ApplicationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationPtrOutput)
}

// ApplicationArrayInput is an input type that accepts ApplicationArray and ApplicationArrayOutput values.
// You can construct a concrete instance of `ApplicationArrayInput` via:
//
//          ApplicationArray{ ApplicationArgs{...} }
type ApplicationArrayInput interface {
	pulumi.Input

	ToApplicationArrayOutput() ApplicationArrayOutput
	ToApplicationArrayOutputWithContext(context.Context) ApplicationArrayOutput
}

type ApplicationArray []ApplicationInput

func (ApplicationArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Application)(nil))
}

func (i ApplicationArray) ToApplicationArrayOutput() ApplicationArrayOutput {
	return i.ToApplicationArrayOutputWithContext(context.Background())
}

func (i ApplicationArray) ToApplicationArrayOutputWithContext(ctx context.Context) ApplicationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationArrayOutput)
}

// ApplicationMapInput is an input type that accepts ApplicationMap and ApplicationMapOutput values.
// You can construct a concrete instance of `ApplicationMapInput` via:
//
//          ApplicationMap{ "key": ApplicationArgs{...} }
type ApplicationMapInput interface {
	pulumi.Input

	ToApplicationMapOutput() ApplicationMapOutput
	ToApplicationMapOutputWithContext(context.Context) ApplicationMapOutput
}

type ApplicationMap map[string]ApplicationInput

func (ApplicationMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Application)(nil))
}

func (i ApplicationMap) ToApplicationMapOutput() ApplicationMapOutput {
	return i.ToApplicationMapOutputWithContext(context.Background())
}

func (i ApplicationMap) ToApplicationMapOutputWithContext(ctx context.Context) ApplicationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationMapOutput)
}

type ApplicationOutput struct {
	*pulumi.OutputState
}

func (ApplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Application)(nil))
}

func (o ApplicationOutput) ToApplicationOutput() ApplicationOutput {
	return o
}

func (o ApplicationOutput) ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput {
	return o
}

func (o ApplicationOutput) ToApplicationPtrOutput() ApplicationPtrOutput {
	return o.ToApplicationPtrOutputWithContext(context.Background())
}

func (o ApplicationOutput) ToApplicationPtrOutputWithContext(ctx context.Context) ApplicationPtrOutput {
	return o.ApplyT(func(v Application) *Application {
		return &v
	}).(ApplicationPtrOutput)
}

type ApplicationPtrOutput struct {
	*pulumi.OutputState
}

func (ApplicationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Application)(nil))
}

func (o ApplicationPtrOutput) ToApplicationPtrOutput() ApplicationPtrOutput {
	return o
}

func (o ApplicationPtrOutput) ToApplicationPtrOutputWithContext(ctx context.Context) ApplicationPtrOutput {
	return o
}

type ApplicationArrayOutput struct{ *pulumi.OutputState }

func (ApplicationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Application)(nil))
}

func (o ApplicationArrayOutput) ToApplicationArrayOutput() ApplicationArrayOutput {
	return o
}

func (o ApplicationArrayOutput) ToApplicationArrayOutputWithContext(ctx context.Context) ApplicationArrayOutput {
	return o
}

func (o ApplicationArrayOutput) Index(i pulumi.IntInput) ApplicationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Application {
		return vs[0].([]Application)[vs[1].(int)]
	}).(ApplicationOutput)
}

type ApplicationMapOutput struct{ *pulumi.OutputState }

func (ApplicationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Application)(nil))
}

func (o ApplicationMapOutput) ToApplicationMapOutput() ApplicationMapOutput {
	return o
}

func (o ApplicationMapOutput) ToApplicationMapOutputWithContext(ctx context.Context) ApplicationMapOutput {
	return o
}

func (o ApplicationMapOutput) MapIndex(k pulumi.StringInput) ApplicationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Application {
		return vs[0].(map[string]Application)[vs[1].(string)]
	}).(ApplicationOutput)
}

func init() {
	pulumi.RegisterOutputType(ApplicationOutput{})
	pulumi.RegisterOutputType(ApplicationPtrOutput{})
	pulumi.RegisterOutputType(ApplicationArrayOutput{})
	pulumi.RegisterOutputType(ApplicationMapOutput{})
}