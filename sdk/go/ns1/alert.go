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

// Provides a NS1 Alert resource. This can be used to create, modify, and delete alerts.
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
//			_, err := ns1.NewAlert(ctx, "example", &ns1.AlertArgs{
//				Name:              pulumi.String("Example Alert"),
//				Type:              pulumi.String("zone"),
//				Subtype:           pulumi.String("transfer_failed"),
//				NotificationLists: pulumi.StringArray{},
//				ZoneNames:         pulumi.StringArray{},
//				RecordIds:         pulumi.StringArray{},
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
// [Alerts Api Doc](https://ns1.com/api#alerts)
//
// ## Import
//
// ```sh
// $ pulumi import ns1:index/alert:Alert <name> <alert_id>`
// ```
type Alert struct {
	pulumi.CustomResourceState

	// (Read Only) The Unix timestamp representing when the alert configuration was created.
	CreatedAt pulumi.IntOutput `pulumi:"createdAt"`
	// (Read Only) The user or apikey that created this alert.
	CreatedBy pulumi.StringOutput `pulumi:"createdBy"`
	// The free-form display name for the alert.
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of id's for notification lists whose notifiers will be triggered by the alert.
	NotificationLists pulumi.StringArrayOutput `pulumi:"notificationLists"`
	// A list of record id's this alert applies to.
	RecordIds pulumi.StringArrayOutput `pulumi:"recordIds"`
	// The type of the alert.
	Subtype pulumi.StringOutput `pulumi:"subtype"`
	// The type of the alert.
	Type pulumi.StringOutput `pulumi:"type"`
	// (Read Only) The Unix timestamp representing when the alert configuration was last modified.
	UpdatedAt pulumi.IntOutput `pulumi:"updatedAt"`
	// (Read Only) The user or apikey that last modified this alert.
	UpdatedBy pulumi.StringOutput `pulumi:"updatedBy"`
	// A list of zones this alert applies to.
	ZoneNames pulumi.StringArrayOutput `pulumi:"zoneNames"`
}

