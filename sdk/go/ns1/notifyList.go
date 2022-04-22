// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ns1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a NS1 Notify List resource. This can be used to create, modify, and delete notify lists.
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
// 		_, err := ns1.NewNotifyList(ctx, "nl", &ns1.NotifyListArgs{
// 			Notifications: NotifyListNotificationArray{
// 				&NotifyListNotificationArgs{
// 					Config: pulumi.AnyMap{
// 						"url": pulumi.Any("http://www.mywebhook.com"),
// 					},
// 					Type: pulumi.String("webhook"),
// 				},
// 				&NotifyListNotificationArgs{
// 					Config: pulumi.AnyMap{
// 						"email": pulumi.Any("test@test.com"),
// 					},
// 					Type: pulumi.String("email"),
// 				},
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
// [NotifyList Api Doc](https://ns1.com/api#notification-lists)
type NotifyList struct {
	pulumi.CustomResourceState

	// The free-form display name for the notify list.
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of notifiers. All notifiers in a notification list will receive notifications whenever an event is send to the list (e.g., when a monitoring job fails). Notifiers are documented below.
	Notifications NotifyListNotificationArrayOutput `pulumi:"notifications"`
}

// NewNotifyList registers a new resource with the given unique name, arguments, and options.
func NewNotifyList(ctx *pulumi.Context,
	name string, args *NotifyListArgs, opts ...pulumi.ResourceOption) (*NotifyList, error) {
	if args == nil {
		args = &NotifyListArgs{}
	}

	var resource NotifyList
	err := ctx.RegisterResource("ns1:index/notifyList:NotifyList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNotifyList gets an existing NotifyList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNotifyList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NotifyListState, opts ...pulumi.ResourceOption) (*NotifyList, error) {
	var resource NotifyList
	err := ctx.ReadResource("ns1:index/notifyList:NotifyList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NotifyList resources.
type notifyListState struct {
	// The free-form display name for the notify list.
	Name *string `pulumi:"name"`
	// A list of notifiers. All notifiers in a notification list will receive notifications whenever an event is send to the list (e.g., when a monitoring job fails). Notifiers are documented below.
	Notifications []NotifyListNotification `pulumi:"notifications"`
}

type NotifyListState struct {
	// The free-form display name for the notify list.
	Name pulumi.StringPtrInput
	// A list of notifiers. All notifiers in a notification list will receive notifications whenever an event is send to the list (e.g., when a monitoring job fails). Notifiers are documented below.
	Notifications NotifyListNotificationArrayInput
}

func (NotifyListState) ElementType() reflect.Type {
	return reflect.TypeOf((*notifyListState)(nil)).Elem()
}

type notifyListArgs struct {
	// The free-form display name for the notify list.
	Name *string `pulumi:"name"`
	// A list of notifiers. All notifiers in a notification list will receive notifications whenever an event is send to the list (e.g., when a monitoring job fails). Notifiers are documented below.
	Notifications []NotifyListNotification `pulumi:"notifications"`
}

// The set of arguments for constructing a NotifyList resource.
type NotifyListArgs struct {
	// The free-form display name for the notify list.
	Name pulumi.StringPtrInput
	// A list of notifiers. All notifiers in a notification list will receive notifications whenever an event is send to the list (e.g., when a monitoring job fails). Notifiers are documented below.
	Notifications NotifyListNotificationArrayInput
}

func (NotifyListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*notifyListArgs)(nil)).Elem()
}

type NotifyListInput interface {
	pulumi.Input

	ToNotifyListOutput() NotifyListOutput
	ToNotifyListOutputWithContext(ctx context.Context) NotifyListOutput
}

func (*NotifyList) ElementType() reflect.Type {
	return reflect.TypeOf((**NotifyList)(nil)).Elem()
}

func (i *NotifyList) ToNotifyListOutput() NotifyListOutput {
	return i.ToNotifyListOutputWithContext(context.Background())
}

func (i *NotifyList) ToNotifyListOutputWithContext(ctx context.Context) NotifyListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotifyListOutput)
}

