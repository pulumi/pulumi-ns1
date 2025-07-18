// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ns1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-ns1/sdk/v3/go/ns1/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a NS1 Global IP Whitelist resource.
//
// This can be used to create, modify, and delete Global IP Whitelists.
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
//			_, err := ns1.NewAccountWhitelist(ctx, "example", &ns1.AccountWhitelistArgs{
//				Name: pulumi.String("Example Whitelist"),
//				Values: pulumi.StringArray{
//					pulumi.String("1.1.1.1"),
//					pulumi.String("2.2.2.2"),
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
//
// > You current source IP must be present in one of the whitelists to prevent locking yourself out.
//
// ## NS1 Documentation
//
// [Global IP Whitelist Doc](https://ns1.com/api?docId=2282)
//
// ## Import
//
// ```sh
// $ pulumi import ns1:index/accountWhitelist:AccountWhitelist example <ID>`
// ```
type AccountWhitelist struct {
	pulumi.CustomResourceState

	// The free form name of the whitelist.
	Name pulumi.StringOutput `pulumi:"name"`
	// Array of IP addresses/networks from which to allow access.
	Values pulumi.StringArrayOutput `pulumi:"values"`
}

// NewAccountWhitelist registers a new resource with the given unique name, arguments, and options.
func NewAccountWhitelist(ctx *pulumi.Context,
	name string, args *AccountWhitelistArgs, opts ...pulumi.ResourceOption) (*AccountWhitelist, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Values == nil {
		return nil, errors.New("invalid value for required argument 'Values'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AccountWhitelist
	err := ctx.RegisterResource("ns1:index/accountWhitelist:AccountWhitelist", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccountWhitelist gets an existing AccountWhitelist resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccountWhitelist(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountWhitelistState, opts ...pulumi.ResourceOption) (*AccountWhitelist, error) {
	var resource AccountWhitelist
	err := ctx.ReadResource("ns1:index/accountWhitelist:AccountWhitelist", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccountWhitelist resources.
type accountWhitelistState struct {
	// The free form name of the whitelist.
	Name *string `pulumi:"name"`
	// Array of IP addresses/networks from which to allow access.
	Values []string `pulumi:"values"`
}

type AccountWhitelistState struct {
	// The free form name of the whitelist.
	Name pulumi.StringPtrInput
	// Array of IP addresses/networks from which to allow access.
	Values pulumi.StringArrayInput
}

func (AccountWhitelistState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountWhitelistState)(nil)).Elem()
}

type accountWhitelistArgs struct {
	// The free form name of the whitelist.
	Name *string `pulumi:"name"`
	// Array of IP addresses/networks from which to allow access.
	Values []string `pulumi:"values"`
}

// The set of arguments for constructing a AccountWhitelist resource.
type AccountWhitelistArgs struct {
	// The free form name of the whitelist.
	Name pulumi.StringPtrInput
	// Array of IP addresses/networks from which to allow access.
	Values pulumi.StringArrayInput
}

func (AccountWhitelistArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accountWhitelistArgs)(nil)).Elem()
}

type AccountWhitelistInput interface {
	pulumi.Input

	ToAccountWhitelistOutput() AccountWhitelistOutput
	ToAccountWhitelistOutputWithContext(ctx context.Context) AccountWhitelistOutput
}

func (*AccountWhitelist) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountWhitelist)(nil)).Elem()
}

func (i *AccountWhitelist) ToAccountWhitelistOutput() AccountWhitelistOutput {
	return i.ToAccountWhitelistOutputWithContext(context.Background())
}

func (i *AccountWhitelist) ToAccountWhitelistOutputWithContext(ctx context.Context) AccountWhitelistOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountWhitelistOutput)
}

// AccountWhitelistArrayInput is an input type that accepts AccountWhitelistArray and AccountWhitelistArrayOutput values.
// You can construct a concrete instance of `AccountWhitelistArrayInput` via:
//
//	AccountWhitelistArray{ AccountWhitelistArgs{...} }
type AccountWhitelistArrayInput interface {
	pulumi.Input

	ToAccountWhitelistArrayOutput() AccountWhitelistArrayOutput
	ToAccountWhitelistArrayOutputWithContext(context.Context) AccountWhitelistArrayOutput
}