// NewAlert registers a new resource with the given unique name, arguments, and options.
func NewAlert(ctx *pulumi.Context,
	name string, args *AlertArgs, opts ...pulumi.ResourceOption) (*Alert, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Subtype == nil {
		return nil, errors.New("invalid value for required argument 'Subtype'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Alert
	err := ctx.RegisterResource("ns1:index/alert:Alert", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAlert gets an existing Alert resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAlert(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AlertState, opts ...pulumi.ResourceOption) (*Alert, error) {
	var resource Alert
	err := ctx.ReadResource("ns1:index/alert:Alert", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Alert resources.
type alertState struct {
	// (Read Only) The Unix timestamp representing when the alert configuration was created.
	CreatedAt *int `pulumi:"createdAt"`
	// (Read Only) The user or apikey that created this alert.
	CreatedBy *string `pulumi:"createdBy"`
	// The free-form display name for the alert.
	Name *string `pulumi:"name"`
	// A list of id's for notification lists whose notifiers will be triggered by the alert.
	NotificationLists []string `pulumi:"notificationLists"`
	// A list of record id's this alert applies to.
	RecordIds []string `pulumi:"recordIds"`
	// The type of the alert.
	Subtype *string `pulumi:"subtype"`
	// The type of the alert.
	Type *string `pulumi:"type"`
	// (Read Only) The Unix timestamp representing when the alert configuration was last modified.
	UpdatedAt *int `pulumi:"updatedAt"`
	// (Read Only) The user or apikey that last modified this alert.
	UpdatedBy *string `pulumi:"updatedBy"`
	// A list of zones this alert applies to.
	ZoneNames []string `pulumi:"zoneNames"`
}

type AlertState struct {
	// (Read Only) The Unix timestamp representing when the alert configuration was created.
	CreatedAt pulumi.IntPtrInput
	// (Read Only) The user or apikey that created this alert.
	CreatedBy pulumi.StringPtrInput
	// The free-form display name for the alert.
	Name pulumi.StringPtrInput
	// A list of id's for notification lists whose notifiers will be triggered by the alert.
	NotificationLists pulumi.StringArrayInput
	// A list of record id's this alert applies to.
	RecordIds pulumi.StringArrayInput
	// The type of the alert.
	Subtype pulumi.StringPtrInput
	// The type of the alert.
	Type pulumi.StringPtrInput
	// (Read Only) The Unix timestamp representing when the alert configuration was last modified.
	UpdatedAt pulumi.IntPtrInput
	// (Read Only) The user or apikey that last modified this alert.
	UpdatedBy pulumi.StringPtrInput
	// A list of zones this alert applies to.
	ZoneNames pulumi.StringArrayInput
}

func (AlertState) ElementType() reflect.Type {
	return reflect.TypeOf((*alertState)(nil)).Elem()
}

type alertArgs struct {
	// The free-form display name for the alert.
	Name *string `pulumi:"name"`
	// A list of id's for notification lists whose notifiers will be triggered by the alert.
	NotificationLists []string `pulumi:"notificationLists"`
	// A list of record id's this alert applies to.
	RecordIds []string `pulumi:"recordIds"`
	// The type of the alert.
	Subtype string `pulumi:"subtype"`
	// The type of the alert.
	Type string `pulumi:"type"`
	// A list of zones this alert applies to.
	ZoneNames []string `pulumi:"zoneNames"`
}

// The set of arguments for constructing a Alert resource.
type AlertArgs struct {
	// The free-form display name for the alert.
	Name pulumi.StringPtrInput
	// A list of id's for notification lists whose notifiers will be triggered by the alert.
	NotificationLists pulumi.StringArrayInput
	// A list of record id's this alert applies to.
	RecordIds pulumi.StringArrayInput
	// The type of the alert.
	Subtype pulumi.StringInput
	// The type of the alert.
	Type pulumi.StringInput
	// A list of zones this alert applies to.
	ZoneNames pulumi.StringArrayInput
}

func (AlertArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*alertArgs)(nil)).Elem()
}

type AlertInput interface {
	pulumi.Input

	ToAlertOutput() AlertOutput
	ToAlertOutputWithContext(ctx context.Context) AlertOutput
}

func (*Alert) ElementType() reflect.Type {
	return reflect.TypeOf((**Alert)(nil)).Elem()
}

func (i *Alert) ToAlertOutput() AlertOutput {
	return i.ToAlertOutputWithContext(context.Background())
}

func (i *Alert) ToAlertOutputWithContext(ctx context.Context) AlertOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertOutput)
}

// AlertArrayInput is an input type that accepts AlertArray and AlertArrayOutput values.
// You can construct a concrete instance of `AlertArrayInput` via:
//
//	AlertArray{ AlertArgs{...} }
type AlertArrayInput interface {
	pulumi.Input

	ToAlertArrayOutput() AlertArrayOutput
	ToAlertArrayOutputWithContext(context.Context) AlertArrayOutput
}

type AlertArray []AlertInput

func (AlertArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Alert)(nil)).Elem()
}

func (i AlertArray) ToAlertArrayOutput() AlertArrayOutput {
	return i.ToAlertArrayOutputWithContext(context.Background())
}

func (i AlertArray) ToAlertArrayOutputWithContext(ctx context.Context) AlertArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertArrayOutput)
}

// AlertMapInput is an input type that accepts AlertMap and AlertMapOutput values.
// You can construct a concrete instance of `AlertMapInput` via:
//
//	AlertMap{ "key": AlertArgs{...} }
type AlertMapInput interface {
	pulumi.Input

	ToAlertMapOutput() AlertMapOutput
	ToAlertMapOutputWithContext(context.Context) AlertMapOutput
}

type AlertMap map[string]AlertInput

func (AlertMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Alert)(nil)).Elem()
}

func (i AlertMap) ToAlertMapOutput() AlertMapOutput {
	return i.ToAlertMapOutputWithContext(context.Background())
}