// NotifyListArrayInput is an input type that accepts NotifyListArray and NotifyListArrayOutput values.
// You can construct a concrete instance of `NotifyListArrayInput` via:
//
//          NotifyListArray{ NotifyListArgs{...} }
type NotifyListArrayInput interface {
	pulumi.Input

	ToNotifyListArrayOutput() NotifyListArrayOutput
	ToNotifyListArrayOutputWithContext(context.Context) NotifyListArrayOutput
}

type NotifyListArray []NotifyListInput

func (NotifyListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NotifyList)(nil)).Elem()
}

func (i NotifyListArray) ToNotifyListArrayOutput() NotifyListArrayOutput {
	return i.ToNotifyListArrayOutputWithContext(context.Background())
}

func (i NotifyListArray) ToNotifyListArrayOutputWithContext(ctx context.Context) NotifyListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotifyListArrayOutput)
}

// NotifyListMapInput is an input type that accepts NotifyListMap and NotifyListMapOutput values.
// You can construct a concrete instance of `NotifyListMapInput` via:
//
//          NotifyListMap{ "key": NotifyListArgs{...} }
type NotifyListMapInput interface {
	pulumi.Input

	ToNotifyListMapOutput() NotifyListMapOutput
	ToNotifyListMapOutputWithContext(context.Context) NotifyListMapOutput
}

type NotifyListMap map[string]NotifyListInput

func (NotifyListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NotifyList)(nil)).Elem()
}

func (i NotifyListMap) ToNotifyListMapOutput() NotifyListMapOutput {
	return i.ToNotifyListMapOutputWithContext(context.Background())
}

func (i NotifyListMap) ToNotifyListMapOutputWithContext(ctx context.Context) NotifyListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotifyListMapOutput)
}

type NotifyListOutput struct{ *pulumi.OutputState }

func (NotifyListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NotifyList)(nil)).Elem()
}

func (o NotifyListOutput) ToNotifyListOutput() NotifyListOutput {
	return o
}

func (o NotifyListOutput) ToNotifyListOutputWithContext(ctx context.Context) NotifyListOutput {
	return o
}

type NotifyListArrayOutput struct{ *pulumi.OutputState }

func (NotifyListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NotifyList)(nil)).Elem()
}

func (o NotifyListArrayOutput) ToNotifyListArrayOutput() NotifyListArrayOutput {
	return o
}

func (o NotifyListArrayOutput) ToNotifyListArrayOutputWithContext(ctx context.Context) NotifyListArrayOutput {
	return o
}

func (o NotifyListArrayOutput) Index(i pulumi.IntInput) NotifyListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NotifyList {
		return vs[0].([]*NotifyList)[vs[1].(int)]
	}).(NotifyListOutput)
}

type NotifyListMapOutput struct{ *pulumi.OutputState }

func (NotifyListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NotifyList)(nil)).Elem()
}

func (o NotifyListMapOutput) ToNotifyListMapOutput() NotifyListMapOutput {
	return o
}

func (o NotifyListMapOutput) ToNotifyListMapOutputWithContext(ctx context.Context) NotifyListMapOutput {
	return o
}

func (o NotifyListMapOutput) MapIndex(k pulumi.StringInput) NotifyListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NotifyList {
		return vs[0].(map[string]*NotifyList)[vs[1].(string)]
	}).(NotifyListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NotifyListInput)(nil)).Elem(), &NotifyList{})
	pulumi.RegisterInputType(reflect.TypeOf((*NotifyListArrayInput)(nil)).Elem(), NotifyListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NotifyListMapInput)(nil)).Elem(), NotifyListMap{})
	pulumi.RegisterOutputType(NotifyListOutput{})
	pulumi.RegisterOutputType(NotifyListArrayOutput{})
	pulumi.RegisterOutputType(NotifyListMapOutput{})
}
