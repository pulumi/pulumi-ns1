// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ns1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a NS1 TSIG Key resource. This can be used to create, modify, and delete TSIG keys.
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
//			_, err := ns1.NewTsigkey(ctx, "example", &ns1.TsigkeyArgs{
//				Algorithm: pulumi.String("hmac-sha256"),
//				Secret:    pulumi.String("Ok1qR5IW1ajVka5cHPEJQIXfLyx5V3PSkFBROAzOn21JumDq6nIpoj6H8rfj5Uo+Ok55ZWQ0Wgrf302fDscHLA=="),
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
// [TSIG Keys Api Doc](https://ns1.com/api/#tsig)
//
// ## Import
//
// ```sh
//
//	$ pulumi import ns1:index/tsigkey:Tsigkey importTest <name>`
//
// ```
type Tsigkey struct {
	pulumi.CustomResourceState

	// The algorithm used to hash the TSIG key's secret.
	Algorithm pulumi.StringOutput `pulumi:"algorithm"`
	// The free form name of the tsigkey.
	Name pulumi.StringOutput `pulumi:"name"`
	// The key's secret to be hashed.
	Secret pulumi.StringOutput `pulumi:"secret"`
}

// NewTsigkey registers a new resource with the given unique name, arguments, and options.
func NewTsigkey(ctx *pulumi.Context,
	name string, args *TsigkeyArgs, opts ...pulumi.ResourceOption) (*Tsigkey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Algorithm == nil {
		return nil, errors.New("invalid value for required argument 'Algorithm'")
	}
	if args.Secret == nil {
		return nil, errors.New("invalid value for required argument 'Secret'")
	}
	var resource Tsigkey
	err := ctx.RegisterResource("ns1:index/tsigkey:Tsigkey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTsigkey gets an existing Tsigkey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTsigkey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TsigkeyState, opts ...pulumi.ResourceOption) (*Tsigkey, error) {
	var resource Tsigkey
	err := ctx.ReadResource("ns1:index/tsigkey:Tsigkey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Tsigkey resources.
type tsigkeyState struct {
	// The algorithm used to hash the TSIG key's secret.
	Algorithm *string `pulumi:"algorithm"`
	// The free form name of the tsigkey.
	Name *string `pulumi:"name"`
	// The key's secret to be hashed.
	Secret *string `pulumi:"secret"`
}

type TsigkeyState struct {
	// The algorithm used to hash the TSIG key's secret.
	Algorithm pulumi.StringPtrInput
	// The free form name of the tsigkey.
	Name pulumi.StringPtrInput
	// The key's secret to be hashed.
	Secret pulumi.StringPtrInput
}

func (TsigkeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*tsigkeyState)(nil)).Elem()
}

type tsigkeyArgs struct {
	// The algorithm used to hash the TSIG key's secret.
	Algorithm string `pulumi:"algorithm"`
	// The free form name of the tsigkey.
	Name *string `pulumi:"name"`
	// The key's secret to be hashed.
	Secret string `pulumi:"secret"`
}

// The set of arguments for constructing a Tsigkey resource.
type TsigkeyArgs struct {
	// The algorithm used to hash the TSIG key's secret.
	Algorithm pulumi.StringInput
	// The free form name of the tsigkey.
	Name pulumi.StringPtrInput
	// The key's secret to be hashed.
	Secret pulumi.StringInput
}

func (TsigkeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tsigkeyArgs)(nil)).Elem()
}

type TsigkeyInput interface {
	pulumi.Input

	ToTsigkeyOutput() TsigkeyOutput
	ToTsigkeyOutputWithContext(ctx context.Context) TsigkeyOutput
}

func (*Tsigkey) ElementType() reflect.Type {
	return reflect.TypeOf((**Tsigkey)(nil)).Elem()
}

func (i *Tsigkey) ToTsigkeyOutput() TsigkeyOutput {
	return i.ToTsigkeyOutputWithContext(context.Background())
}

func (i *Tsigkey) ToTsigkeyOutputWithContext(ctx context.Context) TsigkeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TsigkeyOutput)
}

// TsigkeyArrayInput is an input type that accepts TsigkeyArray and TsigkeyArrayOutput values.
// You can construct a concrete instance of `TsigkeyArrayInput` via:
//
//	TsigkeyArray{ TsigkeyArgs{...} }
type TsigkeyArrayInput interface {
	pulumi.Input

	ToTsigkeyArrayOutput() TsigkeyArrayOutput
	ToTsigkeyArrayOutputWithContext(context.Context) TsigkeyArrayOutput
}