func (i AlertMap) ToAlertMapOutputWithContext(ctx context.Context) AlertMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertMapOutput)
}

type AlertOutput struct{ *pulumi.OutputState }

func (AlertOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Alert)(nil)).Elem()
}

func (o AlertOutput) ToAlertOutput() AlertOutput {
	return o
}

func (o AlertOutput) ToAlertOutputWithContext(ctx context.Context) AlertOutput {
	return o
}

// (Read Only) The Unix timestamp representing when the alert configuration was created.
func (o AlertOutput) CreatedAt() pulumi.IntOutput {
	return o.ApplyT(func(v *Alert) pulumi.IntOutput { return v.CreatedAt }).(pulumi.IntOutput)
}

// (Read Only) The user or apikey that created this alert.
func (o AlertOutput) CreatedBy() pulumi.StringOutput {
	return o.ApplyT(func(v *Alert) pulumi.StringOutput { return v.CreatedBy }).(pulumi.StringOutput)
}

// The free-form display name for the alert.
func (o AlertOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Alert) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A list of id's for notification lists whose notifiers will be triggered by the alert.
func (o AlertOutput) NotificationLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Alert) pulumi.StringArrayOutput { return v.NotificationLists }).(pulumi.StringArrayOutput)
}

// A list of record id's this alert applies to.
func (o AlertOutput) RecordIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Alert) pulumi.StringArrayOutput { return v.RecordIds }).(pulumi.StringArrayOutput)
}

// The type of the alert.
func (o AlertOutput) Subtype() pulumi.StringOutput {
	return o.ApplyT(func(v *Alert) pulumi.StringOutput { return v.Subtype }).(pulumi.StringOutput)
}

// The type of the alert.
func (o AlertOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Alert) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// (Read Only) The Unix timestamp representing when the alert configuration was last modified.
func (o AlertOutput) UpdatedAt() pulumi.IntOutput {
	return o.ApplyT(func(v *Alert) pulumi.IntOutput { return v.UpdatedAt }).(pulumi.IntOutput)
}

// (Read Only) The user or apikey that last modified this alert.
func (o AlertOutput) UpdatedBy() pulumi.StringOutput {
	return o.ApplyT(func(v *Alert) pulumi.StringOutput { return v.UpdatedBy }).(pulumi.StringOutput)
}

// A list of zones this alert applies to.
func (o AlertOutput) ZoneNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Alert) pulumi.StringArrayOutput { return v.ZoneNames }).(pulumi.StringArrayOutput)
}

type AlertArrayOutput struct{ *pulumi.OutputState }

func (AlertArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Alert)(nil)).Elem()
}

func (o AlertArrayOutput) ToAlertArrayOutput() AlertArrayOutput {
	return o
}

func (o AlertArrayOutput) ToAlertArrayOutputWithContext(ctx context.Context) AlertArrayOutput {
	return o
}

func (o AlertArrayOutput) Index(i pulumi.IntInput) AlertOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Alert {
		return vs[0].([]*Alert)[vs[1].(int)]
	}).(AlertOutput)
}

type AlertMapOutput struct{ *pulumi.OutputState }

func (AlertMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Alert)(nil)).Elem()
}

func (o AlertMapOutput) ToAlertMapOutput() AlertMapOutput {
	return o
}

func (o AlertMapOutput) ToAlertMapOutputWithContext(ctx context.Context) AlertMapOutput {
	return o
}

func (o AlertMapOutput) MapIndex(k pulumi.StringInput) AlertOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Alert {
		return vs[0].(map[string]*Alert)[vs[1].(string)]
	}).(AlertOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AlertInput)(nil)).Elem(), &Alert{})
	pulumi.RegisterInputType(reflect.TypeOf((*AlertArrayInput)(nil)).Elem(), AlertArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AlertMapInput)(nil)).Elem(), AlertMap{})
	pulumi.RegisterOutputType(AlertOutput{})
	pulumi.RegisterOutputType(AlertArrayOutput{})
	pulumi.RegisterOutputType(AlertMapOutput{})
}