type AccountWhitelistArray []AccountWhitelistInput

func (AccountWhitelistArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccountWhitelist)(nil)).Elem()
}

func (i AccountWhitelistArray) ToAccountWhitelistArrayOutput() AccountWhitelistArrayOutput {
	return i.ToAccountWhitelistArrayOutputWithContext(context.Background())
}

func (i AccountWhitelistArray) ToAccountWhitelistArrayOutputWithContext(ctx context.Context) AccountWhitelistArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountWhitelistArrayOutput)
}

// AccountWhitelistMapInput is an input type that accepts AccountWhitelistMap and AccountWhitelistMapOutput values.
// You can construct a concrete instance of `AccountWhitelistMapInput` via:
//
//	AccountWhitelistMap{ "key": AccountWhitelistArgs{...} }
type AccountWhitelistMapInput interface {
	pulumi.Input

	ToAccountWhitelistMapOutput() AccountWhitelistMapOutput
	ToAccountWhitelistMapOutputWithContext(context.Context) AccountWhitelistMapOutput
}

type AccountWhitelistMap map[string]AccountWhitelistInput

func (AccountWhitelistMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccountWhitelist)(nil)).Elem()
}

func (i AccountWhitelistMap) ToAccountWhitelistMapOutput() AccountWhitelistMapOutput {
	return i.ToAccountWhitelistMapOutputWithContext(context.Background())
}

func (i AccountWhitelistMap) ToAccountWhitelistMapOutputWithContext(ctx context.Context) AccountWhitelistMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountWhitelistMapOutput)
}

type AccountWhitelistOutput struct{ *pulumi.OutputState }

func (AccountWhitelistOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountWhitelist)(nil)).Elem()
}

func (o AccountWhitelistOutput) ToAccountWhitelistOutput() AccountWhitelistOutput {
	return o
}

func (o AccountWhitelistOutput) ToAccountWhitelistOutputWithContext(ctx context.Context) AccountWhitelistOutput {
	return o
}

// The free form name of the whitelist.
func (o AccountWhitelistOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AccountWhitelist) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Array of IP addresses/networks from which to allow access.
func (o AccountWhitelistOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AccountWhitelist) pulumi.StringArrayOutput { return v.Values }).(pulumi.StringArrayOutput)
}

type AccountWhitelistArrayOutput struct{ *pulumi.OutputState }

func (AccountWhitelistArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccountWhitelist)(nil)).Elem()
}

func (o AccountWhitelistArrayOutput) ToAccountWhitelistArrayOutput() AccountWhitelistArrayOutput {
	return o
}

func (o AccountWhitelistArrayOutput) ToAccountWhitelistArrayOutputWithContext(ctx context.Context) AccountWhitelistArrayOutput {
	return o
}

func (o AccountWhitelistArrayOutput) Index(i pulumi.IntInput) AccountWhitelistOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AccountWhitelist {
		return vs[0].([]*AccountWhitelist)[vs[1].(int)]
	}).(AccountWhitelistOutput)
}

type AccountWhitelistMapOutput struct{ *pulumi.OutputState }

func (AccountWhitelistMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccountWhitelist)(nil)).Elem()
}

func (o AccountWhitelistMapOutput) ToAccountWhitelistMapOutput() AccountWhitelistMapOutput {
	return o
}

func (o AccountWhitelistMapOutput) ToAccountWhitelistMapOutputWithContext(ctx context.Context) AccountWhitelistMapOutput {
	return o
}

func (o AccountWhitelistMapOutput) MapIndex(k pulumi.StringInput) AccountWhitelistOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AccountWhitelist {
		return vs[0].(map[string]*AccountWhitelist)[vs[1].(string)]
	}).(AccountWhitelistOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccountWhitelistInput)(nil)).Elem(), &AccountWhitelist{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccountWhitelistArrayInput)(nil)).Elem(), AccountWhitelistArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccountWhitelistMapInput)(nil)).Elem(), AccountWhitelistMap{})
	pulumi.RegisterOutputType(AccountWhitelistOutput{})
	pulumi.RegisterOutputType(AccountWhitelistArrayOutput{})
	pulumi.RegisterOutputType(AccountWhitelistMapOutput{})
}
