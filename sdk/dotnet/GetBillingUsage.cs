// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ns1
{
    public static class GetBillingUsage
    {
        /// <summary>
        /// Provides billing usage details about a NS1 account.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Ns1 = Pulumi.Ns1;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     // Get query usage data for the given timeframe
        ///     var queries = Ns1.GetBillingUsage.Invoke(new()
        ///     {
        ///         MetricType = "queries",
        ///         From = 1738368000,
        ///         To = 1740787199,
        ///     });
        /// 
        ///     // Get account limits data for the given timeframe
        ///     var limits = Ns1.GetBillingUsage.Invoke(new()
        ///     {
        ///         MetricType = "limits",
        ///         From = 1738368000,
        ///         To = 1740787199,
        ///     });
        /// 
        ///     // Get RUM decisions usage data for the given timeframe
        ///     var decisions = Ns1.GetBillingUsage.Invoke(new()
        ///     {
        ///         MetricType = "decisions",
        ///         From = 1738368000,
        ///         To = 1740787199,
        ///     });
        /// 
        ///     // Get filter chains usage data
        ///     var filterChains = Ns1.GetBillingUsage.Invoke(new()
        ///     {
        ///         MetricType = "filter-chains",
        ///     });
        /// 
        ///     // Get monitoring jobs usage data
        ///     var monitors = Ns1.GetBillingUsage.Invoke(new()
        ///     {
        ///         MetricType = "monitors",
        ///     });
        /// 
        ///     // Get records usage data
        ///     var records = Ns1.GetBillingUsage.Invoke(new()
        ///     {
        ///         MetricType = "records",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["totalQueries"] = queries.Apply(getBillingUsageResult =&gt; getBillingUsageResult.CleanQueries),
        ///         ["totalDdosQueries"] = queries.Apply(getBillingUsageResult =&gt; getBillingUsageResult.DdosQueries),
        ///         ["totalNxdResponses"] = queries.Apply(getBillingUsageResult =&gt; getBillingUsageResult.NxdResponses),
        ///         ["queriesLimit"] = limits.Apply(getBillingUsageResult =&gt; getBillingUsageResult.QueriesLimit),
        ///         ["totalDecisions"] = decisions.Apply(getBillingUsageResult =&gt; getBillingUsageResult.TotalUsage),
        ///         ["decisionsLimit"] = limits.Apply(getBillingUsageResult =&gt; getBillingUsageResult.DecisionsLimit),
        ///         ["totalFilterChains"] = filterChains.Apply(getBillingUsageResult =&gt; getBillingUsageResult.TotalUsage),
        ///         ["filterChainsLimit"] = limits.Apply(getBillingUsageResult =&gt; getBillingUsageResult.FilterChainsLimit),
        ///         ["totalMonitors"] = monitors.Apply(getBillingUsageResult =&gt; getBillingUsageResult.TotalUsage),
        ///         ["monitorsLimit"] = limits.Apply(getBillingUsageResult =&gt; getBillingUsageResult.MonitorsLimit),
        ///         ["totalRecords"] = records.Apply(getBillingUsageResult =&gt; getBillingUsageResult.TotalUsage),
        ///         ["recordsLimit"] = limits.Apply(getBillingUsageResult =&gt; getBillingUsageResult.RecordsLimit),
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Task<GetBillingUsageResult> InvokeAsync(GetBillingUsageArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetBillingUsageResult>("ns1:index/getBillingUsage:getBillingUsage", args ?? new GetBillingUsageArgs(), options.WithDefaults());

        /// <summary>
        /// Provides billing usage details about a NS1 account.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Ns1 = Pulumi.Ns1;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     // Get query usage data for the given timeframe
        ///     var queries = Ns1.GetBillingUsage.Invoke(new()
        ///     {
        ///         MetricType = "queries",
        ///         From = 1738368000,
        ///         To = 1740787199,
        ///     });
        /// 
        ///     // Get account limits data for the given timeframe
        ///     var limits = Ns1.GetBillingUsage.Invoke(new()
        ///     {
        ///         MetricType = "limits",
        ///         From = 1738368000,
        ///         To = 1740787199,
        ///     });
        /// 
        ///     // Get RUM decisions usage data for the given timeframe
        ///     var decisions = Ns1.GetBillingUsage.Invoke(new()
        ///     {
        ///         MetricType = "decisions",
        ///         From = 1738368000,
        ///         To = 1740787199,
        ///     });
        /// 
        ///     // Get filter chains usage data
        ///     var filterChains = Ns1.GetBillingUsage.Invoke(new()
        ///     {
        ///         MetricType = "filter-chains",
        ///     });
        /// 
        ///     // Get monitoring jobs usage data
        ///     var monitors = Ns1.GetBillingUsage.Invoke(new()
        ///     {
        ///         MetricType = "monitors",
        ///     });
        /// 
        ///     // Get records usage data
        ///     var records = Ns1.GetBillingUsage.Invoke(new()
        ///     {
        ///         MetricType = "records",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["totalQueries"] = queries.Apply(getBillingUsageResult =&gt; getBillingUsageResult.CleanQueries),
        ///         ["totalDdosQueries"] = queries.Apply(getBillingUsageResult =&gt; getBillingUsageResult.DdosQueries),
        ///         ["totalNxdResponses"] = queries.Apply(getBillingUsageResult =&gt; getBillingUsageResult.NxdResponses),
        ///         ["queriesLimit"] = limits.Apply(getBillingUsageResult =&gt; getBillingUsageResult.QueriesLimit),
        ///         ["totalDecisions"] = decisions.Apply(getBillingUsageResult =&gt; getBillingUsageResult.TotalUsage),
        ///         ["decisionsLimit"] = limits.Apply(getBillingUsageResult =&gt; getBillingUsageResult.DecisionsLimit),
        ///         ["totalFilterChains"] = filterChains.Apply(getBillingUsageResult =&gt; getBillingUsageResult.TotalUsage),
        ///         ["filterChainsLimit"] = limits.Apply(getBillingUsageResult =&gt; getBillingUsageResult.FilterChainsLimit),
        ///         ["totalMonitors"] = monitors.Apply(getBillingUsageResult =&gt; getBillingUsageResult.TotalUsage),
        ///         ["monitorsLimit"] = limits.Apply(getBillingUsageResult =&gt; getBillingUsageResult.MonitorsLimit),
        ///         ["totalRecords"] = records.Apply(getBillingUsageResult =&gt; getBillingUsageResult.TotalUsage),
        ///         ["recordsLimit"] = limits.Apply(getBillingUsageResult =&gt; getBillingUsageResult.RecordsLimit),
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Output<GetBillingUsageResult> Invoke(GetBillingUsageInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetBillingUsageResult>("ns1:index/getBillingUsage:getBillingUsage", args ?? new GetBillingUsageInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Provides billing usage details about a NS1 account.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Ns1 = Pulumi.Ns1;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     // Get query usage data for the given timeframe
        ///     var queries = Ns1.GetBillingUsage.Invoke(new()
        ///     {
        ///         MetricType = "queries",
        ///         From = 1738368000,
        ///         To = 1740787199,
        ///     });
        /// 
        ///     // Get account limits data for the given timeframe
        ///     var limits = Ns1.GetBillingUsage.Invoke(new()
        ///     {
        ///         MetricType = "limits",
        ///         From = 1738368000,
        ///         To = 1740787199,
        ///     });
        /// 
        ///     // Get RUM decisions usage data for the given timeframe
        ///     var decisions = Ns1.GetBillingUsage.Invoke(new()
        ///     {
        ///         MetricType = "decisions",
        ///         From = 1738368000,
        ///         To = 1740787199,
        ///     });
        /// 
        ///     // Get filter chains usage data
        ///     var filterChains = Ns1.GetBillingUsage.Invoke(new()
        ///     {
        ///         MetricType = "filter-chains",
        ///     });
        /// 
        ///     // Get monitoring jobs usage data
        ///     var monitors = Ns1.GetBillingUsage.Invoke(new()
        ///     {
        ///         MetricType = "monitors",
        ///     });
        /// 
        ///     // Get records usage data
        ///     var records = Ns1.GetBillingUsage.Invoke(new()
        ///     {
        ///         MetricType = "records",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["totalQueries"] = queries.Apply(getBillingUsageResult =&gt; getBillingUsageResult.CleanQueries),
        ///         ["totalDdosQueries"] = queries.Apply(getBillingUsageResult =&gt; getBillingUsageResult.DdosQueries),
        ///         ["totalNxdResponses"] = queries.Apply(getBillingUsageResult =&gt; getBillingUsageResult.NxdResponses),
        ///         ["queriesLimit"] = limits.Apply(getBillingUsageResult =&gt; getBillingUsageResult.QueriesLimit),
        ///         ["totalDecisions"] = decisions.Apply(getBillingUsageResult =&gt; getBillingUsageResult.TotalUsage),
        ///         ["decisionsLimit"] = limits.Apply(getBillingUsageResult =&gt; getBillingUsageResult.DecisionsLimit),
        ///         ["totalFilterChains"] = filterChains.Apply(getBillingUsageResult =&gt; getBillingUsageResult.TotalUsage),
        ///         ["filterChainsLimit"] = limits.Apply(getBillingUsageResult =&gt; getBillingUsageResult.FilterChainsLimit),
        ///         ["totalMonitors"] = monitors.Apply(getBillingUsageResult =&gt; getBillingUsageResult.TotalUsage),
        ///         ["monitorsLimit"] = limits.Apply(getBillingUsageResult =&gt; getBillingUsageResult.MonitorsLimit),
        ///         ["totalRecords"] = records.Apply(getBillingUsageResult =&gt; getBillingUsageResult.TotalUsage),
        ///         ["recordsLimit"] = limits.Apply(getBillingUsageResult =&gt; getBillingUsageResult.RecordsLimit),
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Output<GetBillingUsageResult> Invoke(GetBillingUsageInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetBillingUsageResult>("ns1:index/getBillingUsage:getBillingUsage", args ?? new GetBillingUsageInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBillingUsageArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The start timestamp for the data range in Unix epoch format.
        /// </summary>
        [Input("from")]
        public int? From { get; set; }

        /// <summary>
        /// The type of billing metric to retrieve. Must be one of: `queries`, `limits`, `decisions`, `filter-chains`, `monitors`, `records`.
        /// </summary>
        [Input("metricType", required: true)]
        public string MetricType { get; set; } = null!;

        /// <summary>
        /// The end timestamp for the data range in Unix epoch format.
        /// </summary>
        [Input("to")]
        public int? To { get; set; }

        public GetBillingUsageArgs()
        {
        }
        public static new GetBillingUsageArgs Empty => new GetBillingUsageArgs();
    }

    public sealed class GetBillingUsageInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The start timestamp for the data range in Unix epoch format.
        /// </summary>
        [Input("from")]
        public Input<int>? From { get; set; }

        /// <summary>
        /// The type of billing metric to retrieve. Must be one of: `queries`, `limits`, `decisions`, `filter-chains`, `monitors`, `records`.
        /// </summary>
        [Input("metricType", required: true)]
        public Input<string> MetricType { get; set; } = null!;

        /// <summary>
        /// The end timestamp for the data range in Unix epoch format.
        /// </summary>
        [Input("to")]
        public Input<int>? To { get; set; }

        public GetBillingUsageInvokeArgs()
        {
        }
        public static new GetBillingUsageInvokeArgs Empty => new GetBillingUsageInvokeArgs();
    }


    [OutputType]
    public sealed class GetBillingUsageResult
    {
        /// <summary>
        /// (Computed) A list of network-specific query data containing:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetBillingUsageByNetworkResult> ByNetworks;
        /// <summary>
        /// (Computed) The queries limit for the China network.
        /// </summary>
        public readonly int ChinaQueriesLimit;
        /// <summary>
        /// Clean queries for this day.
        /// </summary>
        public readonly int CleanQueries;
        /// <summary>
        /// (Computed) Whether DDoS Protection is enabled.
        /// </summary>
        public readonly bool DdosProtectionEnabled;
        /// <summary>
        /// DDoS queries for this day.
        /// </summary>
        public readonly int DdosQueries;
        /// <summary>
        /// (Computed) The RUM decisions limit for this billing cycle.
        /// </summary>
        public readonly int DecisionsLimit;
        /// <summary>
        /// (Computed) The filter chains limit for this billing cycle.
        /// </summary>
        public readonly int FilterChainsLimit;
        public readonly int? From;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// (Computed) Whether dedicated DNS usage counts towards managed DNS usage.
        /// </summary>
        public readonly bool IncludeDedicatedDnsNetworkInManagedDnsUsage;
        public readonly string MetricType;
        /// <summary>
        /// (Computed) The monitoring jobs limit for this billing cycle.
        /// </summary>
        public readonly int MonitorsLimit;
        /// <summary>
        /// (Computed) Whether NXD Protection is enabled.
        /// </summary>
        public readonly bool NxdProtectionEnabled;
        /// <summary>
        /// NXD responses for this day.
        /// </summary>
        public readonly int NxdResponses;
        /// <summary>
        /// (Computed) The queries limit for this billing cycle.
        /// </summary>
        public readonly int QueriesLimit;
        /// <summary>
        /// (Computed) The records limit for this billing cycle.
        /// </summary>
        public readonly int RecordsLimit;
        public readonly int? To;
        /// <summary>
        /// (Computed) The total usage count for the metric. Available for `decisions`, `filter-chains`, `monitors`, and `records` metrics.
        /// </summary>
        public readonly int TotalUsage;

        [OutputConstructor]
        private GetBillingUsageResult(
            ImmutableArray<Outputs.GetBillingUsageByNetworkResult> byNetworks,

            int chinaQueriesLimit,

            int cleanQueries,

            bool ddosProtectionEnabled,

            int ddosQueries,

            int decisionsLimit,

            int filterChainsLimit,

            int? from,

            string id,

            bool includeDedicatedDnsNetworkInManagedDnsUsage,

            string metricType,

            int monitorsLimit,

            bool nxdProtectionEnabled,

            int nxdResponses,

            int queriesLimit,

            int recordsLimit,

            int? to,

            int totalUsage)
        {
            ByNetworks = byNetworks;
            ChinaQueriesLimit = chinaQueriesLimit;
            CleanQueries = cleanQueries;
            DdosProtectionEnabled = ddosProtectionEnabled;
            DdosQueries = ddosQueries;
            DecisionsLimit = decisionsLimit;
            FilterChainsLimit = filterChainsLimit;
            From = from;
            Id = id;
            IncludeDedicatedDnsNetworkInManagedDnsUsage = includeDedicatedDnsNetworkInManagedDnsUsage;
            MetricType = metricType;
            MonitorsLimit = monitorsLimit;
            NxdProtectionEnabled = nxdProtectionEnabled;
            NxdResponses = nxdResponses;
            QueriesLimit = queriesLimit;
            RecordsLimit = recordsLimit;
            To = to;
            TotalUsage = totalUsage;
        }
    }
}
