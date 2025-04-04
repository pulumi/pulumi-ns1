// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ns1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-ns1/sdk/v3/go/ns1/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides billing usage details about a NS1 account.
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
//			// Get query usage data for the given timeframe
//			queries, err := ns1.NewBillingUsage(ctx, "queries", &ns1.BillingUsageArgs{
//				MetricType: "queries",
//				From:       1731605824,
//				To:         1734197824,
//			})
//			if err != nil {
//				return err
//			}
//			// Get account limits data for the given timeframe
//			limits, err := ns1.NewBillingUsage(ctx, "limits", &ns1.BillingUsageArgs{
//				MetricType: "limits",
//				From:       1731605824,
//				To:         1734197824,
//			})
//			if err != nil {
//				return err
//			}
//			// Get RUM decisions usage data for the given timeframe
//			decisions, err := ns1.NewBillingUsage(ctx, "decisions", &ns1.BillingUsageArgs{
//				MetricType: "decisions",
//				From:       1731605824,
//				To:         1734197824,
//			})
//			if err != nil {
//				return err
//			}
//			// Get filter chains usage data
//			filterChains, err := ns1.NewBillingUsage(ctx, "filter_chains", &ns1.BillingUsageArgs{
//				MetricType: "filter-chains",
//			})
//			if err != nil {
//				return err
//			}
//			// Get monitoring jobs usage data
//			monitors, err := ns1.NewBillingUsage(ctx, "monitors", &ns1.BillingUsageArgs{
//				MetricType: "monitors",
//			})
//			if err != nil {
//				return err
//			}
//			// Get records usage data
//			records, err := ns1.NewBillingUsage(ctx, "records", &ns1.BillingUsageArgs{
//				MetricType: "records",
//			})
//			if err != nil {
//				return err
//			}
//			ctx.Export("totalQueries", queries.CleanQueries)
//			ctx.Export("totalDdosQueries", queries.DdosQueries)
//			ctx.Export("totalNxdResponses", queries.NxdResponses)
//			ctx.Export("queriesLimit", limits.QueriesLimit)
//			ctx.Export("totalDecisions", decisions.TotalUsage)
//			ctx.Export("decisionsLimit", limits.DecisionsLimit)
//			ctx.Export("totalFilterChains", filterChains.TotalUsage)
//			ctx.Export("filterChainsLimit", limits.FilterChainsLimit)
//			ctx.Export("totalMonitors", monitors.TotalUsage)
//			ctx.Export("monitorsLimit", limits.MonitorsLimit)
//			ctx.Export("totalRecords", records.TotalUsage)
//			ctx.Export("recordsLimit", limits.RecordsLimit)
//			return nil
//		})
//	}
//
// ```
func GetBillingUsage(ctx *pulumi.Context, args *GetBillingUsageArgs, opts ...pulumi.InvokeOption) (*GetBillingUsageResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetBillingUsageResult
	err := ctx.Invoke("ns1:index/getBillingUsage:getBillingUsage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBillingUsage.
type GetBillingUsageArgs struct {
	// The start timestamp for the data range in Unix epoch format.
	From int `pulumi:"from"`
	// The type of billing metric to retrieve. Must be one of: `queries`, `limits`, `decisions`, `filter-chains`, `monitors`, `records`.
	MetricType string `pulumi:"metricType"`
	// The end timestamp for the data range in Unix epoch format.
	To int `pulumi:"to"`
}

// A collection of values returned by getBillingUsage.
type GetBillingUsageResult struct {
	// (Computed) A list of network-specific query data containing:
	ByNetworks []GetBillingUsageByNetwork `pulumi:"byNetworks"`
	// (Computed) The queries limit for the China network.
	ChinaQueriesLimit int `pulumi:"chinaQueriesLimit"`
	// Clean queries for this day.
	CleanQueries int `pulumi:"cleanQueries"`
	// (Computed) Whether DDoS Protection is enabled.
	DdosProtectionEnabled bool `pulumi:"ddosProtectionEnabled"`
	// DDoS queries for this day.
	DdosQueries int `pulumi:"ddosQueries"`
	// (Computed) The RUM decisions limit for this billing cycle.
	DecisionsLimit int `pulumi:"decisionsLimit"`
	// (Computed) The filter chains limit for this billing cycle.
	FilterChainsLimit int `pulumi:"filterChainsLimit"`
	From              int `pulumi:"from"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// (Computed) Whether dedicated DNS usage counts towards managed DNS usage.
	IncludeDedicatedDnsNetworkInManagedDnsUsage bool   `pulumi:"includeDedicatedDnsNetworkInManagedDnsUsage"`
	MetricType                                  string `pulumi:"metricType"`
	// (Computed) The monitoring jobs limit for this billing cycle.
	MonitorsLimit int `pulumi:"monitorsLimit"`
	// (Computed) Whether NXD Protection is enabled.
	NxdProtectionEnabled bool `pulumi:"nxdProtectionEnabled"`
	// NXD responses for this day.
	NxdResponses int `pulumi:"nxdResponses"`
	// (Computed) The queries limit for this billing cycle.
	QueriesLimit int `pulumi:"queriesLimit"`
	// (Computed) The records limit for this billing cycle.
	RecordsLimit int `pulumi:"recordsLimit"`
	To           int `pulumi:"to"`
	// (Computed) The total usage count for the metric. Available for `decisions`, `filter-chains`, `monitors`, and `records` metrics.
	TotalUsage int `pulumi:"totalUsage"`
}

func GetBillingUsageOutput(ctx *pulumi.Context, args GetBillingUsageOutputArgs, opts ...pulumi.InvokeOption) GetBillingUsageResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetBillingUsageResultOutput, error) {
			args := v.(GetBillingUsageArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ns1:index/getBillingUsage:getBillingUsage", args, GetBillingUsageResultOutput{}, options).(GetBillingUsageResultOutput), nil
		}).(GetBillingUsageResultOutput)
}

// A collection of arguments for invoking getBillingUsage.
type GetBillingUsageOutputArgs struct {
	// The start timestamp for the data range in Unix epoch format.
	From pulumi.IntInput `pulumi:"from"`
	// The type of billing metric to retrieve. Must be one of: `queries`, `limits`, `decisions`, `filter-chains`, `monitors`, `records`.
	MetricType pulumi.StringInput `pulumi:"metricType"`
	// The end timestamp for the data range in Unix epoch format.
	To pulumi.IntInput `pulumi:"to"`
}

func (GetBillingUsageOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBillingUsageArgs)(nil)).Elem()
}

// A collection of values returned by getBillingUsage.
type GetBillingUsageResultOutput struct{ *pulumi.OutputState }

func (GetBillingUsageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBillingUsageResult)(nil)).Elem()
}

func (o GetBillingUsageResultOutput) ToGetBillingUsageResultOutput() GetBillingUsageResultOutput {
	return o
}

func (o GetBillingUsageResultOutput) ToGetBillingUsageResultOutputWithContext(ctx context.Context) GetBillingUsageResultOutput {
	return o
}

// (Computed) A list of network-specific query data containing:
func (o GetBillingUsageResultOutput) ByNetworks() GetBillingUsageByNetworkArrayOutput {
	return o.ApplyT(func(v GetBillingUsageResult) []GetBillingUsageByNetwork { return v.ByNetworks }).(GetBillingUsageByNetworkArrayOutput)
}

// (Computed) The queries limit for the China network.
func (o GetBillingUsageResultOutput) ChinaQueriesLimit() pulumi.IntOutput {
	return o.ApplyT(func(v GetBillingUsageResult) int { return v.ChinaQueriesLimit }).(pulumi.IntOutput)
}

// Clean queries for this day.
func (o GetBillingUsageResultOutput) CleanQueries() pulumi.IntOutput {
	return o.ApplyT(func(v GetBillingUsageResult) int { return v.CleanQueries }).(pulumi.IntOutput)
}

// (Computed) Whether DDoS Protection is enabled.
func (o GetBillingUsageResultOutput) DdosProtectionEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetBillingUsageResult) bool { return v.DdosProtectionEnabled }).(pulumi.BoolOutput)
}

// DDoS queries for this day.
func (o GetBillingUsageResultOutput) DdosQueries() pulumi.IntOutput {
	return o.ApplyT(func(v GetBillingUsageResult) int { return v.DdosQueries }).(pulumi.IntOutput)
}

// (Computed) The RUM decisions limit for this billing cycle.
func (o GetBillingUsageResultOutput) DecisionsLimit() pulumi.IntOutput {
	return o.ApplyT(func(v GetBillingUsageResult) int { return v.DecisionsLimit }).(pulumi.IntOutput)
}

// (Computed) The filter chains limit for this billing cycle.
func (o GetBillingUsageResultOutput) FilterChainsLimit() pulumi.IntOutput {
	return o.ApplyT(func(v GetBillingUsageResult) int { return v.FilterChainsLimit }).(pulumi.IntOutput)
}

func (o GetBillingUsageResultOutput) From() pulumi.IntOutput {
	return o.ApplyT(func(v GetBillingUsageResult) int { return v.From }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetBillingUsageResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetBillingUsageResult) string { return v.Id }).(pulumi.StringOutput)
}

// (Computed) Whether dedicated DNS usage counts towards managed DNS usage.
func (o GetBillingUsageResultOutput) IncludeDedicatedDnsNetworkInManagedDnsUsage() pulumi.BoolOutput {
	return o.ApplyT(func(v GetBillingUsageResult) bool { return v.IncludeDedicatedDnsNetworkInManagedDnsUsage }).(pulumi.BoolOutput)
}

func (o GetBillingUsageResultOutput) MetricType() pulumi.StringOutput {
	return o.ApplyT(func(v GetBillingUsageResult) string { return v.MetricType }).(pulumi.StringOutput)
}

// (Computed) The monitoring jobs limit for this billing cycle.
func (o GetBillingUsageResultOutput) MonitorsLimit() pulumi.IntOutput {
	return o.ApplyT(func(v GetBillingUsageResult) int { return v.MonitorsLimit }).(pulumi.IntOutput)
}

// (Computed) Whether NXD Protection is enabled.
func (o GetBillingUsageResultOutput) NxdProtectionEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetBillingUsageResult) bool { return v.NxdProtectionEnabled }).(pulumi.BoolOutput)
}

// NXD responses for this day.
func (o GetBillingUsageResultOutput) NxdResponses() pulumi.IntOutput {
	return o.ApplyT(func(v GetBillingUsageResult) int { return v.NxdResponses }).(pulumi.IntOutput)
}

// (Computed) The queries limit for this billing cycle.
func (o GetBillingUsageResultOutput) QueriesLimit() pulumi.IntOutput {
	return o.ApplyT(func(v GetBillingUsageResult) int { return v.QueriesLimit }).(pulumi.IntOutput)
}

// (Computed) The records limit for this billing cycle.
func (o GetBillingUsageResultOutput) RecordsLimit() pulumi.IntOutput {
	return o.ApplyT(func(v GetBillingUsageResult) int { return v.RecordsLimit }).(pulumi.IntOutput)
}

func (o GetBillingUsageResultOutput) To() pulumi.IntOutput {
	return o.ApplyT(func(v GetBillingUsageResult) int { return v.To }).(pulumi.IntOutput)
}

// (Computed) The total usage count for the metric. Available for `decisions`, `filter-chains`, `monitors`, and `records` metrics.
func (o GetBillingUsageResultOutput) TotalUsage() pulumi.IntOutput {
	return o.ApplyT(func(v GetBillingUsageResult) int { return v.TotalUsage }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(GetBillingUsageResultOutput{})
}
