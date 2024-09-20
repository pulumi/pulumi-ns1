// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ns1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-ns1/sdk/v3/go/ns1/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides details about a NS1 Record. Use this if you would simply like to read
// information from NS1 into your configurations. For read/write operations, you
// should use a resource.
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
//			// Get details about a NS1 Record.
//			_, err := ns1.LookupRecord(ctx, &ns1.LookupRecordArgs{
//				Zone:   "example.io",
//				Domain: "terraform.example.io",
//				Type:   "A",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupRecord(ctx *pulumi.Context, args *LookupRecordArgs, opts ...pulumi.InvokeOption) (*LookupRecordResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRecordResult
	err := ctx.Invoke("ns1:index/getRecord:getRecord", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRecord.
type LookupRecordArgs struct {
	// The records' domain.
	Domain string `pulumi:"domain"`
	// The records' RR type.
	Type string `pulumi:"type"`
	// The zone the record belongs to.
	Zone string `pulumi:"zone"`
}

// A collection of values returned by getRecord.
type LookupRecordResult struct {
	// List of NS1 answers.
	Answers []GetRecordAnswer `pulumi:"answers"`
	Domain  string            `pulumi:"domain"`
	// List of NS1 filters.
	Filters []GetRecordFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The target record this links to.
	Link string `pulumi:"link"`
	// Map of metadata
	Meta        map[string]string `pulumi:"meta"`
	OverrideTtl bool              `pulumi:"overrideTtl"`
	// List of regions.
	Regions      []GetRecordRegion `pulumi:"regions"`
	ShortAnswers []string          `pulumi:"shortAnswers"`
	Tags         map[string]string `pulumi:"tags"`
	// The records' time to live (in seconds).
	Ttl  int    `pulumi:"ttl"`
	Type string `pulumi:"type"`
	// Whether to use EDNS client subnet data when available (in filter chain).
	UseClientSubnet bool   `pulumi:"useClientSubnet"`
	Zone            string `pulumi:"zone"`
}

func LookupRecordOutput(ctx *pulumi.Context, args LookupRecordOutputArgs, opts ...pulumi.InvokeOption) LookupRecordResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRecordResultOutput, error) {
			args := v.(LookupRecordArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv LookupRecordResult
			secret, err := ctx.InvokePackageRaw("ns1:index/getRecord:getRecord", args, &rv, "", opts...)
			if err != nil {
				return LookupRecordResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(LookupRecordResultOutput)
			if secret {
				return pulumi.ToSecret(output).(LookupRecordResultOutput), nil
			}
			return output, nil
		}).(LookupRecordResultOutput)
}

// A collection of arguments for invoking getRecord.
type LookupRecordOutputArgs struct {
	// The records' domain.
	Domain pulumi.StringInput `pulumi:"domain"`
	// The records' RR type.
	Type pulumi.StringInput `pulumi:"type"`
	// The zone the record belongs to.
	Zone pulumi.StringInput `pulumi:"zone"`
}

func (LookupRecordOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRecordArgs)(nil)).Elem()
}

// A collection of values returned by getRecord.
type LookupRecordResultOutput struct{ *pulumi.OutputState }

func (LookupRecordResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRecordResult)(nil)).Elem()
}

func (o LookupRecordResultOutput) ToLookupRecordResultOutput() LookupRecordResultOutput {
	return o
}

func (o LookupRecordResultOutput) ToLookupRecordResultOutputWithContext(ctx context.Context) LookupRecordResultOutput {
	return o
}

// List of NS1 answers.
func (o LookupRecordResultOutput) Answers() GetRecordAnswerArrayOutput {
	return o.ApplyT(func(v LookupRecordResult) []GetRecordAnswer { return v.Answers }).(GetRecordAnswerArrayOutput)
}

func (o LookupRecordResultOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRecordResult) string { return v.Domain }).(pulumi.StringOutput)
}

// List of NS1 filters.
func (o LookupRecordResultOutput) Filters() GetRecordFilterArrayOutput {
	return o.ApplyT(func(v LookupRecordResult) []GetRecordFilter { return v.Filters }).(GetRecordFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRecordResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRecordResult) string { return v.Id }).(pulumi.StringOutput)
}

// The target record this links to.
func (o LookupRecordResultOutput) Link() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRecordResult) string { return v.Link }).(pulumi.StringOutput)
}

// Map of metadata
func (o LookupRecordResultOutput) Meta() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRecordResult) map[string]string { return v.Meta }).(pulumi.StringMapOutput)
}

func (o LookupRecordResultOutput) OverrideTtl() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRecordResult) bool { return v.OverrideTtl }).(pulumi.BoolOutput)
}

// List of regions.
func (o LookupRecordResultOutput) Regions() GetRecordRegionArrayOutput {
	return o.ApplyT(func(v LookupRecordResult) []GetRecordRegion { return v.Regions }).(GetRecordRegionArrayOutput)
}

func (o LookupRecordResultOutput) ShortAnswers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRecordResult) []string { return v.ShortAnswers }).(pulumi.StringArrayOutput)
}

func (o LookupRecordResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRecordResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The records' time to live (in seconds).
func (o LookupRecordResultOutput) Ttl() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRecordResult) int { return v.Ttl }).(pulumi.IntOutput)
}

func (o LookupRecordResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRecordResult) string { return v.Type }).(pulumi.StringOutput)
}

// Whether to use EDNS client subnet data when available (in filter chain).
func (o LookupRecordResultOutput) UseClientSubnet() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRecordResult) bool { return v.UseClientSubnet }).(pulumi.BoolOutput)
}

func (o LookupRecordResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRecordResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRecordResultOutput{})
}