type TsigkeyArray []TsigkeyInput

func (TsigkeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Tsigkey)(nil)).Elem()
}

func (i TsigkeyArray) ToTsigkeyArrayOutput() TsigkeyArrayOutput {
	return i.ToTsigkeyArrayOutputWithContext(context.Background())
}

func (i TsigkeyArray) ToTsigkeyArrayOutputWithContext(ctx context.Context) TsigkeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TsigkeyArrayOutput)
}

// TsigkeyMapInput is an input type that accepts TsigkeyMap and TsigkeyMapOutput values.
// You can construct a concrete instance of `TsigkeyMapInput` via:
//
//	TsigkeyMap{ "key": TsigkeyArgs{...} }
type TsigkeyMapInput interface {
	pulumi.Input

	ToTsigkeyMapOutput() TsigkeyMapOutput
	ToTsigkeyMapOutputWithContext(context.Context) TsigkeyMapOutput
}

type TsigkeyMap map[string]TsigkeyInput

func (TsigkeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Tsigkey)(nil)).Elem()
}

func (i TsigkeyMap) ToTsigkeyMapOutput() TsigkeyMapOutput {
	return i.ToTsigkeyMapOutputWithContext(context.Background())
}

func (i TsigkeyMap) ToTsigkeyMapOutputWithContext(ctx context.Context) TsigkeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TsigkeyMapOutput)
}

type TsigkeyOutput struct{ *pulumi.OutputState }

func (TsigkeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Tsigkey)(nil)).Elem()
}

func (o TsigkeyOutput) ToTsigkeyOutput() TsigkeyOutput {
	return o
}

func (o TsigkeyOutput) ToTsigkeyOutputWithContext(ctx context.Context) TsigkeyOutput {
	return o
}

// The algorithm used to hash the TSIG key's secret.
func (o TsigkeyOutput) Algorithm() pulumi.StringOutput {
	return o.ApplyT(func(v *Tsigkey) pulumi.StringOutput { return v.Algorithm }).(pulumi.StringOutput)
}

// The free form name of the tsigkey.
func (o TsigkeyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Tsigkey) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The key's secret to be hashed.
func (o TsigkeyOutput) Secret() pulumi.StringOutput {
	return o.ApplyT(func(v *Tsigkey) pulumi.StringOutput { return v.Secret }).(pulumi.StringOutput)
}

type TsigkeyArrayOutput struct{ *pulumi.OutputState }

func (TsigkeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Tsigkey)(nil)).Elem()
}

func (o TsigkeyArrayOutput) ToTsigkeyArrayOutput() TsigkeyArrayOutput {
	return o
}

func (o TsigkeyArrayOutput) ToTsigkeyArrayOutputWithContext(ctx context.Context) TsigkeyArrayOutput {
	return o
}

func (o TsigkeyArrayOutput) Index(i pulumi.IntInput) TsigkeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Tsigkey {
		return vs[0].([]*Tsigkey)[vs[1].(int)]
	}).(TsigkeyOutput)
}

type TsigkeyMapOutput struct{ *pulumi.OutputState }

func (TsigkeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Tsigkey)(nil)).Elem()
}

func (o TsigkeyMapOutput) ToTsigkeyMapOutput() TsigkeyMapOutput {
	return o
}

func (o TsigkeyMapOutput) ToTsigkeyMapOutputWithContext(ctx context.Context) TsigkeyMapOutput {
	return o
}

func (o TsigkeyMapOutput) MapIndex(k pulumi.StringInput) TsigkeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Tsigkey {
		return vs[0].(map[string]*Tsigkey)[vs[1].(string)]
	}).(TsigkeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TsigkeyInput)(nil)).Elem(), &Tsigkey{})
	pulumi.RegisterInputType(reflect.TypeOf((*TsigkeyArrayInput)(nil)).Elem(), TsigkeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TsigkeyMapInput)(nil)).Elem(), TsigkeyMap{})
	pulumi.RegisterOutputType(TsigkeyOutput{})
	pulumi.RegisterOutputType(TsigkeyArrayOutput{})
	pulumi.RegisterOutputType(TsigkeyMapOutput{})
}
